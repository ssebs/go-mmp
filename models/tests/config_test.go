package modelstests

import (
	"testing"

	"github.com/ssebs/go-mmp/models"
	"github.com/stretchr/testify/assert"
)

func TestConfigM(t *testing.T) {
	_macros := []models.Macro{
		models.NewMacro("TestFirst", nil),
		models.NewMacro("TestSecond", nil),
	}

	want := models.ConfigM{
		Metadata: models.NewDefaultMetadata(),
		Macros:   _macros,
	}

	t.Run("test constructor", func(t *testing.T) {
		got := models.NewConfigM(models.NewDefaultMetadata(), _macros)
		assert.Equal(t, want, got)
	})

	t.Run("Test parser", func(t *testing.T) {
		expectedStr := `Metadata:
    Columns: 2
    SerialPortName: ""
    SerialBaudRate: 9600
    Delay: 125ms
    GUIMode: GUIOnly
Macros:
    - Name: TestFirst
      Actions: []
    - Name: TestSecond
      Actions: []`
		// if err := os.WriteFile("../../tmp/expectedStr.txt", []byte(expectedStr), 0644); err != nil {
		// 	t.Fatal(err)
		// }

		assert.YAMLEq(
			t,
			expectedStr,
			want.String(),
		)
	})

}
