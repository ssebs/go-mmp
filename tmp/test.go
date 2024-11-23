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

	cm := models.NewConfig(models.NewDefaultMetadata(), []*models.Macro{
		models.NewMacro("Undo", []*models.Action{
			models.NewAction("Shortcut", "CTRL+Z"),
		}),
		models.NewMacro("gg", []*models.Action{
			models.NewAction("PressRelease", "Enter"),
			models.NewAction("Delay", "100ms"),
			models.NewAction("SendText", "gg"),
			models.NewAction("PressRelease", "Enter"),
		}),
		models.NewMacro("close game", []*models.Action{
			models.NewAction("Shortcut", "ALT+F4"),
		}),
	})
	cv := views.NewConfigEditorView(cm.Columns)
	cc := controllers.NewConfigController(cm, cv)

	// cm.UpdateMetadataView()

	win.SetContent(container.NewBorder(
		widget.NewSeparator(),
		widget.NewSeparator(),
		widget.NewSeparator(),
		widget.NewSeparator(),
		cc.ConfigEditorView,
	))

	win.CenterOnScreen()
	win.Resize(fyne.NewSize(300, 500))
	win.ShowAndRun()
}

// test medatata editor

// func main() {
// 	testApp := app.New()
// 	win := testApp.NewWindow("TEST")

// 	mm := models.NewDefaultMetadata()
// 	mv := views.NewMetadataEditorView()
// 	mc := controllers.NewMetadataController(mm, mv)

// 	mc.UpdateMetadataView()

// 	win.SetContent(container.NewBorder(
// 		widget.NewSeparator(),
// 		widget.NewSeparator(),
// 		widget.NewSeparator(),
// 		widget.NewSeparator(),
// 		mv,
// 	))

// 	win.CenterOnScreen()
// 	win.Resize(fyne.NewSize(300, 500))
// 	win.ShowAndRun()

// }

// test MacroEditorView

// func main() {
// 	testApp := app.New()
// 	win := testApp.NewWindow("TEST")

// 	mm := models.NewMacro("TestMacro", []*models.Action{
// 		models.NewAction("PressRelease", "ENTER"),
// 		models.NewAction("Delay", "200ms"),
// 		models.NewAction("SendText", "GG"),
// 		models.NewAction("PressRelease", "ENTER"),
// 	})
// 	mv := views.NewMacroEditorView()
// 	mc := controllers.NewMacroController(mm, mv)
// 	mc.UpdateMacroView()

// 	win.SetContent(container.NewBorder(
// 		widget.NewSeparator(),
// 		widget.NewSeparator(),
// 		widget.NewSeparator(),
// 		widget.NewSeparator(),
// 		mc.MacroEditorView,
// 	))

// 	win.CenterOnScreen()
// 	win.Resize(fyne.NewSize(300, 500))
// 	win.ShowAndRun()
// }

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
