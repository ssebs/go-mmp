package views

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// Ensure interface implementation.
var _ fyne.Widget = (*MacroEditorView)(nil)

type MacroEditorView struct {
	widget.BaseWidget
	macroNameEntry *widget.Entry
	actionsScroll  *container.Scroll
}

func NewMacroEditorView() *MacroEditorView {
	view := &MacroEditorView{
		macroNameEntry: widget.NewEntry(),
		actionsScroll:  container.NewVScroll(container.NewVBox()),
	}
	view.macroNameEntry.Validator = nil
	view.actionsScroll.Resize(view.actionsScroll.Size().AddWidthHeight(0, 400))

	view.ExtendBaseWidget(view)
	return view
}

func (v *MacroEditorView) SetMacroName(s string) {
	v.macroNameEntry.SetText(s)
	v.macroNameEntry.Refresh()
}
func (v *MacroEditorView) SetActions(actions []*ActionItemEditorView) {
	v.actionsScroll.Content.(*fyne.Container).RemoveAll()
	for _, action := range actions {
		v.actionsScroll.Content.(*fyne.Container).Add(action)
	}
	v.actionsScroll.Refresh()
}

func (v *MacroEditorView) CreateRenderer() fyne.WidgetRenderer {
	c := container.NewBorder(
		container.NewVBox(
			widget.NewLabelWithStyle("Edit", fyne.TextAlignCenter,
				fyne.TextStyle{Bold: true},
			),
			widget.NewForm(
				widget.NewFormItem("Name/Title:", v.macroNameEntry),
				widget.NewFormItem("Actions", layout.NewSpacer()),
				widget.NewFormItem("", layout.NewSpacer()),
			),
		),
		container.NewHBox(
			widget.NewButton("Close", func() {
				fmt.Println("CLOSE WINDOW")
			}),
			widget.NewButton("+ Add Action", func() {
				fmt.Println("ADD ACTION")
			}),
			widget.NewButton("Save", func() {
				msg := fmt.Sprintf("Saved %s", v.macroNameEntry.Text)
				fmt.Println(msg)
			}),
		),
		nil, nil,
		v.actionsScroll,
	)
	return widget.NewSimpleRenderer(c)
}
