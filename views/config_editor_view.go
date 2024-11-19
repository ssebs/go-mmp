package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
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
}

func NewConfigEditorView(cols int) *ConfigEditorView {
	view := &ConfigEditorView{
		cols: cols,
		titleLabel: widget.NewLabelWithStyle("Edit Config", fyne.TextAlignCenter,
			fyne.TextStyle{Bold: true},
		),
		metadataBtn:     widget.NewButton("Edit Metadata", nil),
		macrosContainer: NewDragAndDropView(container.NewGridWithColumns(cols)),
	}

	view.ExtendBaseWidget(view)
	return view
}

// SEE: MacroEditorView.SetActions(...)
func (v *ConfigEditorView) SetMacros(macros []*models.Macro) {
	var stuff []fyne.CanvasObject

	for _, macro := range macros {
		stuff = append(stuff, widget.NewButton(macro.Name, func() {
			v.OnMacroTapped(macro)
		}))
		// ALSO ALLOW X+Y DRAGGING ON DRAGANDDROPVIEW
	}
	v.macrosContainer.SetDragItems(stuff)
	v.macrosContainer.SetOnItemsSwapped(v.OnMacrosSwapped)
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

func (v *ConfigEditorView) SetCols(c int) {
	v.cols = c
}

func (v *ConfigEditorView) CreateRenderer() fyne.WidgetRenderer {
	c := container.NewVBox(v.titleLabel, v.metadataBtn, v.macrosContainer)
	return widget.NewSimpleRenderer(c)
}
