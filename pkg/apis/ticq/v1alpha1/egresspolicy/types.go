package egresspolicy

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type EgressPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec Spec `json:"spec"`
}

type Spec struct {
	Egress []EgressList `json:"egress"`
	Groups []string     `json:"groups"`
}

type EgressList struct {
	Type    string  `json:"type"`
	DnsName *string `json:"dnsName"`
	Cidr    *string `json:"cidr"`
}

type EgressPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []EgressPolicy `json:"items"`
}
