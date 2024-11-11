package utils

import (
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CanWeTakeOwnership(old []v12.OwnerReference, new []v12.OwnerReference) bool {
	if len(old) == 0 {
		return true // take over ownership
	}
	if len(old) == 1 {
		if old[0].Name == new[0].Name {
			if old[0].Kind == new[0].Kind {
				if old[0].APIVersion == new[0].APIVersion {
					return true
				}
			}
		}
	}
	return false
}
