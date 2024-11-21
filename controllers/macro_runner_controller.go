package controllers

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"github.com/ssebs/go-mmp/models"
	"github.com/ssebs/go-mmp/utils"
	"github.com/ssebs/go-mmp/views"
)

type MacroRunnerController struct {
	*models.ConfigM
	*views.MacroRunnerView
	*views.ConfigEditorView
	*ConfigController
}

func NewMacroRunnerController(m *models.ConfigM, v *views.MacroRunnerView) *MacroRunnerController {
	cc := &MacroRunnerController{
		ConfigM:          m,
		MacroRunnerView:  v,
		ConfigEditorView: views.NewConfigEditorView(m.Columns),
	}
	cc.ConfigController = NewConfigController(cc.ConfigM, cc.ConfigEditorView)

	cc.MacroRunnerView.SetOnMacroTapped(func(macro *models.Macro) {

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

		if err := cc.ConfigM.OpenConfig(yamlPath); err != nil {
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
	cc.MacroRunnerView.SetCols(cc.ConfigM.Columns)
	cc.MacroRunnerView.SetMacros(cc.ConfigM.Macros)
}
