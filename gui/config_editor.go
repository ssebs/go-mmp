package gui

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	osdialog "github.com/sqweek/dialog"
	"github.com/ssebs/go-mmp/widgets"
)

// Open a new Window and use it to edit the config
func (g *GUI) EditConfig() {
	editorWindow := g.App.NewWindow("Config Editor") // TOOD: Move this so it can be used elsewhere
	g.initEditorGUI(editorWindow)

	// Editor features:
	// [ ] Click btn to edit Macro
	// [ ] Create/Delete Macro btns
	// [x] Drag and Drop button positions on grid

	editorWindow.CenterOnScreen()
	editorWindow.Show()
}

func (g *GUI) initEditorGUI(win fyne.Window) {
	delayEntryBinding := binding.NewString()
	delayEntryBinding.Set(g.config.Delay.String())

	serialPortEntryBinding := binding.NewString()
	serialPortEntryBinding.Set(g.config.SerialDevice.PortName)

	serialBaudEntryBinding := binding.NewString()
	serialBaudEntryBinding.Set(fmt.Sprintf("%d", g.config.SerialDevice.BaudRate))

	dragBox := widgets.NewDragBox(g.App, g.config, color.RGBA{20, 20, 20, 255}, color.White)

	vbox := container.NewVBox()

	vbox.Add(widget.NewLabelWithStyle("Edit Config", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}))
	vbox.Add(widget.NewForm(
		widget.NewFormItem("Delay", widget.NewEntryWithData(delayEntryBinding)),
		widget.NewFormItem("Serial Port", widget.NewEntryWithData(serialPortEntryBinding)),
		widget.NewFormItem("Serial Baud", widget.NewEntryWithData(serialBaudEntryBinding)),
	))
	vbox.Add(widget.NewLabelWithStyle("Edit Macros", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}))
	vbox.Add(dragBox)
	vbox.Add(layout.NewSpacer())

	saveBtn := widget.NewButton("Save", func() {
		g.config.SaveConfig("")
	})
	saveBtn.Importance = widget.HighImportance

	vbox.Add(container.NewHBox(
		widget.NewButton("Open Config", g.OpenConfig),
		widget.NewButton("+ Add Macro", func() {
			// g.config.AddMacro() or something like that
			fmt.Println("ADD MACRO")
		}),
		saveBtn,
		widget.NewButton("Save As", func() {
			filename, err := getYAMLFilename(true)
			if err != nil {
				fmt.Println(err)
			}
			g.config.SaveConfig(filename)
		}),
	))
	win.SetContent(vbox)
	vbox.Refresh()
}

func (g *GUI) OpenConfig() {
	fmt.Println("OPEN CONFIG")
	filename, err := getYAMLFilename(false)
	if err != nil {
		fmt.Println(err)
	}
	g.config.OpenConfig(filename)
	g.SetContent(widget.NewLabel("test"))
	// TODO: refresh!

	// g.initEditorGUI(win) // reload
	// win.Content().Refresh()
}

// returns path to .yaml|.yml file
// isSaving sets the type to save file instead of open file.
func getYAMLFilename(isSaving bool) (string, error) {
	var filename string
	var err error

	if isSaving {
		filename, err = osdialog.File().Filter("YAML config file", "yaml", "yml").Save()
	} else {
		filename, err = osdialog.File().Filter("YAML config file", "yaml", "yml").Load()
	}

	if err != nil {
		err = fmt.Errorf("could not open YAML config file, err: %s", err)
	}
	return filename, err
}
