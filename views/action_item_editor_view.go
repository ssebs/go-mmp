package views

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/ssebs/go-mmp/models"
)

// Ensure interface implementation.
var _ fyne.Widget = (*ActionItemEditorView)(nil)

type ActionItemEditorView struct {
	widget.BaseWidget
	funcSelect     *widget.Select
	funcParamEntry *widget.Entry
	delBtn         *widget.Button
}

func NewActionItemEditorView() *ActionItemEditorView {
	view := &ActionItemEditorView{
		funcSelect: widget.NewSelect(models.GetActionFunctions(),
			func(s string) { fmt.Println("Use SetOnFuncNameChanged() to overwrite this behavior!\ns:", s) },
		),
		funcParamEntry: widget.NewEntry(),
		delBtn: widget.NewButtonWithIcon("", theme.NewErrorThemedResource(theme.WindowCloseIcon()), func() {
			fmt.Println("DELETE")
		}),
	}
	view.funcParamEntry.Validator = nil

	return view
}

func (v *ActionItemEditorView) CreateRenderer() fyne.WidgetRenderer {
	c := container.NewBorder(
		nil, nil,
		v.funcSelect,
		v.delBtn,
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

func (v *ActionItemEditorView) SetOnFuncNameChanged(f func(string)) {
	v.funcSelect.OnChanged = f
}
func (v *ActionItemEditorView) SetOnFuncParamChanged(f func(string)) {
	v.funcParamEntry.OnChanged = f
}
func (v *ActionItemEditorView) SetOnDelete(f func()) {
	v.delBtn.OnTapped = f
}
