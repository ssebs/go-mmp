package controllers

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"github.com/ssebs/go-mmp/models"
	"github.com/ssebs/go-mmp/utils"
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
		win.Resize(fyne.NewSize(300, 500))

		win.SetContent(cc.metaController.MetadataEditorView)

		win.Show()
		win.SetOnClosed(func() {
			fmt.Println("cols", cc.ConfigM.Columns)
			cc.UpdateConfigView()
		})
	})

	cc.ConfigEditorView.SetOnMacrosSwapped(func(idx1, idx2 int) {
		if err := cc.ConfigM.SwapMacroPositions(idx1, idx2); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		cc.UpdateConfigView()
	})

	cc.ConfigEditorView.SetOnMacroTapped(func(macro *models.Macro) {
		win := fyne.CurrentApp().NewWindow("Macro Editor")
		win.CenterOnScreen()
		win.Resize(fyne.NewSize(300, 500))

		mv := views.NewMacroEditorView()
		mc := NewMacroController(macro, mv)
		win.SetContent(mc.MacroEditorView)

		win.Show()
		win.SetOnClosed(func() {
			cc.UpdateConfigView()
		})
	})

	cc.ConfigEditorView.SetOnAddMacro(func() {
		cc.ConfigM.AddMacro(models.NewDefaultMacro())
		cc.UpdateConfigView()
	})

	cc.ConfigEditorView.SetOnSave(func() {
		fmt.Println("Saving")
		fmt.Println(cc.ConfigM)
		// cc.ConfigM.Save("")
	})
	cc.ConfigEditorView.SetOnSaveAs(func() {
		yamlPath, err := utils.GetYAMLFilename(true)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
		}
		fmt.Println("Saving to", yamlPath)
		fmt.Println(cc.ConfigM)
		// cc.ConfigM.Save(yamlPath)
	})

	cc.ConfigEditorView.SetOnMacroDeleted(func(i int) {
		if err := cc.ConfigM.DeleteMacro(i); err != nil {
			fmt.Fprint(os.Stderr, err)
		}
		cc.UpdateConfigView()
	})

	cc.UpdateConfigView()
	return cc
}

func (cc *ConfigController) UpdateConfigView() {
	cc.SetCols(cc.ConfigM.Columns)
	cc.ConfigEditorView.SetMacros(cc.ConfigM.Macros)
}
