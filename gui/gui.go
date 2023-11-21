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
	config  *config.Config
	App     fyne.App
	RootWin fyne.Window
	// MacroGrid map[int]config.Macro
}

// Create a new GUI, given a GUISize and a monitor #
func NewGUI(config *config.Config) *GUI {
	gs := GUISize{float32(config.MacroLayout.Width), float32(config.MacroLayout.Height)}
	gui := &GUI{Size: gs, config: config}

	gui.App = app.New()
	gui.RootWin = gui.App.NewWindow(utils.ProjectName)
	gui.RootWin.Resize(fyne.NewSize(gs.Width, gs.Height))
	gui.RootWin.CenterOnScreen()

	gui.initMacroGrid()

	return gui
}

// useful?!?
// // MacroGrid
// type MacroGrid struct {
// 	win fyne.Window
// }

func replace_me() {
	fmt.Println("Replace the use of this function!")
}

func (g *GUI) initMacroGrid() {
	// Generate grid from g.config.MacroLayout.SizeX/Y
	grid := container.New(layout.NewGridLayout(g.config.MacroLayout.SizeX))

	for pos := 1; pos <= len(g.config.Macros); pos++ {
		macro := g.config.Macros[pos]

		// fmt.Println(pos, ":", macro)
		// Add to the grid
		btn := widget.NewButton(macro.Name, replace_me)
		grid.Add(btn)
		// In each container, add the macro Name & Id info
		// add keys to Macros[position] for new stuff, don't copy that map
		// TODO: see above
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
