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
	*models.Config
	*views.ConfigEditorView
	*MetadataController
}

func NewConfigController(m *models.Config, v *views.ConfigEditorView) *ConfigController {
	cc := &ConfigController{
		Config:             m,
		ConfigEditorView:   v,
		MetadataController: NewMetadataController(m.Metadata, views.NewMetadataEditorView()),
	}
	cc.MetadataController.SetOnSubmit(func(m models.Metadata) {
		cc.MetadataController.UpdateAllFields(m)
		fmt.Println("Updated metadata")
		// fmt.Println(mc.Metadata)
		cc.UpdateConfigView()
	})
	cc.ConfigEditorView.SetMetadataView(cc.MetadataEditorView)

	cc.ConfigEditorView.SetOnMacrosSwapped(func(idx1, idx2 int) {
		if err := cc.Config.SwapMacroPositions(idx1, idx2); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		cc.UpdateConfigView()
	})

	cc.ConfigEditorView.SetOnMacroTapped(func(macro *models.Macro) {
		win := fyne.CurrentApp().NewWindow("Macro Editor")
		win.CenterOnScreen()
		win.Resize(fyne.NewSize(300, 500))

		mv := views.NewMacroEditorView(win)
		mc := NewMacroController(macro, mv)
		win.SetContent(mc.MacroEditorView)

		win.Show()
		win.SetOnClosed(func() {
			cc.UpdateConfigView()
		})
	})

	cc.ConfigEditorView.SetOnAddMacro(func() {
		cc.Config.AddMacro(models.NewDefaultMacro())
		cc.UpdateConfigView()
	})

	cc.ConfigEditorView.SetOnSave(func() {
		fmt.Println("Saving")
		// fmt.Println(cc.Config)
		if err := cc.Config.SaveConfig(""); err != nil {
			fmt.Fprint(os.Stderr, err)
		}
		cc.ConfigEditorView.CloseWindow()
	})

	cc.ConfigEditorView.SetOnSaveAs(func() {
		yamlPath, err := utils.GetYAMLFilename(true)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
		}
		fmt.Println("Saving to", yamlPath)
		// fmt.Println(cc.Config)
		if err := cc.Config.SaveConfig(yamlPath); err != nil {
			fmt.Fprint(os.Stderr, err)
		}
		cc.ConfigEditorView.CloseWindow()
	})

	cc.ConfigEditorView.SetOnMacroDeleted(func(i int) {
		if err := cc.Config.DeleteMacro(i); err != nil {
			fmt.Fprint(os.Stderr, err)
		}
		cc.UpdateConfigView()
	})

	cc.UpdateConfigView()
	return cc
}

func (cc *ConfigController) UpdateConfigView() {
	cc.SetCols(cc.Config.Columns)
	cc.ConfigEditorView.SetMacros(cc.Config.Macros)
}

func (cc *ConfigController) SetRootWin(rootWin fyne.Window) {
	cc.ConfigEditorView.SetRootWin(rootWin)
}
