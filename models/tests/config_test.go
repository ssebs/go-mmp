package models_test

import (
	"testing"

	"github.com/ssebs/go-mmp/models"
	"github.com/stretchr/testify/assert"
)

func TestConfigM(t *testing.T) {
	_macros := []*models.Macro{
		models.NewMacro("TestFirst", nil),
		models.NewMacro("TestSecond", nil),
	}

	want := &models.Config{
		Metadata: models.NewDefaultMetadata(),
		Macros:   _macros,
	}

	t.Run("test constructor", func(t *testing.T) {
		got := models.NewConfig(models.NewDefaultMetadata(), _macros)
		assert.Equal(t, want, got)
	})

	t.Run("Test AddMacro", func(t *testing.T) {
		got := models.NewConfig(models.NewDefaultMetadata(), nil)
		got.AddMacro(models.NewMacro("TestFirst", nil))
		got.AddMacro(models.NewMacro("TestSecond", nil))

		assert.Equal(t, want, got)
	})

	t.Run("Test GetMacro", func(t *testing.T) {
		got, err := want.GetMacro(0)
		assert.Nil(t, err)
		assert.Equal(t, models.NewMacro("TestFirst", nil), got)

		got, err = want.GetMacro(999)
		assert.Error(t, err, "idx out of bounds of Macros")
		assert.Nil(t, got)

	})

	t.Run("Test UpdateMacro", func(t *testing.T) {
		m_copy := make([]*models.Macro, len(_macros))
		copy(m_copy, _macros)

		got := models.NewConfig(models.NewDefaultMetadata(), m_copy)
		got.UpdateMacro(0, models.NewMacro("ReplacedFirst", nil))

		gotMacro, err := got.GetMacro(0)
		assert.Nil(t, err)

		assert.Equal(t, models.NewMacro("ReplacedFirst", nil), gotMacro)
	})

	t.Run("Test DeleteMacro", func(t *testing.T) {
		m_copy := make([]*models.Macro, len(_macros))
		copy(m_copy, _macros)

		got := models.NewConfig(models.NewDefaultMetadata(), m_copy)
		err := got.DeleteMacro(0)
		assert.Nil(t, err)

		gotMacro, err := got.GetMacro(0)
		assert.Nil(t, err)

		assert.Equal(t, _macros[1], gotMacro)

	})
	t.Run("Test SwapMacroPositions", func(t *testing.T) {
		m_copy := make([]*models.Macro, len(_macros))
		copy(m_copy, _macros)
		got := models.NewConfig(models.NewDefaultMetadata(), m_copy)
		err := got.SwapMacroPositions(0, 1)
		assert.Nil(t, err)

		swapped0, err := got.GetMacro(0)
		assert.Nil(t, err)
		assert.Equal(t, _macros[1].Name, swapped0.Name)

		swapped1, err := got.GetMacro(1)
		assert.Nil(t, err)
		assert.Equal(t, _macros[0].Name, swapped1.Name)
	})

	// 	t.Run("Test parser", func(t *testing.T) {
	// 		expectedStr := `Metadata:
	//     Columns: 2
	//     SerialPortName: ""
	//     SerialBaudRate: 9600
	//     Delay: 125ms
	//     GUIMode: GUIOnly
	// Macros:
	//     - Name: TestFirst
	//       Actions: []
	//     - Name: TestSecond
	//       Actions: []`
	// 		// if err := os.WriteFile("../../tmp/expectedStr.txt", []byte(want.String()), 0644); err != nil {
	// 		// 	t.Fatal(err)
	// 		// }

	// 		assert.YAMLEq(
	// 			t,
	// 			expectedStr,
	// 			want.String(),
	// 		)
	// 	})

}
