package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Ensure interface implementation.
var _ fyne.Widget = (*MacroEditorView)(nil)

type MacroEditorView struct {
	widget.BaseWidget
	macroNameEntry   *widget.Entry
	actionsScroll    *container.Scroll
	titleLabel       *widget.Label
	addActionBtn     *widget.Button
	saveBtn          *widget.Button
	rootWin          fyne.Window
	OnActionDeleted  func(idx int)
	OnActionsSwapped func(idx1, idx2 int)
}

func NewMacroEditorView(rootWin fyne.Window) *MacroEditorView {
	view := &MacroEditorView{
		macroNameEntry: widget.NewEntry(),
		actionsScroll:  container.NewVScroll(NewDragAndDropView(container.NewVBox(), DRAG_VERTICAL)),
		titleLabel: widget.NewLabelWithStyle("Edit", fyne.TextAlignCenter,
			fyne.TextStyle{Bold: true},
		),
		addActionBtn: widget.NewButton("+ Add Action", nil),
		saveBtn:      widget.NewButton("Save and Close", nil),
		rootWin:      rootWin,
	}
	view.saveBtn.Importance = widget.HighImportance
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
			layout.NewSpacer(),
			v.addActionBtn,
			v.saveBtn,
			layout.NewSpacer(),
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
	var stuff []fyne.CanvasObject

	for idx, action := range actions {
		delBtn := widget.NewButtonWithIcon("", theme.NewErrorThemedResource(theme.WindowCloseIcon()), func() {
			if v.OnActionDeleted != nil {
				v.OnActionDeleted(idx)
			}
		})
		stuff = append(stuff, container.NewBorder(nil, nil, nil, delBtn, action))
	}
	v.actionsScroll.Content.(*DragAndDropView).SetDragItems(stuff)
	v.actionsScroll.Content.(*DragAndDropView).SetOnItemsSwapped(v.OnActionsSwapped)
	v.actionsScroll.Refresh()
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
func (v *MacroEditorView) SetOnActionDeleted(f func(idx int)) {
	v.OnActionDeleted = f
}
func (v *MacroEditorView) SetOnActionsSwapped(f func(idx1, idx2 int)) {
	v.OnActionsSwapped = f
}
func (v *MacroEditorView) CloseWindow() {
	v.rootWin.Close()
}
