package utils

import (
	"testing"
)

func TestSliceContains(t *testing.T) {
	t.Run("test matching string", func(t *testing.T) {
		s := []string{"test", "foo"}
		want := "test"
		got, isFound := SliceContains[string](&s, want)

		if !isFound {
			t.Fatalf("got %v want %v", got, want)
		}

		if s[got] != want {
			t.Fatalf("got %v want %v. match != want", got, want)
		}
	})
	t.Run("test missing string", func(t *testing.T) {
		s := []string{"test", "foo"}
		want := "bar"
		got, isFound := SliceContains[string](&s, want)
		if isFound {
			t.Fatalf("should not have found %v in %v", want, s)
		}
		if got != 0 {
			t.Fatal("should not have found an index, default is 0")
		}
	})
}
