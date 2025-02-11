package models_test

import (
	"testing"
	"time"

	"github.com/ssebs/go-mmp/models"
	"github.com/stretchr/testify/assert"
)

func TestMetadataModel(t *testing.T) {
	want := &models.Metadata{
		Columns:        1,
		SerialPortName: "TEST",
		SerialBaudRate: 1,
		Delay:          time.Second,
		GUIMode:        models.TESTING,
	}
	t.Run("test constructor", func(t *testing.T) {
		got := models.NewMetadata("TEST", 1, models.TESTING, 1, time.Second, 0)
		assert.Equal(t, got, want)
	})

	// 	t.Run("Test parser", func(t *testing.T) {
	// 		expectedStr := `Columns: 1
	// SerialPortName: TEST
	// SerialBaudRate: 1
	// Delay: 1s
	// GUIMode: TESTING`

	// 		assert.YAMLEq(
	// 			t,
	// 			expectedStr,
	// 			want.String(),
	// 		)
	// 	})

	t.Run("Test default constructor", func(t *testing.T) {
		want = &models.Metadata{
			Columns:        2,
			SerialPortName: "",
			SerialBaudRate: 9600,
			Delay:          125 * time.Millisecond,
			GUIMode:        models.GUIOnly,
		}
		got := models.NewDefaultMetadata()
		assert.Equal(t, want, got)
	})
}
