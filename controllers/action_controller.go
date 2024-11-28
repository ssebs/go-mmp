package controllers

import (
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

	ac.SetValidator(CheckValidParams)

	ac.UpdateActionView()
	return ac
}

func (ac *ActionController) UpdateActionView() {
	ac.ActionItemEditorView.SetAction(ac.Action)
}

func CheckValidParams(a models.Action) error {
	err := a.Validate()
	return err
}
