package networkopenshift

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	controller "oetp/pkg/apis/openshift"
)

type Client struct {
	restClient rest.Interface
}

type ClientInterface interface {
	EgressNetworkPolicy(namespace string) EgressNetworkPolicyInterface
}

func New(c *rest.Config) (*Client, error) {
	config := *c
	config.ContentConfig.GroupVersion = &schema.GroupVersion{Group: controller.Name, Version: controller.Version}
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	config.UserAgent = rest.DefaultKubernetesUserAgent()

	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}

	return &Client{restClient: client}, nil
}

func (c *Client) EgressNetworkPolicy(namespace string) EgressNetworkPolicyInterface {
	return &client{
		restClient: c.restClient,
		ns:         namespace,
	}
}
