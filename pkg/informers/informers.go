package informers

import "oetp/pkg/clients"

type ReconcileNamespace func(string)

type Informers struct {
	clientSet clients.ClientInterface
	reconcile ReconcileNamespace
}

func NewInformers(clientSet clients.ClientInterface, namespace ReconcileNamespace) (informer *Informers) {
	return &Informers{
		clientSet: clientSet,
		reconcile: namespace,
	}
}
