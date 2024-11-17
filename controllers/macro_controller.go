package controllers

import (
	"fmt"

	"github.com/ssebs/go-mmp/models"
	"github.com/ssebs/go-mmp/views"
)

type MacroController struct {
	*models.Macro
	*views.MacroEditorView
	actionViews []*views.ActionItemEditorView
}

func NewMacroController(m *models.Macro, v *views.MacroEditorView) *MacroController {
	mc := &MacroController{
		Macro:           m,
		MacroEditorView: v,
		actionViews:     make([]*views.ActionItemEditorView, 0),
	}

	mc.SetOnMacroNameChanged(func(s string) {
		mc.Macro.Name = s
		mc.SetTitleLabel(fmt.Sprintf("Edit %s", mc.Macro.Name))
	})

	return mc
}

func (mc *MacroController) UpdateMacroView() {
	mc.SetMacroName(mc.Name)
	mc.SetTitleLabel(fmt.Sprintf("Edit %s", mc.Macro.Name))
	mc.UpdateActionsInView()
}

func (mc *MacroController) UpdateActionsInView() {
	mc.actionViews = make([]*views.ActionItemEditorView, len(mc.Macro.Actions))
	for i, action := range mc.Macro.Actions {
		mc.actionViews[i] = views.NewActionItemEditorView(action)
	}
	mc.MacroEditorView.SetActions(mc.actionViews)
}
