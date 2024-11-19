package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Ensure interface implementation.
var _ fyne.Widget = (*ConfigEditorView)(nil)

type ConfigEditorView struct {
	widget.BaseWidget
	titleLabel  *widget.Label
	metadataBtn *widget.Button
}

func NewConfigEditorView() *ConfigEditorView {
	view := &ConfigEditorView{
		titleLabel: widget.NewLabelWithStyle("Edit Config", fyne.TextAlignCenter,
			fyne.TextStyle{Bold: true},
		),
		metadataBtn: widget.NewButton("Edit Metadata", func() {
			win := fyne.CurrentApp().NewWindow("Metadata Editor")
			win.CenterOnScreen()
			win.SetContent(NewMetadataEditorView())
			win.Show()
		}),
	}

	view.ExtendBaseWidget(view)
	return view
}

func (v *ConfigEditorView) CreateRenderer() fyne.WidgetRenderer {
	c := container.NewVBox(v.titleLabel, v.metadataBtn)
	return widget.NewSimpleRenderer(c)
}
