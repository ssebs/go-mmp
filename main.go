package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

const projName = "Go-MMP"

func main() {
	app := app.New()
	win := app.NewWindow(projName)

	win.SetContent(widget.NewLabel(projName))
	win.ShowAndRun()
}
