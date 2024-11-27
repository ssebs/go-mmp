package views

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/ssebs/go-mmp/models"
)

// Ensure interface implementation.
var _ fyne.Widget = (*ActionItemEditorView)(nil)

type ActionItemEditorView struct {
	widget.BaseWidget
	funcSelect     *widget.Select
	funcParamEntry *widget.Entry
}

func NewActionItemEditorView(initialAction *models.Action) *ActionItemEditorView {
	view := &ActionItemEditorView{
		funcSelect:     widget.NewSelect(models.GetActionFunctions(), nil),
		funcParamEntry: widget.NewEntry(),
	}
	view.funcParamEntry.Validator = nil

	if initialAction.FuncName != "" {
		view.funcSelect.SetSelected(initialAction.FuncName)
	}
	if initialAction.FuncParam != "" {
		view.funcParamEntry.SetText(initialAction.FuncParam)
	}

	view.ExtendBaseWidget(view)
	return view
}

func (v *ActionItemEditorView) CreateRenderer() fyne.WidgetRenderer {
	c := container.NewBorder(
		nil, nil,
		v.funcSelect,
		nil,
		v.funcParamEntry,
	)

	return widget.NewSimpleRenderer(c)
}

func (v *ActionItemEditorView) SetAction(a *models.Action) {
	v.funcParamEntry.SetText(a.FuncParam)
	v.funcParamEntry.Refresh()

	v.funcSelect.SetSelected(a.FuncName)
	v.funcSelect.Refresh()
}

func (v *ActionItemEditorView) GetAction() models.Action {
	return models.Action{
		FuncName:  v.funcSelect.Selected,
		FuncParam: v.funcParamEntry.Text,
	}
}

func (v *ActionItemEditorView) SetValidator(f func(models.Action) error) {
	v.funcParamEntry.Validator = func(s string) error {
		err := f(v.GetAction())
		fmt.Println(err)
		return err
	}
}

func (v *ActionItemEditorView) SetOnFuncNameChanged(f func(string)) {
	v.funcSelect.OnChanged = f
}
func (v *ActionItemEditorView) SetOnFuncParamChanged(f func(string)) {
	v.funcParamEntry.OnChanged = f
}
