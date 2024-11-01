package gui

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	osdialog "github.com/sqweek/dialog"
	"github.com/ssebs/go-mmp/widgets"
)

// Open a new Window and use it to edit the config
func (g *GUI) EditConfig() {
	editorWindow := g.App.NewWindow("Config Editor")
	g.initGUI(editorWindow)

	// Editor features:
	// [ ] Click btn to edit Macro
	// [ ] Create/Delete Macro btns
	// [x] Drag and Drop button positions on grid

	editorWindow.CenterOnScreen()
	editorWindow.Show()
}

func (g *GUI) initGUI(win fyne.Window) {
	vbox := container.NewVBox(widget.NewLabelWithStyle("Edit Macros", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}))
	dragBox := widgets.NewDragBox(g.config, color.RGBA{20, 20, 20, 255}, color.White)

	vbox.Add(dragBox)
	vbox.Add(layout.NewSpacer())

	saveBtn := widget.NewButton("Save", func() {
		g.config.SaveConfig("")
	})
	saveBtn.Importance = widget.HighImportance

	vbox.Add(container.NewHBox(
		widget.NewButton("Open Config", func() {
			filename, err := osdialog.File().Filter("YAML config file", "yaml", "yml").Load()
			if err != nil {
				ShowErrorDialogAndRun(err)
			}
			g.config.OpenConfig(filename)
			g.SetContent(widget.NewLabel("test"))
			g.initGUI(win)
		}),
		widget.NewButton("+ Add Macro", func() {
			fmt.Println("ADD MACRO")
		}),
		saveBtn,
		widget.NewButton("Save As", func() {
			filename, err := osdialog.File().Filter("YAML config file", "yaml", "yml").Load()
			if err != nil {
				ShowErrorDialogAndRun(err)
			}
			g.config.SaveConfig(filename)
		}),
	))
	win.SetContent(vbox)
}

func (g *GUI) OpenConfig() {
	fmt.Println("OPEN CONFIG")
}
