package egressnetworkpolicy

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type EgressNetworkPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec Spec `json:"spec"`
}

type Spec struct {
	Egress []EgressList `json:"egress"`
}

type EgressList struct {
	Type string   `json:"type"`
	To   EgressTo `json:"to"`
}

type EgressTo struct {
	CidrSelector *string `json:"cidrSelector"`
	DnsName      *string `json:"dnsName"`
}

type EgressNetworkPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []EgressNetworkPolicy `json:"items"`
}
