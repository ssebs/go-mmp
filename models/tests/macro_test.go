package models_test

import (
	"testing"

	"github.com/ssebs/go-mmp/models"
	"github.com/stretchr/testify/assert"
)

func TestModelMacroActions(t *testing.T) {
	want := &models.Macro{
		Name: "TEST",
		Actions: []*models.Action{
			models.NewAction("TESTONE", "TESTONE"),
			models.NewAction("TESTTWO", "TESTTWO"),
			models.NewAction("TESTTHREE", "TESTTHREE"),
		},
	}

	t.Run("Test AddAction", func(t *testing.T) {
		got := models.NewMacro("TEST", nil)
		got.AddAction(models.NewAction("TESTONE", "TESTONE"))
		got.AddAction(models.NewAction("TESTTWO", "TESTTWO"))
		got.AddAction(models.NewAction("TESTTHREE", "TESTTHREE"))

		assert.Equal(t, want, got)
	})

	t.Run("Test DeleteAction", func(t *testing.T) {
		got := models.NewMacro("TEST", []*models.Action{
			models.NewAction("TESTONE", "TESTONE"),
			models.NewAction("TESTTWO", "TESTTWO"),
			models.NewAction("TESTTHREE", "TESTTHREE"),
			models.NewAction("DELETE_ME", "DELETE_ME"),
		})
		// Valid delete
		err := got.DeleteAction(3)
		assert.Nil(t, err)
		assert.Equal(t, want, got)

		// Invalid delete
		err = got.DeleteAction(-1)
		assert.Error(t, err)

	})

	t.Run("Test UpdateAction", func(t *testing.T) {
		got := models.NewMacro("TEST", []*models.Action{
			models.NewAction("TESTONE", "TESTONE"),
			models.NewAction("TESTTWO", "TESTTWO"),
			models.NewAction("replace_me", "replace_me"),
		})

		err := got.UpdateAction(2, models.NewAction("TESTTHREE", "TESTTHREE"))
		assert.Nil(t, err)

		assert.Equal(t, want, got)
	})

	t.Run("Test SwapActionPositions", func(t *testing.T) {
		got := models.NewMacro("TEST", []*models.Action{
			models.NewAction("TESTTHREE", "TESTTHREE"),
			models.NewAction("TESTTWO", "TESTTWO"),
			models.NewAction("TESTONE", "TESTONE"),
		})
		assert.NotEqual(t, want, got)

		err := got.SwapActionPositions(0, 2)
		assert.Nil(t, err)

		assert.Equal(t, want, got)
	})

	t.Run("Test GetAction", func(t *testing.T) {
		got, err := want.GetAction(1)
		assert.Nil(t, err)
		assert.Equal(t, want.Actions[1], got)
	})
}

func TestMacroModel(t *testing.T) {
	want := &models.Macro{
		Name:    "TestingName",
		Actions: make([]*models.Action, 0),
	}
	t.Run("Test empty actions constructor", func(t *testing.T) {
		got := models.NewMacro("TestingName", nil)
		assert.Equal(t, want, got)
	})

	t.Run("Test non-empty actions constructor", func(t *testing.T) {
		_actions := []*models.Action{
			{FuncName: "a", FuncParam: "a"},
			{FuncName: "b", FuncParam: "b"},
		}
		want2 := &models.Macro{
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
			&models.Action{FuncName: "a", FuncParam: "aa"},
			&models.Action{FuncName: "b", FuncParam: "bb"},
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
