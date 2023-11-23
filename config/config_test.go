package config

import (
	"os"
	"reflect"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	configPath := "../res/defaultConfig.yml"
	t.Run("make sure loadconfig works", func(t *testing.T) {
		// expected usage
		c, err := NewConfigFromFile(configPath)
		if err != nil {
			t.Fatalf("expected no error, got %s", err)
		}

		// manual to test
		// TODO: make this path agnostic to where you are running the test
		f, err := os.Open(configPath)
		if err != nil {
			t.Fatalf("could not open file for test, err: %s. %+v", err, f)
		}
		defer f.Close()
		c2, err := LoadConfig(f)
		if err != nil {
			t.Fatalf("expected no error, got %s", err)
		}

		if !reflect.DeepEqual(c, c2) {
			t.Fatalf("expected NewConfigFromFile('') to load res/defaultconfig.yml. got %+v, want %+v", c2, c)
		}
	})
}
