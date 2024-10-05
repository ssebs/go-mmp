package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
)

func (g *GUI) initMenu() {
	var fileMenuItems = []*fyne.MenuItem{
		fyne.NewMenuItem("Open Config...", g.OpenConfig),
		fyne.NewMenuItem("Save Config", g.SaveConfig),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Quit", g.Quit),
	}

	fileMenu := fyne.NewMenu("File", fileMenuItems...)
	editMenu := fyne.NewMenu("Edit", fyne.NewMenuItem("Edit Config", g.EditConfig))

	g.menu = fyne.NewMainMenu(fileMenu, editMenu)
	g.RootWin.SetMainMenu(g.menu)
}

func (g *GUI) Quit() {
	fmt.Println("Quitting")
	close(g.QuitCh)
}
