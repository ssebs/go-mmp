package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/ssebs/go-mmp/models"
)

// Ensure interface implementation.
var _ fyne.Widget = (*ConfigEditorView)(nil)

type ConfigEditorView struct {
	widget.BaseWidget
	cols            int
	titleLabel      *widget.Label
	metadataBtn     *widget.Button
	macrosContainer *DragAndDropView
	OnMacroTapped   func(m *models.Macro)
	OnMacrosSwapped func(idx1, idx2 int)
	OnMacroDeleted  func(idx1 int)
	addMacroBtn     *widget.Button
	saveBtn         *widget.Button
	saveAsBtn       *widget.Button
}

func NewConfigEditorView(cols int) *ConfigEditorView {
	view := &ConfigEditorView{
		cols: cols,
		titleLabel: widget.NewLabelWithStyle("Edit Config", fyne.TextAlignCenter,
			fyne.TextStyle{Bold: true},
		),
		metadataBtn:     widget.NewButton("Edit Metadata", nil),
		macrosContainer: NewDragAndDropView(container.NewGridWithColumns(cols), DRAG_BOTH),
		addMacroBtn:     widget.NewButton("+ Add Macro", nil),
		saveAsBtn:       widget.NewButton("Save as", nil),
		saveBtn:         widget.NewButton("Save", nil),
	}

	view.ExtendBaseWidget(view)
	return view
}

// SEE: MacroEditorView.SetActions(...)
func (v *ConfigEditorView) SetMacros(macros []*models.Macro) {
	var stuff []fyne.CanvasObject

	for idx, macro := range macros {
		delBtn := widget.NewButtonWithIcon("", theme.NewErrorThemedResource(theme.WindowCloseIcon()), func() {
			if v.OnMacroDeleted != nil {
				v.OnMacroDeleted(idx)
			}
		})
		macroBtn := widget.NewButton(macro.Name, func() {
			v.OnMacroTapped(macro)
		})
		stuff = append(stuff, container.NewBorder(nil, nil, nil, delBtn, macroBtn))
	}
	v.macrosContainer.SetDragItems(stuff)
	v.macrosContainer.SetOnItemsSwapped(v.OnMacrosSwapped)
	v.macrosContainer.Refresh()
}
func (v *ConfigEditorView) SetOnMacroDeleted(f func(int)) {
	v.OnMacroDeleted = f
}
func (v *ConfigEditorView) SetOnMacroTapped(f func(*models.Macro)) {
	v.OnMacroTapped = f
}
func (v *ConfigEditorView) SetOnMetadataTapped(f func()) {
	v.metadataBtn.OnTapped = f
}
func (v *ConfigEditorView) SetOnMacrosSwapped(f func(idx1, idx2 int)) {
	v.OnMacrosSwapped = f
}
func (v *ConfigEditorView) SetOnAddMacro(f func()) {
	v.addMacroBtn.OnTapped = f
}
func (v *ConfigEditorView) SetOnSave(f func()) {
	v.saveBtn.OnTapped = f
}
func (v *ConfigEditorView) SetOnSaveAs(f func()) {
	v.saveAsBtn.OnTapped = f
}
func (v *ConfigEditorView) SetCols(c int) {
	v.cols = c
	v.macrosContainer.dragItems.Layout = (container.NewGridWithColumns(c)).Layout
	v.macrosContainer.dragItems.Refresh()
}

func (v *ConfigEditorView) CreateRenderer() fyne.WidgetRenderer {
	c := container.NewVBox(
		container.NewVBox(v.titleLabel, v.metadataBtn),
		v.macrosContainer,
		container.NewHBox(v.addMacroBtn, v.saveAsBtn, v.saveBtn),
	)
	return widget.NewSimpleRenderer(c)
}
