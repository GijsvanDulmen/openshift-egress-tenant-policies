package utils

import (
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	controller "oetp/pkg/apis/ovn"
	"oetp/pkg/apis/ovn/v1/egressfirewall"
	"oetp/pkg/apis/ticq/v1alpha1/egressbase"
	"oetp/pkg/apis/ticq/v1alpha1/egresspolicy"
)

func CreateEgressFirewall(egressBases *egressbase.EgressBaseList, egressPolicies []egresspolicy.EgressPolicy, reconcileNamespace string) *egressfirewall.EgressFirewall {
	var list []egressfirewall.EgressList

	// add before
	for _, i2 := range egressBases.Items[0].Spec.Before {
		list = append(list, egressfirewall.EgressList{
			Type: i2.Type,
			To: egressfirewall.EgressTo{
				CidrSelector: i2.Cidr,
				DnsName:      i2.DnsName,
			},
		})
	}

	// add policies
	for _, policy := range egressPolicies {
		for _, groupName := range policy.Spec.Groups {
			if groupName != "" {
				for _, group := range egressBases.Items[0].Spec.Groups {
					if group.Name == groupName {
						for _, groupPolicyRule := range group.Egress {
							list = append(list, egressfirewall.EgressList{
								Type: groupPolicyRule.Type,
								To: egressfirewall.EgressTo{
									CidrSelector: groupPolicyRule.Cidr,
									DnsName:      groupPolicyRule.DnsName,
								},
							})
						}
					}
				}
			}
		}

		for _, policyRule := range policy.Spec.Egress {
			list = append(list, egressfirewall.EgressList{
				Type: policyRule.Type,
				To: egressfirewall.EgressTo{
					CidrSelector: policyRule.Cidr,
					DnsName:      policyRule.DnsName,
				},
			})
		}
	}

	// add after
	for _, i2 := range egressBases.Items[0].Spec.After {
		list = append(list, egressfirewall.EgressList{
			Type: i2.Type,
			To: egressfirewall.EgressTo{
				CidrSelector: i2.Cidr,
				DnsName:      i2.DnsName,
			},
		})
	}

	eg := &egressfirewall.EgressFirewall{
		TypeMeta: v12.TypeMeta{
			Kind:       controller.Name,
			APIVersion: controller.GroupName + "/" + controller.Version,
		},
		ObjectMeta: v12.ObjectMeta{
			Name:      controller.DefaultName,
			Namespace: reconcileNamespace,
			OwnerReferences: []v12.OwnerReference{
				{
					APIVersion: egressBases.Items[0].APIVersion,
					Kind:       egressBases.Items[0].Kind,
					Name:       egressBases.Items[0].Name,
					UID:        egressBases.Items[0].UID,
				},
			},
		},
		Spec: egressfirewall.Spec{
			Egress: list,
		},
	}
	return eg
}
