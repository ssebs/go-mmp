package controllers

import (
	"fyne.io/fyne/v2"
	"github.com/ssebs/go-mmp/models"
	"github.com/ssebs/go-mmp/views"
)

type ConfigController struct {
	*models.ConfigM
	*views.ConfigEditorView
	metaController *MetadataController
}

func NewConfigController(m *models.ConfigM, v *views.ConfigEditorView) *ConfigController {
	cc := &ConfigController{
		ConfigM:          m,
		ConfigEditorView: v,
		metaController:   NewMetadataController(m.Metadata, views.NewMetadataEditorView()),
	}

	cc.SetOnMetadataTapped(func() {
		win := fyne.CurrentApp().NewWindow("Metadata Editor")
		win.CenterOnScreen()
		win.SetContent(cc.metaController.MetadataEditorView)
		win.Show()
	})

	return cc
}

func (cc *ConfigController) UpdateConfigView() {
	// cc.ConfigEditorView.
}
