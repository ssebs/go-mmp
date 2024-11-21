package controllers

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"github.com/ssebs/go-mmp/macro"
	"github.com/ssebs/go-mmp/models"
	"github.com/ssebs/go-mmp/utils"
	"github.com/ssebs/go-mmp/views"
)

type MacroRunnerController struct {
	*models.Config
	*views.MacroRunnerView
	*views.ConfigEditorView
	*ConfigController
	*macro.MacroManager
}

func NewMacroRunnerController(m *models.Config, v *views.MacroRunnerView, mm *macro.MacroManager) *MacroRunnerController {
	cc := &MacroRunnerController{
		Config:           m,
		MacroRunnerView:  v,
		ConfigEditorView: views.NewConfigEditorView(m.Columns),
		MacroManager:     mm,
	}
	cc.ConfigController = NewConfigController(cc.Config, cc.ConfigEditorView)

	cc.MacroRunnerView.SetOnMacroTapped(func(macro *models.Macro) {
		cc.MacroManager.RunMacro(macro)
	})

	cc.MacroRunnerView.SetOnEditConfig(func() {
		win := fyne.CurrentApp().NewWindow("Macro Editor")
		win.CenterOnScreen()
		win.Resize(fyne.NewSize(300, 500))

		win.SetContent(cc.ConfigController.ConfigEditorView)

		win.Show()
		win.SetOnClosed(func() {
			cc.UpdateConfigView()
		})
	})

	cc.MacroRunnerView.SetOnOpenConfig(func() {
		yamlPath, err := utils.GetYAMLFilename(false)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
		}
		fmt.Println("Saving to", yamlPath)

		if err := cc.Config.OpenConfig(yamlPath); err != nil {
			fmt.Fprint(os.Stderr, err)
		}
		cc.ConfigController.UpdateConfigView()
		cc.UpdateConfigView()
	})

	cc.MacroRunnerView.SetOnQuit(func() {
		fyne.CurrentApp().Quit()
	})

	cc.UpdateConfigView()
	return cc
}

func (cc *MacroRunnerController) UpdateConfigView() {
	cc.MacroRunnerView.SetCols(cc.Config.Columns)
	cc.MacroRunnerView.SetMacros(cc.Config.Macros)
}
