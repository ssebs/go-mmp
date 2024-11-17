package controllers

import (
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

	return mc
}

func (mc *MacroController) UpdateMacroView() {
	mc.SetMacroName(mc.Name)
	mc.UpdateActionsInView()
	mc.Refresh()
}

func (mc *MacroController) UpdateActionsInView() {
	mc.actionViews = make([]*views.ActionItemEditorView, len(mc.Macro.Actions))
	for i, action := range mc.Macro.Actions {
		mc.actionViews[i] = views.NewActionItemEditorView(action)
	}
	mc.MacroEditorView.SetActions(mc.actionViews)
	mc.MacroEditorView.Refresh()
}
