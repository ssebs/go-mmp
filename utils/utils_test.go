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
		if got != -1 {
			t.Fatal("should not have found an index, default is -1 if not found")
		}
	})
}

func TestStringToInt(t *testing.T) {
	t.Run("test string to int", func(t *testing.T) {
		got, err := StringToInt("2")
		want := 2
		if got != want {
			t.Fatalf("expected conversion from string to int. got %v, want %v", got, want)
		}
		if err != nil {
			t.Fatalf("should not have received error. error %s", err)
		}
	})

	t.Run("test fail string to int using word", func(t *testing.T) {
		_, err := StringToInt("bar")
		if err == nil {
			t.Fatalf("expected an error. got %v", err)
		}
	})
}
