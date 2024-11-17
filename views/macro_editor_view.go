package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// TODO: implement action drag and drop to position action items
// widget.NewIcon(theme.MenuIcon()),

// Ensure interface implementation.
var _ fyne.Widget = (*MacroEditorView)(nil)

type MacroEditorView struct {
	widget.BaseWidget
	macroNameEntry  *widget.Entry
	actionsScroll   *container.Scroll
	titleLabel      *widget.Label
	addActionBtn    *widget.Button
	saveBtn         *widget.Button
	OnActionDeleted func(idx int)
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

/* Setters */
func (v *MacroEditorView) SetMacroName(s string) {
	v.macroNameEntry.SetText(s)
	v.macroNameEntry.Refresh()
}
func (v *MacroEditorView) SetActions(actions []*ActionItemEditorView) {
	v.actionsScroll.Content.(*fyne.Container).RemoveAll()

	for idx, action := range actions {
		dragIcon := widget.NewIcon(theme.MenuIcon())
		delBtn := widget.NewButtonWithIcon("", theme.NewErrorThemedResource(theme.WindowCloseIcon()), func() {
			if v.OnActionDeleted != nil {
				v.OnActionDeleted(idx)
			}
		})

		v.actionsScroll.Content.(*fyne.Container).Add(container.NewBorder(nil, nil, dragIcon, delBtn, action))
	}
	v.actionsScroll.Refresh()
}

func (v *MacroEditorView) SetOnActionDeleted(f func(idx int)) {
	v.OnActionDeleted = f
}

func (v *MacroEditorView) SetTitleLabel(s string) {
	v.titleLabel.SetText(s)
	v.titleLabel.Refresh()
}

/* Callback overrides */
func (v *MacroEditorView) SetOnMacroNameChanged(f func(string)) {
	v.macroNameEntry.OnChanged = f
}
func (v *MacroEditorView) SetOnAddAction(f func()) {
	v.addActionBtn.OnTapped = f
}
func (v *MacroEditorView) SetOnSave(f func()) {
	v.saveBtn.OnTapped = f
}
