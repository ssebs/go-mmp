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
	got := models.NewAction("TestingFunction", "TestingParameter")

	assert.Equal(t, want, got)
	assert.YAMLEq(
		t,
		"FuncName: TestingFunction\nFuncParam: TestingParameter\n",
		got.String(),
	)

}
