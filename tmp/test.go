package main

import (
	"fmt"

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

	am := models.NewAction("Shortcut", "CTRL+V")
	av := views.NewActionItemEditorView(am)
	ac := controllers.NewActionController(am, av)

	ac.UpdateActionView()

	win.SetContent(container.NewBorder(
		widget.NewSeparator(),
		widget.NewButton("Get Action", func() {
			fmt.Println(ac.Action)
		}),
		widget.NewSeparator(),
		widget.NewSeparator(),
		ac.ActionItemEditorView,
	))

	win.CenterOnScreen()
	win.Resize(fyne.NewSize(300, 100))
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
