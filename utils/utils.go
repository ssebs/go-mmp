package utils

import (
	"reflect"
	"strconv"
)

const ProjectName = "Go-MMP"

// Errors
type ErrCannotParseIntFromEmptyString struct{}

func (e ErrCannotParseIntFromEmptyString) Error() string {
	return "cannot parse int from empty string"
}

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

// Convert a string to an integer, if possible.
// If not, return an error
func StringToInt(s string) (int, error) {
	i64, err := strconv.ParseInt(s, 10, 64)
	if err != nil && err.Error() == "strconv.ParseInt: parsing \"\": invalid syntax" {
		return 0, &ErrCannotParseIntFromEmptyString{}
	}
	return int(i64), err
}
