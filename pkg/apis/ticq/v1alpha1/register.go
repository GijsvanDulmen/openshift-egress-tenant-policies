package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	controller "oetp/pkg/apis/ticq"
	"oetp/pkg/apis/ticq/v1alpha1/egressbase"
	"oetp/pkg/apis/ticq/v1alpha1/egresspolicy"
)

var SchemeGroupVersion = schema.GroupVersion{Group: controller.GroupName, Version: "v1alpha1"}

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&egressbase.EgressBase{},
		&egressbase.EgressBaseList{},
		&egresspolicy.EgressPolicy{},
		&egresspolicy.EgressPolicyList{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
