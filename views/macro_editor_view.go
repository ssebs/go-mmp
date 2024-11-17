package views

import (
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
	titleLabel     *widget.Label
	addActionBtn   *widget.Button
	saveBtn        *widget.Button
}

func NewMacroEditorView() *MacroEditorView {
	view := &MacroEditorView{
		macroNameEntry: widget.NewEntry(),
		actionsScroll:  container.NewVScroll(container.NewVBox()),
		titleLabel: widget.NewLabelWithStyle("Edit", fyne.TextAlignCenter,
			fyne.TextStyle{Bold: true},
		),
		addActionBtn: widget.NewButton("+ Add Action", nil),
		saveBtn:      widget.NewButton("Save", nil),
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

func (v *MacroEditorView) SetTitleLabel(s string) {
	v.titleLabel.SetText(s)
	v.titleLabel.Refresh()
}

func (v *MacroEditorView) SetOnMacroNameChanged(f func(string)) {
	v.macroNameEntry.OnChanged = f
}

func (v *MacroEditorView) SetOnAddAction(f func()) {
	v.addActionBtn.OnTapped = f
}

func (v *MacroEditorView) SetOnSave(f func()) {
	v.saveBtn.OnTapped = f
}

func (v *MacroEditorView) CreateRenderer() fyne.WidgetRenderer {
	c := container.NewBorder(
		container.NewVBox(
			v.titleLabel,
			widget.NewForm(
				widget.NewFormItem("Name/Title:", v.macroNameEntry),
				widget.NewFormItem("Actions", layout.NewSpacer()),
				widget.NewFormItem("", layout.NewSpacer()),
			),
		),
		container.NewHBox(
			v.addActionBtn,
			v.saveBtn,
		),
		nil, nil,
		v.actionsScroll,
	)
	return widget.NewSimpleRenderer(c)
}
