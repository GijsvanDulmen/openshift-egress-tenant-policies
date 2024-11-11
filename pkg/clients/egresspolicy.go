package clients

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"oetp/pkg/apis/ticq/v1alpha1/egresspolicy"
)

const egressPolicyPlural = "egresspolicies"

type EgressPolicyInterface interface {
	List(opts metav1.ListOptions) (*egresspolicy.EgressPolicyList, error)
	Get(name string, options metav1.GetOptions) (*egresspolicy.EgressPolicy, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Update(*egresspolicy.EgressPolicy, metav1.UpdateOptions) (*egresspolicy.EgressPolicy, error)
}

type egressPolicyClient struct {
	restClient rest.Interface
	ns         string
}

func (c *egressPolicyClient) Update(ep *egresspolicy.EgressPolicy, opts metav1.UpdateOptions) (*egresspolicy.EgressPolicy, error) {
	result := egresspolicy.EgressPolicy{}
	err := c.restClient.Put().
		Namespace(c.ns).
		Resource(egressPolicyPlural).
		Name(ep.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(ep).
		Do(context.Background()).
		Into(&result)

	return &result, err
}

func (c *egressPolicyClient) List(opts metav1.ListOptions) (*egresspolicy.EgressPolicyList, error) {
	result := egresspolicy.EgressPolicyList{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource(egressPolicyPlural).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(context.Background()).
		Into(&result)

	return &result, err
}

func (c *egressPolicyClient) Get(name string, opts metav1.GetOptions) (*egresspolicy.EgressPolicy, error) {
	result := egresspolicy.EgressPolicy{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource(egressPolicyPlural).
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(context.Background()).
		Into(&result)

	return &result, err
}

func (c *egressPolicyClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.restClient.
		Get().
		Namespace(c.ns).
		Resource(egressPolicyPlural).
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch(context.Background())
}
