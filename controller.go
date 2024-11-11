package main

import (
	"flag"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	openshiftcontroller "oetp/pkg/apis/openshift"
	v13 "oetp/pkg/apis/openshift/v1"
	ovncontroller "oetp/pkg/apis/ovn"
	v1 "oetp/pkg/apis/ovn/v1"
	v1alpha12 "oetp/pkg/apis/ticq/v1alpha1"
	"oetp/pkg/clients"
	"oetp/pkg/clients/networkopenshift"
	"oetp/pkg/clients/ovn"
	"oetp/pkg/informers"
	logger "oetp/pkg/log"
	"oetp/pkg/signals"
	"oetp/pkg/utils"
	"os"
	"time"
)

var kubeconfig string
var enableEgressFirewall bool
var enableEgressNetworkPolicy bool
var log = logger.Logger()

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "path to Kubernetes config file")
	flag.BoolVar(&enableEgressFirewall, "egressfirewall", true, "enable egress firewall ovn")
	flag.BoolVar(&enableEgressNetworkPolicy, "egressnetworkpolicy", false, "enable egress network policy")
	flag.Parse()
}

func main() {
	var restConfig *rest.Config
	var err error

	if enableEgressFirewall {
		log.Info().Msg("EgressFirewall enabled")
	}

	if enableEgressNetworkPolicy {
		log.Info().Msg("EgressNetworkPolicy enabled")
	}

	signals.SetupSignalHandler()

	if kubeconfig == "" {
		log.Info().Msg("using in-cluster configuration")
		restConfig, err = rest.InClusterConfig()
	} else {
		log.Info().Msgf("using configuration from '%s'", kubeconfig)
		restConfig, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	}

	if err != nil {
		log.Error().Msg(err.Error())
		os.Exit(3)
		return
	}

	_ = v1alpha12.AddToScheme(scheme.Scheme)
	_ = v1.AddToScheme(scheme.Scheme)
	_ = v13.AddToScheme(scheme.Scheme)

	egressBaseClient, err := clients.NewFor(restConfig)
	if err != nil {
		log.Error().Msg(err.Error())
		os.Exit(3)
		return
	}

	egressFirewallClient, err := ovn.New(restConfig)
	if err != nil {
		log.Error().Msg(err.Error())
		os.Exit(3)
		return
	}

	egressNetworkPolicyClient, err := networkopenshift.New(restConfig)
	if err != nil {
		log.Error().Msg(err.Error())
		os.Exit(3)
		return
	}

	if err != nil {
		log.Error().Msg(err.Error())
		os.Exit(3)
		return
	}

	var egressBaseController *cache.Controller
	var egressPolicyController *cache.Controller

	reconcileChannel := make(chan string)

	// empty string = all queues
	reconcileRequest := func(namespace string) {
		go func() {
			reconcileChannel <- namespace
		}()
	}

	newInformers := informers.NewInformers(egressBaseClient, func(namespace string) {
		log.Info().Msgf("Should reconcile for namespace %s", namespace)
		reconcileRequest(namespace)
	})

	_, eb := newInformers.WatchEgressBase()
	egressBaseController = &eb

	_, epc := newInformers.WatchEgressPolicy()
	egressPolicyController = &epc

	go func() {
		for {
			time.Sleep(30 * time.Second)
		}
	}()

	go func() {
		time.Sleep(5 * time.Second)

		log.Info().Msg("Initial reconciliation during startup")
		egressBases, err := egressBaseClient.EgressBase("").List(v12.ListOptions{})
		if err != nil {
			log.Error().Msgf("Error listing egressbases for initial run %s", err.Error())
			return
		}

		for _, egressBase := range (*egressBases).Items {
			reconcileRequest(egressBase.Namespace)
		}
	}()

	for {
		reconcileNamespace := <-reconcileChannel

		if !(*egressBaseController).HasSynced() {
			log.Debug().Msg("waiting for full sync of egressbase")
			continue
		}

		if !(*egressPolicyController).HasSynced() {
			log.Debug().Msg("waiting for full sync of egresspolicies")
			continue
		}

		log.Debug().Msgf("reconciling namespace: %s", reconcileNamespace)

		egressBases, err := egressBaseClient.EgressBase(reconcileNamespace).List(v12.ListOptions{})
		if err != nil {
			log.Error().Msgf("Error listing EgressBase: %s", err.Error())
			continue
		}

		if len(egressBases.Items) > 1 {
			log.Warn().Msgf("More than one egressbase for %s", reconcileNamespace)
			continue
		} else if len(egressBases.Items) == 0 {
			log.Warn().Msgf("No egressbase for %s", reconcileNamespace)
			continue
		}

		egressPolicies, err := egressBaseClient.EgressPolicy(reconcileNamespace).List(v12.ListOptions{})
		if err != nil {
			log.Error().Msgf("Error listing egresspolicies: %s", err.Error())
			continue
		}

		if enableEgressFirewall {
			toCreateEgressFirewall := utils.CreateEgressFirewall(egressBases, egressPolicies.Items, reconcileNamespace)

			firewall := egressFirewallClient.EgressFirewall(reconcileNamespace)

			if len(toCreateEgressFirewall.Spec.Egress) == 0 {
				log.Info().Msgf("Deleting EgressFirewall for %s", reconcileNamespace)
				err := firewall.Delete(toCreateEgressFirewall, v12.DeleteOptions{})
				if err != nil {
					log.Error().Msgf("Error creating EgressFirewall in %s: %s", reconcileNamespace, err.Error())
				}
			} else {
				currentEgresFirewall, err := firewall.Get(ovncontroller.DefaultName, v12.GetOptions{})
				if err != nil {
					// assume not existing
					log.Info().Msgf("Creating EgressFirewall for %s", reconcileNamespace)
					_, err := firewall.Create(toCreateEgressFirewall)
					if err != nil {
						log.Error().Msgf("Error creating EgressFirewall: %s", err.Error())
					}
				} else if !utils.CanWeTakeOwnership((*currentEgresFirewall).OwnerReferences, (*toCreateEgressFirewall).OwnerReferences) {
					log.Info().Msgf("EgressFirewall cant update because it is not the owner or cant take ownership %s", reconcileNamespace)
				} else if currentEgresFirewall.NeedsUpdate(*toCreateEgressFirewall) {
					log.Info().Msgf("Updating EgressFirewall for %s", reconcileNamespace)

					// update fields which need an update
					currentEgresFirewall.Spec.Egress = toCreateEgressFirewall.Spec.Egress
					currentEgresFirewall.ObjectMeta.OwnerReferences = toCreateEgressFirewall.ObjectMeta.OwnerReferences

					_, err := firewall.Update(currentEgresFirewall, v12.UpdateOptions{})
					if err != nil {
						log.Error().Msgf("Error updating EgressFirewall: %s", err.Error())
					}
				} else {
					log.Info().Msgf("EgressFirewall doesnt need an update for %s", reconcileNamespace)
				}
			}
		}

		if enableEgressNetworkPolicy {
			toCreateEgressNetworkPolicy := utils.CreateEgressNetworkPolicy(egressBases, egressPolicies.Items, reconcileNamespace)

			egressNetworkPolicy := egressNetworkPolicyClient.EgressNetworkPolicy(reconcileNamespace)
			if len(toCreateEgressNetworkPolicy.Spec.Egress) == 0 {
				log.Info().Msgf("Deleting EgressNetworkPolicy for %s", reconcileNamespace)
				err := egressNetworkPolicy.Delete(toCreateEgressNetworkPolicy, v12.DeleteOptions{})
				if err != nil {
					log.Error().Msgf("Error deleting EgressNetworkPolicy in %s: %s", reconcileNamespace, err.Error())
				}
			} else {

				currentEgressNetworkPolicy, err := egressNetworkPolicy.Get(openshiftcontroller.DefaultName, v12.GetOptions{})
				if err != nil {
					// assume not existing
					log.Info().Msgf("Creating EgressNetworkPolicy for %s", reconcileNamespace)
					_, err := egressNetworkPolicy.Create(toCreateEgressNetworkPolicy)
					log.Info().Msg(err.Error())
					if err != nil {
						log.Error().Msgf("Error creating EgressNetworkPolicy: %s", err.Error())
					}
				} else if !utils.CanWeTakeOwnership((*currentEgressNetworkPolicy).OwnerReferences, (*toCreateEgressNetworkPolicy).OwnerReferences) {
					log.Info().Msgf("EgressNetworkPolicy cant update because it is not the owner or cant take ownership %s", reconcileNamespace)
				} else if currentEgressNetworkPolicy.NeedsUpdate(*toCreateEgressNetworkPolicy) {
					log.Info().Msgf("Updating EgressNetworkPolicy for %s", reconcileNamespace)

					// update fields which need an update
					currentEgressNetworkPolicy.Spec.Egress = toCreateEgressNetworkPolicy.Spec.Egress
					currentEgressNetworkPolicy.ObjectMeta.OwnerReferences = toCreateEgressNetworkPolicy.ObjectMeta.OwnerReferences

					_, err := egressNetworkPolicy.Update(currentEgressNetworkPolicy, v12.UpdateOptions{})
					if err != nil {
						log.Error().Msgf("Error updating EgressNetworkPolicy: %s", err.Error())
					}
				} else {
					log.Info().Msgf("EgressNetworkPolicy doesnt need an update for %s", reconcileNamespace)
				}
			}
		}
	}
}
