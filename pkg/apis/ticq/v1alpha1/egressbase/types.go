package egressbase

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type EgressBase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec Spec `json:"spec"`
}

type Spec struct {
	Before []EgressList  `json:"before"`
	After  []EgressList  `json:"after"`
	Groups []EgressGroup `json:"groups"`
}

type EgressGroup struct {
	Name   string       `json:"name"`
	Egress []EgressList `json:"egress"`
}

type EgressList struct {
	Type    string  `json:"type"`
	DnsName *string `json:"dnsName"`
	Cidr    *string `json:"cidr"`
}

type EgressBaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []EgressBase `json:"items"`
}
