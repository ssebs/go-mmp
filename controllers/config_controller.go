package controllers

import (
	"fmt"
	"os"

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

	cc.ConfigEditorView.SetOnMetadataTapped(func() {
		win := fyne.CurrentApp().NewWindow("Metadata Editor")
		win.CenterOnScreen()
		win.SetContent(cc.metaController.MetadataEditorView)
		win.Resize(fyne.NewSize(300, 500))
		win.Show()
	})

	cc.ConfigEditorView.SetOnMacrosSwapped(func(idx1, idx2 int) {
		fmt.Println("Swapping macros", idx1, idx2)
		if err := cc.ConfigM.SwapMacroPositions(idx1, idx2); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	})

	cc.ConfigEditorView.SetOnMacroTapped(func(macro *models.Macro) {
		win := fyne.CurrentApp().NewWindow("Macro Editor")
		win.CenterOnScreen()

		mv := views.NewMacroEditorView()
		mc := NewMacroController(macro, mv)
		win.SetContent(mc.MacroEditorView)
		win.Resize(fyne.NewSize(300, 500))
		win.Show()
		win.SetOnClosed(func() {
			cc.UpdateConfigView()
		})

	})

	cc.UpdateConfigView()
	return cc
}

func (cc *ConfigController) UpdateConfigView() {
	cc.SetCols(cc.ConfigM.Columns)
	cc.ConfigEditorView.SetMacros(cc.ConfigM.Macros)
}
