package modelstests

import (
	"testing"

	"github.com/ssebs/go-mmp/models"
	"github.com/stretchr/testify/assert"
)

func TestMacroModel(t *testing.T) {
	t.Run("Test empty actions constructor", func(t *testing.T) {
		want := models.Macro{
			Name:    "TestingName",
			Actions: make([]models.Action, 0),
		}
		got := models.NewMacro("TestingName", nil)
		assert.Equal(t, want, got)
	})

	t.Run("Test non-empty actions constructor", func(t *testing.T) {
		_actions := []models.Action{
			{FuncName: "a", FuncParam: "a"},
			{FuncName: "b", FuncParam: "b"},
		}
		want := models.Macro{
			Name:    "TestingName",
			Actions: _actions,
		}

		// First test with same array
		got := models.NewMacro("TestingName", _actions)
		assert.Equal(t, want, got)

		// Next test appending the values
		got = models.NewMacro("TestingName", nil)
		got.Actions = append(got.Actions, _actions...)
		assert.Equal(t, want, got)

	})

}
