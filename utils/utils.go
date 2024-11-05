package utils

import (
	"fmt"
	"io"
	"os"
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

// CheckFileExists will check if file is there, and is readable
// Returns a bool with the answer
func CheckFileExists(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	if fi.Size() == 0 {
		return false
	}
	return true
}

// CopyFile will copy a file from src to dest.
// Returns an error if one occurs
func CopyFile(src, dest string) error {
	// Open the source file for reading
	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file: %v", err)
	}
	defer srcFile.Close()
	// Open the destination file for writing
	dstFile, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %v", err)
	}
	defer dstFile.Close()
	// Copy the contents of the source file to the destination file
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return fmt.Errorf("failed to copy file: %v", err)
	}
	return nil
}

func GetKeyVal(m map[string]string) (string, string) {
	for k, v := range m {
		return k, v
	}
	return "", ""
}
