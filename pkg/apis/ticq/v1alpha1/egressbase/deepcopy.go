package egressbase

import (
	"k8s.io/apimachinery/pkg/runtime"
)

func (eb *EgressBase) DeepCopyInto(out *EgressBase) {
	*out = *eb
	out.TypeMeta = eb.TypeMeta
	eb.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = eb.Spec
}

func (eb *EgressBase) DeepCopy() *EgressBase {
	if eb == nil {
		return nil
	}
	out := new(EgressBase)
	eb.DeepCopyInto(out)
	return out
}

func (eb *EgressBase) DeepCopyObject() runtime.Object {
	if c := eb.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (ebl *EgressBaseList) DeepCopyInto(out *EgressBaseList) {
	*out = *ebl
	out.TypeMeta = ebl.TypeMeta
	ebl.ListMeta.DeepCopyInto(&out.ListMeta)
	if ebl.Items != nil {
		in, out := &ebl.Items, &out.Items
		*out = make([]EgressBase, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

func (ebl *EgressBaseList) DeepCopy() *EgressBaseList {
	if ebl == nil {
		return nil
	}
	out := new(EgressBaseList)
	ebl.DeepCopyInto(out)
	return out
}

func (ebl *EgressBaseList) DeepCopyObject() runtime.Object {
	if c := ebl.DeepCopy(); c != nil {
		return c
	}
	return nil
}
