package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/ssebs/go-mmp/controllers"
	"github.com/ssebs/go-mmp/models"
	"github.com/ssebs/go-mmp/views"
)

func main() {
	testApp := app.New()
	win := testApp.NewWindow("TEST")

	mm := models.NewMacro("TestMacro", []*models.Action{
		models.NewAction("Shortcut", "CTRL+C"),
		models.NewAction("Delay", "200ms"),
		models.NewAction("Shortcut", "CTRL+V"),
	})
	mv := views.NewMacroEditorView()
	mc := controllers.NewMacroController(mm, mv)
	mc.UpdateMacroView()

	win.SetContent(container.NewBorder(
		widget.NewSeparator(),
		widget.NewSeparator(),
		widget.NewSeparator(),
		widget.NewSeparator(),
		mc.MacroEditorView,
	))

	win.CenterOnScreen()
	win.Resize(fyne.NewSize(300, 300))
	win.ShowAndRun()
}

// func main() {
// 	cliFlags := config.ParseFlags()
// 	conf, err := config.NewConfig(cliFlags)
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, err.Error())
// 	}

// 	macroMgr, err := macro.NewMacroManager(conf)
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, err.Error())
// 	}

// 	if cliFlags.GUIMode != models.NOTSET {
// 		macroMgr.Config.GUIMode = cliFlags.GUIMode
// 	}

// 	g := gui.NewGUI(macroMgr)

// 	g.EditConfig()
// 	g.App.Run()
// }
