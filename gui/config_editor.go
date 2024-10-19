package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/ssebs/go-mmp/config"
)

// Open a new Window and use it to edit the config
func (g *GUI) EditConfig() {
	editorWindow := g.App.NewWindow("Config Editor")
	grid := g.initEditorGrid()

	// Editor features:
	// - Click btn to edit Macro
	// - Create/Delete Macro btns
	// - Drag and Drop button positions on grid

	editorWindow.SetContent(grid)
	editorWindow.CenterOnScreen()
	editorWindow.Show()
}

func (g *GUI) initEditorGrid() *fyne.Container {
	grid := container.NewGridWithColumns(g.config.MacroLayout.SizeX)

	for pos := 1; pos <= len(g.config.Macros); pos++ {
		macro := g.config.Macros[config.BtnId(pos)]

		grid.Add(widget.NewButton(macro.Name, func() {
			g.macroManager.RunActionFromID(config.BtnId(pos))
		}))
	}
	return grid
}

// Save and Load configs, and refresh current running config
func (g *GUI) SaveConfig() {
	fmt.Println("SAVE CONFIG...")
}
func (g *GUI) OpenConfig() {
	fmt.Println("OPEN CONFIG")
}
