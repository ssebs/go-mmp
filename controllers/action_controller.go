package controllers

import (
	"fmt"

	"github.com/ssebs/go-mmp/models"
	"github.com/ssebs/go-mmp/views"
)

type ActionController struct {
	*models.Action
	*views.ActionItemEditorView
}

func NewActionController(a *models.Action, v *views.ActionItemEditorView) *ActionController {
	ac := &ActionController{
		Action:               a,
		ActionItemEditorView: v,
	}

	ac.ActionItemEditorView.SetOnFuncNameChanged(func(s string) {
		ac.Action.FuncName = s
	})

	ac.ActionItemEditorView.SetOnFuncParamChanged(func(s string) {
		ac.Action.FuncParam = s
	})

	ac.ActionItemEditorView.SetOnDelete(func() {
		fmt.Println("This should be handled in the macro editor")
	})

	return ac
}

func (ac *ActionController) UpdateActionView() {
	ac.ActionItemEditorView.SetAction(ac.Action)
}

func (ac *ActionController) CheckValidParams() bool {
	// TODO: implement
	return false
}
