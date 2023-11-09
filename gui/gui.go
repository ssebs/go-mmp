package gui

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"github.com/ssebs/go-mmp/config"
	"github.com/ssebs/go-mmp/utils"
)

type GUISize fyne.Size

var dialogSize = fyne.Size{Width: 200, Height: 150}

func (g *GUISize) GetSizeAsString() string {
	return fmt.Sprintf("%f.2x%f.2", g.Width, g.Height)
}

// GUI
type GUI struct {
	Size    GUISize
	Monitor int
	App     fyne.App
	RootWin fyne.Window
}

// Create a new GUI, given a GUISize and a monitor #
func NewGUI(config *config.Config, monitor int) *GUI {
	gs := GUISize{float32(config.MacroLayout.Width), float32(config.MacroLayout.Height)}
	gui := &GUI{Size: gs, Monitor: monitor}

	gui.App = app.New()
	gui.RootWin = gui.App.NewWindow(utils.ProjectName)
	gui.RootWin.Resize(fyne.NewSize(gs.Width, gs.Height))
	gui.RootWin.CenterOnScreen()

	return gui
}

// Add content to the GUI.RootWin
func (g *GUI) SetContent(c fyne.CanvasObject) {
	g.RootWin.SetContent(c)
}

// Run GUI.RootWin.ShowAndRun()
func (g *GUI) ShowAndRun() {
	g.RootWin.ShowAndRun()
}

// Create a fyne app & attach a window
// Returns the window
// TODO: make this not quit the app if closed
func newAppWindow(title string, size fyne.Size) fyne.Window {
	app := app.New()
	win := app.NewWindow(title)
	win.CenterOnScreen()
	win.Resize(size)
	return win
}

// Static dialogs
// TODO: fix these...
func ShowDialog(title, msg string) {
	win := newAppWindow(title, dialogSize)
	d := dialog.NewInformation(title, msg, win)
	d.Resize(dialogSize)
	d.SetOnClosed(func() {
		win.Close()
	})
	d.Show()
	win.ShowAndRun()
}

func ShowErrorDialog(err error) {
	win := newAppWindow(err.Error(), dialogSize)
	d := dialog.NewError(err, win)
	d.Resize(dialogSize)
	d.SetOnClosed(func() {
		win.Close()
		log.Fatal(err)
	})
	d.Show()
	win.ShowAndRun()
}
