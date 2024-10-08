package gui

import (
	"fmt"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/ssebs/go-mmp/config"
)

type LayoutEditor struct {
	Layout *config.MacroLayout
}
type ActionEditor struct {
	Title   string
	Actions []*config.Macro
}

type ConfigEditor struct {
	LayoutEditor
	ActionEditor
}

// Open a new Window and use it to edit the config
func (g *GUI) EditConfig() {
	fmt.Println("EDIT CONFIG")

	// Create window
	editor := g.App.NewWindow("Config Editor")
	grid := container.NewGridWithColumns(g.config.MacroLayout.SizeX)

	for btnId, macro := range g.config.Macros {
		// TODO: sort this properly
		grid.Add(widget.NewButton(
			macro.Name,
			func() { g.macroManager.RunActionFromID(btnId - 1) },
		),
		)
	}

	editor.SetContent(grid)
	editor.CenterOnScreen()
	// Fill with layout contents

	// Each item will be editable

	editor.Show()
}

// Save and Load configs, and refresh current running config
func (g *GUI) SaveConfig() {
	fmt.Println("SAVE CONFIG...")
}
func (g *GUI) OpenConfig() {
	fmt.Println("OPEN CONFIG")
}
