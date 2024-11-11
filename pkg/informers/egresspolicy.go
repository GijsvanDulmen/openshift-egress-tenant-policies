package informers

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
	"oetp/pkg/apis/ticq/v1alpha1/egresspolicy"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"time"
)

func (informer *Informers) WatchEgressPolicy() (cache.Store, cache.Controller) {
	store, controller := cache.NewInformer(
		&cache.ListWatch{
			ListFunc: func(lo metav1.ListOptions) (result runtime.Object, err error) {
				return informer.clientSet.EgressPolicy("").List(lo)
			},
			WatchFunc: func(lo metav1.ListOptions) (watch.Interface, error) {
				return informer.clientSet.EgressPolicy("").Watch(lo)
			},
		},
		&egresspolicy.EgressPolicy{},
		1*time.Minute,
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				typed := obj.(*egresspolicy.EgressPolicy)
				log.Print(*typed, "added")
				informer.ReconcileEgressPolicy(typed)
			},
			UpdateFunc: func(old, new interface{}) {
				typed := new.(*egresspolicy.EgressPolicy)
				log.Print(*typed, "updated")
				informer.ReconcileEgressPolicy(typed)
			},
			DeleteFunc: func(obj interface{}) {
				typed := obj.(*egresspolicy.EgressPolicy)
				log.Print(*typed, "deleted")
				informer.ReconcileEgressPolicy(typed)
			},
		},
	)

	go controller.Run(wait.NeverStop)
	return store, controller
}

func (informer *Informers) ReconcileEgressPolicy(egressPolicy *egresspolicy.EgressPolicy) {
	if egressPolicy.ObjectMeta.DeletionTimestamp.IsZero() {
		if !controllerutil.ContainsFinalizer(egressPolicy, finalizerName) {
			log.Print(*egressPolicy, "adding finalizer")
			controllerutil.AddFinalizer(egressPolicy, finalizerName)

			_, err := informer.clientSet.EgressPolicy(egressPolicy.ObjectMeta.Namespace).Update(egressPolicy, metav1.UpdateOptions{})
			if err != nil {
				log.Print(*egressPolicy, err.Error())
			}
			return
		}
		informer.reconcile(egressPolicy.Namespace)
	} else {
		if controllerutil.ContainsFinalizer(egressPolicy, finalizerName) {
			controllerutil.RemoveFinalizer(egressPolicy, finalizerName)

			log.Print(*egressPolicy, "removing finalizer")

			_, err := informer.clientSet.EgressPolicy(egressPolicy.ObjectMeta.Namespace).Update(egressPolicy, metav1.UpdateOptions{})
			if err != nil {
				log.Print(*egressPolicy, "could not remove finalizer")
				log.Print(*egressPolicy, err.Error())
			}
		}
		informer.reconcile(egressPolicy.Namespace)
	}
}
