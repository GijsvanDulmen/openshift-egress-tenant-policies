package informers

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
	"oetp/pkg/apis/ticq/v1alpha1/egressbase"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"time"
)

func (informer *Informers) WatchEgressBase() (cache.Store, cache.Controller) {
	store, controller := cache.NewInformer(
		&cache.ListWatch{
			ListFunc: func(lo metav1.ListOptions) (result runtime.Object, err error) {
				return informer.clientSet.EgressBase("").List(lo)
			},
			WatchFunc: func(lo metav1.ListOptions) (watch.Interface, error) {
				return informer.clientSet.EgressBase("").Watch(lo)
			},
		},
		&egressbase.EgressBase{},
		1*time.Minute,
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				typed := obj.(*egressbase.EgressBase)
				log.Print(*typed, "added")
				informer.ReconcileEgressBase(typed)
			},
			UpdateFunc: func(old, new interface{}) {
				typed := new.(*egressbase.EgressBase)
				log.Print(*typed, "updated")
				informer.ReconcileEgressBase(typed)
			},
			DeleteFunc: func(obj interface{}) {
				typed := obj.(*egressbase.EgressBase)
				log.Print(*typed, "deleted")
				informer.ReconcileEgressBase(typed)
			},
		},
	)

	go controller.Run(wait.NeverStop)
	return store, controller
}

func (informer *Informers) ReconcileEgressBase(egressBase *egressbase.EgressBase) {
	if egressBase.ObjectMeta.DeletionTimestamp.IsZero() {
		if !controllerutil.ContainsFinalizer(egressBase, finalizerName) {
			log.Print(*egressBase, "adding finalizer")
			controllerutil.AddFinalizer(egressBase, finalizerName)

			_, err := informer.clientSet.EgressBase(egressBase.ObjectMeta.Namespace).Update(egressBase, metav1.UpdateOptions{})
			if err != nil {
				log.Print(*egressBase, err.Error())
			}
			return
		}
		informer.reconcile(egressBase.Namespace)
	} else {
		if controllerutil.ContainsFinalizer(egressBase, finalizerName) {
			controllerutil.RemoveFinalizer(egressBase, finalizerName)

			log.Print(*egressBase, "removing finalizer")

			_, err := informer.clientSet.EgressBase(egressBase.ObjectMeta.Namespace).Update(egressBase, metav1.UpdateOptions{})
			if err != nil {
				log.Print(*egressBase, "could not remove finalizer")
				log.Print(*egressBase, err.Error())
			}
		}
		informer.reconcile(egressBase.Namespace)
	}
}
