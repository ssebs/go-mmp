package modelstests

import (
	"testing"

	"github.com/ssebs/go-mmp/models"
	"github.com/stretchr/testify/assert"
)

func TestGUIModeModel(t *testing.T) {
	got := models.NOTSET

	got.Set("NORMAL")
	assert.Equal(t, models.NORMAL, got)

	err := got.Set("INVALID")
	assert.Error(t, err)
	assert.Equal(t, "NOTSET", got.Type())
}
func TestGUIModeYAML(t *testing.T) {
	// Test Mrashal and Unmarshal YAML
}
