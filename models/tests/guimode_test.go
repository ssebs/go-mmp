package models_test

import (
	"fmt"
	"testing"

	"github.com/ssebs/go-mmp/models"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestGUIModeModel(t *testing.T) {
	got := models.NOTSET

	t.Run("Test Set", func(t *testing.T) {
		got.Set("NORMAL")
		assert.Equal(t, models.NORMAL, got)
	})
	t.Run("Test Setting invalid option", func(t *testing.T) {
		err := got.Set("INVALID")
		assert.Error(t, err)
		assert.Equal(t, "NOTSET", got.Type())
	})

	t.Run("Test String/Type/Get", func(t *testing.T) {
		got = models.TESTING
		assert.Equal(t, "TESTING", got.String())
		assert.Equal(t, "TESTING", got.Type())
	})
}
func TestGUIModeYAML(t *testing.T) {
	got := models.TESTING
	t.Run("Test MarshalYAML", func(t *testing.T) {
		gotYaml, err := got.MarshalYAML()
		assert.Nil(t, err)
		assert.Equal(t, "TESTING", gotYaml)
	})

	t.Run("Test UnmarshalYAML", func(t *testing.T) {
		want := TestStruct{StrField: "TEST", GMode: models.TESTING}
		got := TestStruct{}

		sampleData := []byte("StrField: \"TEST\"\nGMode: \"TESTING\"\n")
		err := yaml.Unmarshal(sampleData, &got)
		assert.Nil(t, err)

		fmt.Println(want)
		assert.Equal(t, want, got)
	})
}

type TestStruct struct {
	StrField string         `yaml:"StrField"`
	GMode    models.GUIMode `yaml:"GMode"`
}
