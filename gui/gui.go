package gui

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/ssebs/go-mmp/config"
	"github.com/ssebs/go-mmp/macro"
	"github.com/ssebs/go-mmp/utils"
)

var dialogSize = fyne.Size{Width: 200, Height: 150}

// GUI
type GUI struct {
	Size         fyne.Size
	config       *config.Config
	macroManager *macro.MacroManager
	App          fyne.App
	RootWin      fyne.Window
	// MacroGrid map[int]config.Macro
}

// Create a new GUI, given a MacroManager ptr
func NewGUI(mm *macro.MacroManager) *GUI {
	gs := fyne.NewSize(float32(mm.Config.MacroLayout.Width), float32(mm.Config.MacroLayout.Height))
	gui := &GUI{Size: gs, config: mm.Config, macroManager: mm}

	gui.App = app.New()
	gui.RootWin = gui.App.NewWindow(utils.ProjectName)
	gui.RootWin.Resize(gs)
	gui.RootWin.CenterOnScreen()

	gui.initMacroGrid()

	return gui
}

// initMacroGrid will generate grid from g.config.MacroLayout.SizeX & number of Macros
func (g *GUI) initMacroGrid() {
	grid := container.New(layout.NewGridLayout(g.config.MacroLayout.SizeX))

	for pos := 1; pos <= len(g.config.Macros); pos++ {
		macro := g.config.Macros[pos]
		// fmt.Println(pos, ":", macro)

		// Copy pos to p so it doesn't get set to the len of Macros
		p := pos
		// Create btn with lambda to run function
		btn := widget.NewButton(macro.Name, func() {
			// Runs the macro from the btn id that was clicked
			g.macroManager.RunActionFromID(fmt.Sprint(p))
			// fmt.Printf("running action from %d\n", p)
		})
		// Add to the grid
		grid.Add(btn)

	}
	// Add chans for each button to show pressed state when a (physical) button is pressed

	// Add grid to g.RootWin
	g.RootWin.SetContent(grid)
}

/* Usefull stuff from the demo app:
- Containers
	- Grid
	- Split (colors look good)
- Collections
	- GridWrap
- Data Binding
- Widgets
	- Text (RichText Heading)
	- Button
- Windows
*/

/*"static" stuff below*/

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

/* Helpers */
// Add content to the GUI.RootWin
func (g *GUI) SetContent(c fyne.CanvasObject) {
	g.RootWin.SetContent(c)
}

// Run GUI.RootWin.ShowAndRun()
func (g *GUI) ShowAndRun() {
	g.RootWin.ShowAndRun()
}
