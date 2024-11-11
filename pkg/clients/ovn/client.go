package ovn

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type Client struct {
	restClient rest.Interface
}

type ClientInterface interface {
	EgressFirewall(namespace string) EgressFirewallInterface
}

func New(c *rest.Config) (*Client, error) {
	config := *c
	config.ContentConfig.GroupVersion = &schema.GroupVersion{Group: "k8s.ovn.org", Version: "v1"}
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	config.UserAgent = rest.DefaultKubernetesUserAgent()

	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}

	return &Client{restClient: client}, nil
}

func (c *Client) EgressFirewall(namespace string) EgressFirewallInterface {
	return &client{
		restClient: c.restClient,
		ns:         namespace,
	}
}
