package gui

import (
	"fyne.io/fyne/v2"
)

func (g *GUI) initMenu() {
	fileMenu := fyne.NewMenu("File",
		fyne.NewMenuItem("Open Config", g.OpenConfig),
		fyne.NewMenuItem("Save Config", g.SaveConfig),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Quit", g.Quit),
	)
	editMenu := fyne.NewMenu("Edit",
		fyne.NewMenuItem("Edit Config", g.EditConfig),
	)

	g.menu = fyne.NewMainMenu(fileMenu, editMenu)
	g.RootWin.SetMainMenu(g.menu)
}
