package utils

import "reflect"

// Return index of match, isFound if a value in in a slice
// If there's no match, index is -1
func SliceContains[T comparable](slic *[]T, item T) (index int, isFound bool) {
	for i, elem := range *slic {
		if reflect.DeepEqual(elem, item) {
			return i, true
		}
	}
	return -1, false
}
