package egressnetworkpolicy

import (
	"k8s.io/apimachinery/pkg/runtime"
)

func (eb *EgressNetworkPolicy) DeepCopyInto(out *EgressNetworkPolicy) {
	*out = *eb
	out.TypeMeta = eb.TypeMeta
	eb.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = eb.Spec
}

func (eb *EgressNetworkPolicy) DeepCopy() *EgressNetworkPolicy {
	if eb == nil {
		return nil
	}
	out := new(EgressNetworkPolicy)
	eb.DeepCopyInto(out)
	return out
}

func (eb *EgressNetworkPolicy) DeepCopyObject() runtime.Object {
	if c := eb.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (ebl *EgressNetworkPolicyList) DeepCopyInto(out *EgressNetworkPolicyList) {
	*out = *ebl
	out.TypeMeta = ebl.TypeMeta
	ebl.ListMeta.DeepCopyInto(&out.ListMeta)
	if ebl.Items != nil {
		in, out := &ebl.Items, &out.Items
		*out = make([]EgressNetworkPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

func (ebl *EgressNetworkPolicyList) DeepCopy() *EgressNetworkPolicyList {
	if ebl == nil {
		return nil
	}
	out := new(EgressNetworkPolicyList)
	ebl.DeepCopyInto(out)
	return out
}

func (ebl *EgressNetworkPolicyList) DeepCopyObject() runtime.Object {
	if c := ebl.DeepCopy(); c != nil {
		return c
	}
	return nil
}
