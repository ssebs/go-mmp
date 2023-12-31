package gui

import (
	"fmt"
	"log"
	"net/url"
	"path/filepath"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/ssebs/go-mmp/config"
	"github.com/ssebs/go-mmp/macro"
	"github.com/ssebs/go-mmp/utils"
)

// GUI
type GUI struct {
	Size         fyne.Size
	App          fyne.App
	RootWin      fyne.Window
	macroManager *macro.MacroManager
	config       *config.Config
	grid         *fyne.Container
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
	g.grid = container.New(layout.NewGridLayout(g.config.MacroLayout.SizeX))

	for pos := 1; pos <= len(g.config.Macros); pos++ {
		macro := g.config.Macros[pos]
		// fmt.Println(pos, ":", macro)

		// Copy pos to p so it doesn't get set to the len of Macros
		p := pos
		// Create btn with lambda to run function
		btn := widget.NewButton(macro.Name, func() {
			// Runs the macro from the btn id that was clicked
			g.macroManager.RunActionFromID(fmt.Sprint(p))
		})
		// Add to the grid
		g.grid.Add(btn)
	}

	// Add grid to g.RootWin
	g.RootWin.SetContent(g.grid)
}

// ListenForDisplayButtonPress will listen for a button press then visibly update
// the button so it looks like it was pressed
func (g *GUI) ListenForDisplayButtonPress(displayBtnch chan string, quitch chan struct{}) {
free:
	for {
		select {
		case btnStr := <-displayBtnch:
			// fmt.Printf("~~DISPLAY BTN %s~~\n", btnStr)
			// If you can convert the btnstr into an int
			if iBtn, err := utils.StringToInt(btnStr); err == nil {
				// Since the buttons start at 1 in the Config, get the btn - 1
				btn := g.grid.Objects[iBtn-1].(*widget.Button)
				// Display the button press
				ShowPressedAnimation(g.macroManager.Config.Delay, btn)
			}
		case <-quitch:
			break free
		}
	}
}

// ShowPressedAnimation will change the color of the button for the delay given
func ShowPressedAnimation(delay time.Duration, btn *widget.Button) {
	btn.Importance = widget.HighImportance
	btn.Refresh()
	time.Sleep(delay)
	btn.Importance = widget.MediumImportance
	btn.Refresh()
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

/* Dialogs */

// ShowErrorDialogAndRunWithLink will create a new error window displaying the text of the error.
// Takes in an error, and an optional link. If the link is added, a hyperlink will be created
// at the bottom so the user can click on it.
func ShowErrorDialogAndRunWithLink(err error, link string) {
	w := fyne.CurrentApp().NewWindow("Error!")

	// What to do if the button / close btn are pressed
	errFunc := func() {
		log.Fatal("error", err.Error())
	}
	// Container for the dialog stuff
	container := container.NewVBox()

	// Add widgets to it
	lbl := widget.NewLabel(fmt.Sprintf("Error: %s", err.Error()))
	btn := widget.NewButton("OK", errFunc)
	container.Add(lbl)

	// Add link if it's not empty
	if link != "" {
		link = filepath.ToSlash(link)
		filePrefix := "file://"
		if !strings.HasPrefix(filePrefix, link) {
			link = filePrefix + link
		}
		if url, err := url.Parse(link); err == nil {
			fmt.Println("url:", url)
			container.Add(widget.NewHyperlink(link[len(filePrefix):], url))
		} else {
			fmt.Println(err)
		}
	}

	// Add button at the bottom
	container.Add(btn)

	w.SetContent(container)
	w.SetOnClosed(errFunc)
	w.CenterOnScreen()
	w.ShowAndRun()
}

// ShowErrorDialogAndRun will create a new error window displaying the text of the error.
// Takes in an error, and an optional link. If the link is added, a hyperlink will be created
// at the bottom so the user can click on it.
func ShowErrorDialogAndRun(err error) {
	ShowErrorDialogAndRunWithLink(err, "")
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
