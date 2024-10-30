package gui

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/ssebs/go-mmp/widgets"
)

// Open a new Window and use it to edit the config
func (g *GUI) EditConfig() {
	editorWindow := g.App.NewWindow("Config Editor")
	g.initGUI(editorWindow)

	// Editor features:
	// - Click btn to edit Macro
	// - Create/Delete Macro btns
	// - Drag and Drop button positions on grid

	editorWindow.CenterOnScreen()
	editorWindow.Show()
}

func (g *GUI) initGUI(win fyne.Window) {
	vbox := container.NewVBox(widget.NewLabelWithStyle("Edit Macros", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}))
	grid := container.NewGridWithColumns(g.config.MacroLayout.SizeX)

	// dragBox := widgets.NewDragBoxWidget("title", g.config, color.RGBA{20, 20, 20, 255}, color.White, func(btnId config.BtnId) {
	// 	fmt.Println("Edit", g.config.Macros[btnId].Name)
	// })
	dragBox := widgets.NewDragBox(g.config, color.RGBA{20, 20, 20, 255}, color.White)

	vbox.Add(dragBox)
	vbox.Add(layout.NewSpacer())
	vbox.Add(grid)

	saveBtn := widget.NewButton("Save", func() {
		fmt.Println("SAVE")
	})
	saveBtn.Importance = widget.HighImportance

	vbox.Add(container.NewHBox(
		widget.NewButton("+ Add Macro", func() {
			fmt.Println("ADD MACRO")
		}),
		saveBtn,
	))
	win.SetContent(vbox)
}

// Save and Load configs, and refresh current running config
func (g *GUI) SaveConfig() {
	fmt.Println("SAVE CONFIG...")
}
func (g *GUI) OpenConfig() {
	fmt.Println("OPEN CONFIG")
}
