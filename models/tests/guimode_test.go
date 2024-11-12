package modelstests

import (
	"testing"

	"github.com/ssebs/go-mmp/models"
	"github.com/stretchr/testify/assert"
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
		got.MarshalYAML()
	})
}
