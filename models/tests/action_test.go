package modelstests

import (
	"testing"

	"github.com/ssebs/go-mmp/models"
	"github.com/stretchr/testify/assert"
)

func TestActionModel(t *testing.T) {
	want := models.Action{
		FuncName:  "TestingFunction",
		FuncParam: "TestingParameter",
	}
	t.Run("Test constructor", func(t *testing.T) {
		got := models.NewAction("TestingFunction", "TestingParameter")
		assert.Equal(t, want, got)
	})
	t.Run("Test parser", func(t *testing.T) {
		assert.YAMLEq(
			t,
			"FuncName: TestingFunction\nFuncParam: TestingParameter\n",
			want.String(),
		)
	})

}
