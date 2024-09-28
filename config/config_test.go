package config

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	cliFlags := &CLIFlags{
		GUIMode:     TESTING,
		ConfigPath:  "../res/defaultConfig.yml",
		ResetConfig: false,
	}

	t.Run("make sure loadconfig works and matches the contents in defaultConfig", func(t *testing.T) {
		// expected usage
		c, err := NewConfig(cliFlags)
		if err != nil {
			t.Fatalf("expected no error, got %s", err)
		}

		// manual to test
		// TODO: make this path agnostic to where you are running the test
		f, err := os.Open(cliFlags.ConfigPath)
		if err != nil {
			t.Fatalf("could not open file for test, err: %s. %+v", err, f)
		}
		defer f.Close()
		c2, err := loadConfig(f)
		if err != nil {
			t.Fatalf("expected no error, got %s", err)
		}

		if !reflect.DeepEqual(c, c2) {
			t.Fatalf("res/defaultconfig.yml not the same as output from NewConfig(). got %+v, want %+v", c2, c)
		}
	})

	t.Run("test various cliflags to confirm correct paths", func(t *testing.T) {
		// with --path, no defaultconfig

		// with invalid --path, no defaultconfig

		// without --path, no defaultconfig

		// with --path, has defaultconfig

		// with invalid --path, has defaultconfig

		// without --path, has defaultconfig

	})

	t.Run("test loading file to confirm error handling", func(t *testing.T) {

	})

	t.Run("test expansion of ~/mmpConfig.yml", func(t *testing.T) {
		home, _ := os.UserHomeDir()
		short := "~/mmpConfig.yml"
		expected := filepath.Clean(filepath.FromSlash(home) + "/mmpConfig.yml")

		if short != DefaultConfigPath {
			t.Fatal("DefaultConfigPath has changed from ~/mmpConfig.yml")
		}

		defaultFullPath, err := filepath.Abs(os.ExpandEnv(DefaultConfigPath))
		if err != nil {
			t.Fatalf("failed to expand ~/mmpConfig.yml to a full path, %e", err)
		}

		if expected != defaultFullPath {
			t.Fatalf("expanded home env not as expected. got %s, want %s", defaultFullPath, expected)
		}

	})
}
