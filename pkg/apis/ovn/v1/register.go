package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	controller "oetp/pkg/apis/ovn"
	"oetp/pkg/apis/ovn/v1/egressfirewall"
)

var SchemeGroupVersion = schema.GroupVersion{Group: controller.GroupName, Version: controller.Version}

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&egressfirewall.EgressFirewall{},
		&egressfirewall.EgressFirewallList{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
