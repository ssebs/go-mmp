package controllers

import (
	"fmt"
	"os"

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
		fmt.Println("\"Saving\" Macro")
		// fmt.Println(mc.Macro)
		mc.MacroEditorView.CloseWindow()
	})

	mc.SetOnActionDeleted(func(idx int) {
		// TODO: allow cancel / undo. add delete to stack?
		if err := mc.Macro.DeleteAction(idx); err != nil {
			fmt.Fprintln(os.Stderr, "failed to delete action", err)
		}
		mc.UpdateActionsInView()
		fmt.Printf("Deleted %d", idx)
	})

	mc.SetOnActionsSwapped(func(idx1, idx2 int) {
		if err := mc.Macro.MoveActionPositions(idx1, idx2); err != nil {
			fmt.Fprintln(os.Stderr, "failed to move actions", err)
		}
		mc.UpdateActionsInView()
		fmt.Printf("Moved Actions %d and %d\n", idx1, idx2)
	})

	mc.UpdateMacroView()
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
