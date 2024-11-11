package egressfirewall

import (
	"k8s.io/apimachinery/pkg/runtime"
)

func (egressFirewall *EgressFirewall) DeepCopyInto(out *EgressFirewall) {
	*out = *egressFirewall
	out.TypeMeta = egressFirewall.TypeMeta
	egressFirewall.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = egressFirewall.Spec
}

func (egressFirewall *EgressFirewall) DeepCopy() *EgressFirewall {
	if egressFirewall == nil {
		return nil
	}
	out := new(EgressFirewall)
	egressFirewall.DeepCopyInto(out)
	return out
}

func (egressFirewall *EgressFirewall) DeepCopyObject() runtime.Object {
	if c := egressFirewall.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (egressFirewallList *EgressFirewallList) DeepCopyInto(out *EgressFirewallList) {
	*out = *egressFirewallList
	out.TypeMeta = egressFirewallList.TypeMeta
	egressFirewallList.ListMeta.DeepCopyInto(&out.ListMeta)
	if egressFirewallList.Items != nil {
		in, out := &egressFirewallList.Items, &out.Items
		*out = make([]EgressFirewall, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

func (egressFirewallList *EgressFirewallList) DeepCopy() *EgressFirewallList {
	if egressFirewallList == nil {
		return nil
	}
	out := new(EgressFirewallList)
	egressFirewallList.DeepCopyInto(out)
	return out
}

func (egressFirewallList *EgressFirewallList) DeepCopyObject() runtime.Object {
	if c := egressFirewallList.DeepCopy(); c != nil {
		return c
	}
	return nil
}
