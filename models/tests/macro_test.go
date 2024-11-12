package models_test

import (
	"testing"

	"github.com/ssebs/go-mmp/models"
	"github.com/stretchr/testify/assert"
)

func TestMacroModel(t *testing.T) {
	want := models.Macro{
		Name:    "TestingName",
		Actions: make([]models.Action, 0),
	}
	t.Run("Test empty actions constructor", func(t *testing.T) {
		got := models.NewMacro("TestingName", nil)
		assert.Equal(t, want, got)
	})

	t.Run("Test non-empty actions constructor", func(t *testing.T) {
		_actions := []models.Action{
			{FuncName: "a", FuncParam: "a"},
			{FuncName: "b", FuncParam: "b"},
		}
		want2 := models.Macro{
			Name:    "TestingName",
			Actions: _actions,
		}

		// First test with same array
		got := models.NewMacro("TestingName", _actions)
		assert.Equal(t, want2, got)

		// Next test appending the values
		got = models.NewMacro("TestingName", nil)
		got.Actions = append(got.Actions, _actions...)
		assert.Equal(t, want2, got)

	})

	t.Run("Test parser", func(t *testing.T) {
		assert.YAMLEq(
			t,
			"Name: TestingName\nActions: []\n",
			want.String(),
		)

		want.Actions = append(want.Actions,
			models.Action{FuncName: "a", FuncParam: "aa"},
			models.Action{FuncName: "b", FuncParam: "bb"},
		)

		expectedStr := `Name: TestingName
Actions:
    - FuncName: a
      FuncParam: aa
    - FuncName: b
      FuncParam: bb`

		assert.YAMLEq(
			t,
			expectedStr,
			want.String(),
		)
	})
}
