package clients

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"oetp/pkg/apis/ticq/v1alpha1/egressbase"
)

const egressBasePlural = "egressbases"

type EgressBaseInterface interface {
	List(opts metav1.ListOptions) (*egressbase.EgressBaseList, error)
	Get(name string, options metav1.GetOptions) (*egressbase.EgressBase, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Update(*egressbase.EgressBase, metav1.UpdateOptions) (*egressbase.EgressBase, error)
}

type egressBaseClient struct {
	restClient rest.Interface
	ns         string
}

func (c *egressBaseClient) Update(eb *egressbase.EgressBase, opts metav1.UpdateOptions) (*egressbase.EgressBase, error) {
	result := egressbase.EgressBase{}
	err := c.restClient.Put().
		Namespace(c.ns).
		Resource(egressBasePlural).
		Name(eb.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(eb).
		Do(context.Background()).
		Into(&result)

	return &result, err
}

func (c *egressBaseClient) List(opts metav1.ListOptions) (*egressbase.EgressBaseList, error) {
	result := egressbase.EgressBaseList{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource(egressBasePlural).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(context.Background()).
		Into(&result)

	return &result, err
}

func (c *egressBaseClient) Get(name string, opts metav1.GetOptions) (*egressbase.EgressBase, error) {
	result := egressbase.EgressBase{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource(egressBasePlural).
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(context.Background()).
		Into(&result)

	return &result, err
}

func (c *egressBaseClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.restClient.
		Get().
		Namespace(c.ns).
		Resource(egressBasePlural).
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch(context.Background())
}
