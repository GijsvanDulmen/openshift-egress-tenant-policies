package ovn

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"oetp/pkg/apis/ovn/v1/egressfirewall"
)

const egressFirewallPlural = "egressfirewalls"

type EgressFirewallInterface interface {
	List(opts metav1.ListOptions) (*egressfirewall.EgressFirewallList, error)
	Get(name string, options metav1.GetOptions) (*egressfirewall.EgressFirewall, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Update(*egressfirewall.EgressFirewall, metav1.UpdateOptions) (*egressfirewall.EgressFirewall, error)
	Create(*egressfirewall.EgressFirewall) (*egressfirewall.EgressFirewall, error)
	Delete(*egressfirewall.EgressFirewall, metav1.DeleteOptions) error
}

type client struct {
	restClient rest.Interface
	ns         string
}

func (c *client) Create(ef *egressfirewall.EgressFirewall) (*egressfirewall.EgressFirewall, error) {
	result := egressfirewall.EgressFirewall{}
	err := c.restClient.
		Post().
		Namespace(c.ns).
		Resource(egressFirewallPlural).
		Body(ef).
		Do(context.Background()).
		Into(&result)

	return &result, err
}

func (c *client) Update(ep *egressfirewall.EgressFirewall, opts metav1.UpdateOptions) (*egressfirewall.EgressFirewall, error) {
	result := egressfirewall.EgressFirewall{}
	err := c.restClient.Put().
		Namespace(c.ns).
		Resource(egressFirewallPlural).
		Name(ep.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(ep).
		Do(context.Background()).
		Into(&result)

	return &result, err
}

func (c *client) List(opts metav1.ListOptions) (*egressfirewall.EgressFirewallList, error) {
	result := egressfirewall.EgressFirewallList{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource(egressFirewallPlural).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(context.Background()).
		Into(&result)

	return &result, err
}

func (c *client) Delete(ef *egressfirewall.EgressFirewall, opts metav1.DeleteOptions) error {
	result := egressfirewall.EgressFirewall{}
	err := c.restClient.
		Delete().
		Namespace(c.ns).
		Resource(egressFirewallPlural).
		Name(ef.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(context.Background()).
		Into(&result)

	return err
}

func (c *client) Get(name string, opts metav1.GetOptions) (*egressfirewall.EgressFirewall, error) {
	result := egressfirewall.EgressFirewall{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource(egressFirewallPlural).
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
		Resource(egressFirewallPlural).
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch(context.Background())
}
