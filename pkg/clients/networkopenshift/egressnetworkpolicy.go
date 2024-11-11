package networkopenshift

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"oetp/pkg/apis/openshift/v1/egressnetworkpolicy"
)

const egressNetworkPolicyPlural = "egressnetworkpolicis"

type EgressNetworkPolicyInterface interface {
	List(opts metav1.ListOptions) (*egressnetworkpolicy.EgressNetworkPolicyList, error)
	Get(name string, options metav1.GetOptions) (*egressnetworkpolicy.EgressNetworkPolicy, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Update(*egressnetworkpolicy.EgressNetworkPolicy, metav1.UpdateOptions) (*egressnetworkpolicy.EgressNetworkPolicy, error)
	Create(*egressnetworkpolicy.EgressNetworkPolicy) (*egressnetworkpolicy.EgressNetworkPolicy, error)
	Delete(*egressnetworkpolicy.EgressNetworkPolicy, metav1.DeleteOptions) error
}

type client struct {
	restClient rest.Interface
	ns         string
}

func (c *client) Create(ef *egressnetworkpolicy.EgressNetworkPolicy) (*egressnetworkpolicy.EgressNetworkPolicy, error) {
	result := egressnetworkpolicy.EgressNetworkPolicy{}
	err := c.restClient.
		Post().
		Namespace(c.ns).
		Resource(egressNetworkPolicyPlural).
		Body(ef).
		Do(context.Background()).
		Into(&result)

	return &result, err
}

func (c *client) Update(ep *egressnetworkpolicy.EgressNetworkPolicy, opts metav1.UpdateOptions) (*egressnetworkpolicy.EgressNetworkPolicy, error) {
	result := egressnetworkpolicy.EgressNetworkPolicy{}
	err := c.restClient.Put().
		Namespace(c.ns).
		Resource(egressNetworkPolicyPlural).
		Name(ep.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(ep).
		Do(context.Background()).
		Into(&result)

	return &result, err
}

func (c *client) Delete(ef *egressnetworkpolicy.EgressNetworkPolicy, opts metav1.DeleteOptions) error {
	result := egressnetworkpolicy.EgressNetworkPolicy{}
	err := c.restClient.
		Delete().
		Namespace(c.ns).
		Resource(egressNetworkPolicyPlural).
		Name(ef.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(context.Background()).
		Into(&result)

	return err
}

func (c *client) List(opts metav1.ListOptions) (*egressnetworkpolicy.EgressNetworkPolicyList, error) {
	result := egressnetworkpolicy.EgressNetworkPolicyList{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource(egressNetworkPolicyPlural).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(context.Background()).
		Into(&result)

	return &result, err
}

func (c *client) Get(name string, opts metav1.GetOptions) (*egressnetworkpolicy.EgressNetworkPolicy, error) {
	result := egressnetworkpolicy.EgressNetworkPolicy{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource(egressNetworkPolicyPlural).
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(context.Background()).
		Into(&result)

	return &result, err
}

func (c *client) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.restClient.
		Get().
		Namespace(c.ns).
		Resource(egressNetworkPolicyPlural).
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch(context.Background())
}
