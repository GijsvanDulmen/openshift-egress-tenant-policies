package egressfirewall

import (
	"reflect"
)

func (egressFirewall *EgressFirewall) NeedsUpdate(compareTo EgressFirewall) bool {
	return !reflect.DeepEqual(egressFirewall, compareTo)
}
