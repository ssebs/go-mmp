package views

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/ssebs/go-mmp/models"
	"github.com/ssebs/go-mmp/utils"
)

// TODO: Add validation!

// Ensure interface implementation.
var _ fyne.Widget = (*MetadataEditorView)(nil)

type MetadataEditorView struct {
	widget.BaseWidget
	form                *widget.Form
	colsEntry           *widget.Entry
	serialPortNameEntry *widget.SelectEntry
	serialPortBaudEntry *widget.Entry
	delayEntry          *widget.Entry
	guiModeSelect       *widget.Select
}

func NewMetadataEditorView() *MetadataEditorView {
	view := &MetadataEditorView{
		colsEntry:           widget.NewEntry(),
		serialPortNameEntry: widget.NewSelectEntry(models.GetSerialPortsList()),
		serialPortBaudEntry: widget.NewEntry(),
		delayEntry:          widget.NewEntry(),
		guiModeSelect:       widget.NewSelect(models.GetGUIModesList(), nil),
		form:                widget.NewForm(),
	}
	view.form.SubmitText = "Save Metadata"

	view.form.Append("Grid columns", view.colsEntry)
	view.form.Append("Serial Port Name", view.serialPortNameEntry)
	view.form.Append("Serial Port Baud Rate", view.serialPortBaudEntry)
	view.form.Append("Default Delay", view.delayEntry)
	view.form.Append("GUI Mode", view.guiModeSelect)

	view.ExtendBaseWidget(view)
	return view
}

func (v *MetadataEditorView) SetMetadata(m *models.Metadata) {
	v.colsEntry.SetText(fmt.Sprintf("%d", m.Columns))
	v.serialPortBaudEntry.SetText(fmt.Sprintf("%d", m.SerialBaudRate))
	v.serialPortNameEntry.SetText(m.SerialPortName)
	v.delayEntry.SetText(m.Delay.String())
	v.guiModeSelect.SetSelected(m.GUIMode.String())
	v.Refresh()
}
func (v *MetadataEditorView) SetOnSubmit(f func(models.Metadata)) {
	v.form.OnSubmit = func() {
		f(v.GetMetadata())
	}
	v.form.Refresh()
}

func (v *MetadataEditorView) GetMetadata() models.Metadata {
	cols, _ := utils.StringToInt(v.colsEntry.Text)
	baud, _ := utils.StringToInt(v.serialPortBaudEntry.Text)
	delay, _ := time.ParseDuration(v.delayEntry.Text)
	guimode, _ := models.ParseGUIModeString(v.guiModeSelect.Selected)

	return models.Metadata{
		Columns:        cols,
		SerialPortName: v.serialPortNameEntry.Text,
		SerialBaudRate: baud,
		Delay:          delay,
		GUIMode:        guimode,
	}
}

func (v *MetadataEditorView) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(v.form)
}
