package clients

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	controller "oetp/pkg/apis/ticq"
)

type Client struct {
	restClient rest.Interface
}

type ClientInterface interface {
	EgressBase(namespace string) EgressBaseInterface
	EgressPolicy(namespace string) EgressPolicyInterface
}

func NewFor(c *rest.Config) (*Client, error) {
	config := *c
	config.ContentConfig.GroupVersion = &schema.GroupVersion{Group: controller.GroupName, Version: controller.Version}
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	config.UserAgent = rest.DefaultKubernetesUserAgent()

	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}

	return &Client{restClient: client}, nil
}

func (c *Client) EgressBase(namespace string) EgressBaseInterface {
	return &egressBaseClient{
		restClient: c.restClient,
		ns:         namespace,
	}
}

func (c *Client) EgressPolicy(namespace string) EgressPolicyInterface {
	return &egressPolicyClient{
		restClient: c.restClient,
		ns:         namespace,
	}
}
