package utils

import "reflect"

func SliceContains[T comparable](slic *[]T, item T) (index int, found bool) {
	for i, elem := range *slic {
		if reflect.DeepEqual(elem, item) {
			return i, true
		}
	}
	return 0, false
}
