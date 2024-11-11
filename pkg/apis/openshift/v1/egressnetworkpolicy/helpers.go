package egressnetworkpolicy

import "reflect"

func (eb *EgressNetworkPolicy) NeedsUpdate(compareTo EgressNetworkPolicy) bool {
	return !reflect.DeepEqual(eb, compareTo)
}
