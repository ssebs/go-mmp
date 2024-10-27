package gui

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
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

	dragBox := NewDragBoxWidget("title", g.config, color.RGBA{20, 20, 20, 255}, color.White, func() {
		fmt.Println("Edit button pressed")
	})

	vbox.Add(dragBox)
	vbox.Add(layout.NewSpacer())

	// for pos := 1; pos <= len(g.config.Macros); pos++ {
	// 	macro := g.config.Macros[config.BtnId(pos)]

	// 	cont := container.NewBorder(
	// 		// Top right 'x' button
	// 		container.NewHBox(
	// 			widget.NewLabelWithStyle(macro.Name, fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
	// 			layout.NewSpacer(),
	// 			// TODO: maybe I shouldn't use a lambda here...
	// 			func(callback func()) *widget.Button {
	// 				xbtn := widget.NewButton("x", callback)
	// 				xbtn.Importance = widget.DangerImportance
	// 				return xbtn
	// 			}(func() {
	// 				fmt.Println("Del ", macro.Name)
	// 			}),
	// 		),
	// 		// Bottom Edit btn
	// 		container.NewHBox(
	// 			layout.NewSpacer(),
	// 			widget.NewButton("Edit", func() {
	// 				fmt.Println("Edit actions")
	// 			}),
	// 			layout.NewSpacer(),
	// 		),
	// 		nil,
	// 		nil,
	// 		// Main content
	// 	)
	// 	cont.Resize(fyne.NewSize(80, 80))
	// 	// contBorder := container.NewBorder(widget.NewSeparator(), widget.NewSeparator(), widget.NewSeparator(), widget.NewSeparator(), cont)
	// 	contBorder := container.NewStack(
	// 		func() *canvas.Rectangle {
	// 			r := canvas.NewRectangle(color.Black)
	// 			r.Resize(fyne.NewSize(90, 90))
	// 			// TODO: MAKE THIS CLICKABLE
	// 			// TODO: Make colored button widget
	// 			return r
	// 		}(),
	// 		cont,
	// 	)

	// 	grid.Add(contBorder)
	// }
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
