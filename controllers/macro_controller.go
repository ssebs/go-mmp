package controllers

import (
	"fmt"

	"github.com/ssebs/go-mmp/models"
	"github.com/ssebs/go-mmp/views"
)

type MacroController struct {
	*models.Macro
	*views.MacroEditorView
	actionControllers []*ActionController
}

func NewMacroController(m *models.Macro, v *views.MacroEditorView) *MacroController {
	mc := &MacroController{
		Macro:             m,
		MacroEditorView:   v,
		actionControllers: make([]*ActionController, 0),
	}

	mc.SetOnMacroNameChanged(func(s string) {
		mc.Macro.Name = s
		mc.SetTitleLabel(fmt.Sprintf("Edit %s", mc.Macro.Name))
	})

	mc.SetOnAddAction(func() {
		mc.Macro.AddAction(models.NewDefaultAction())
		mc.UpdateActionsInView()
	})
	mc.SetOnSave(func() {
		fmt.Println("Saving Macro!")
		fmt.Println(mc.Macro)
	})

	// Set on delete

	return mc
}

func (mc *MacroController) UpdateMacroView() {
	mc.SetMacroName(mc.Name)
	mc.SetTitleLabel(fmt.Sprintf("Edit %s", mc.Macro.Name))
	mc.UpdateActionsInView()
}

func (mc *MacroController) UpdateActionsInView() {
	actionViews := make([]*views.ActionItemEditorView, 0, len(mc.Macro.Actions))

	for _, action := range mc.Macro.Actions {
		av := views.NewActionItemEditorView(action)
		ac := NewActionController(action, av)

		mc.actionControllers = append(mc.actionControllers, ac)
		actionViews = append(actionViews, av)
	}

	mc.MacroEditorView.SetActions(actionViews)
}
