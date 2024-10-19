package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/ssebs/go-mmp/config"
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

	for pos := 1; pos <= len(g.config.Macros); pos++ {
		macro := g.config.Macros[config.BtnId(pos)]

		cont := container.NewBorder(
			// Top right 'x' button
			container.NewHBox(
				widget.NewLabel(macro.Name),
				layout.NewSpacer(),
				// TODO: maybe I shouldn't use a lambda here...
				func(callback func()) *widget.Button {
					xbtn := widget.NewButton("x", callback)
					xbtn.Importance = widget.DangerImportance
					return xbtn
				}(func() {
					fmt.Println("Del ", macro.Name)
				}),
			),
			// Bottom Edit btn
			container.NewHBox(
				layout.NewSpacer(),
				widget.NewButton("Edit", func() {
					fmt.Println("Edit actions")
				}),
				layout.NewSpacer(),
			),
			nil,
			nil,
			// Main content

		)
		cont.Resize(fyne.NewSize(80, 80))

		grid.Add(cont)
	}
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
