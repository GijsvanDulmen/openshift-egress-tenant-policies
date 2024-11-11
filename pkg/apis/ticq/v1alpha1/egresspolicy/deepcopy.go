package egresspolicy

import (
	"k8s.io/apimachinery/pkg/runtime"
)

func (ep *EgressPolicy) DeepCopyInto(out *EgressPolicy) {
	*out = *ep
	out.TypeMeta = ep.TypeMeta
	ep.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = ep.Spec
}

func (ep *EgressPolicy) DeepCopy() *EgressPolicy {
	if ep == nil {
		return nil
	}
	out := new(EgressPolicy)
	ep.DeepCopyInto(out)
	return out
}

func (ep *EgressPolicy) DeepCopyObject() runtime.Object {
	if c := ep.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (epl *EgressPolicyList) DeepCopyInto(out *EgressPolicyList) {
	*out = *epl
	out.TypeMeta = epl.TypeMeta
	epl.ListMeta.DeepCopyInto(&out.ListMeta)
	if epl.Items != nil {
		in, out := &epl.Items, &out.Items
		*out = make([]EgressPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

func (epl *EgressPolicyList) DeepCopy() *EgressPolicyList {
	if epl == nil {
		return nil
	}
	out := new(EgressPolicyList)
	epl.DeepCopyInto(out)
	return out
}

func (epl *EgressPolicyList) DeepCopyObject() runtime.Object {
	if c := epl.DeepCopy(); c != nil {
		return c
	}
	return nil
}
