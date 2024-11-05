package widgets

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/ssebs/go-mmp/config"
)

/*
Main gui
- VBox:
	- Label: Edit (macroname)
	- Hbox:
		- Label (name)
		- Entry for (macroname)
	- Label: (Actions)
	- List:
		- <See below>
	- HBox:
		- Undo btn?
		- +NewAction btn
		- Save btn
- List for actions: (each item)
	- Hbox:
		- icon for drag/moving position
		- Select for action
		- entry for action param
		- x icon for delete
*/

// Ensure interface implementation.
var _ fyne.Widget = (*MacroEditBox)(nil)

type MacroEditBox struct {
	widget.BaseWidget
	Config *config.Config
	Macro  config.Macro
	app    fyne.App
}

func NewEditBox(app fyne.App, conf *config.Config, macro config.Macro) *MacroEditBox {
	eb := &MacroEditBox{
		app:    app,
		Config: conf,
		Macro:  macro,
	}
	eb.ExtendBaseWidget(eb)
	return eb
}

func (eb *MacroEditBox) CreateRenderer() fyne.WidgetRenderer {
	c := container.NewBorder(
		nil,
		widget.NewButton("Edit", func() {
			fmt.Printf("Edit %s, id:%d\n", eb.Macro.Name, eb.getIdxFromMacro(eb.Macro.Name))
			eb.runActionEditorWindow()
		}),
		nil, nil,
		widget.NewLabelWithStyle(eb.Macro.Name, fyne.TextAlignCenter, fyne.TextStyle{}),
	)

	outer := canvas.NewRectangle(color.RGBA{30, 30, 30, 255})

	// outer.Resize(c.Size().AddWidthHeight(20, 20)) // how to make this work??

	return widget.NewSimpleRenderer(container.NewStack(outer, c))
}

func (eb *MacroEditBox) runActionEditorWindow() error {
	newWin := eb.app.NewWindow("Edit Actions")
	nameEntryBinding := binding.NewString()
	nameEntryBinding.Set(eb.Macro.Name)

	nameEntry := widget.NewEntryWithData(nameEntryBinding)
	nameEntry.Validator = nil

	actionsScroll := container.NewVScroll(container.NewVBox())
	actionsScroll.Resize(actionsScroll.Size().AddWidthHeight(0, 400))

	// Within a Macro, there's a list of Actions to run.
	// we want to run each one in order here
	for _, action := range eb.Macro.Actions {
		// Get the key/vals from the action
		for funcName, funcParam := range action {
			actionsScroll.Content.(*fyne.Container).Add(
				widget.NewLabel(fmt.Sprintf("%s : %s", funcName, funcParam)),
			)
		}
	}

	newWin.SetContent(container.NewBorder(
		container.NewVBox(
			widget.NewLabelWithStyle(
				fmt.Sprintf("Edit %s", nameEntry.Text),
				fyne.TextAlignCenter,
				fyne.TextStyle{Bold: true},
			),
			widget.NewForm(
				widget.NewFormItem("Name", nameEntry),
				widget.NewFormItem("Actions", layout.NewSpacer()),
			),
			layout.NewSpacer(),
		),
		widget.NewButton("Save", func() {
			fmt.Println("SAVED", nameEntry.Text)
		}),
		nil, nil,
		actionsScroll,
	))

	newWin.Resize(fyne.NewSquareSize(400))
	newWin.CenterOnScreen()
	newWin.Show()
	return nil
}

func (eb *MacroEditBox) newActionItemEditor() *fyne.Container {

	// dragIcon := nil
	// delIcon := nil

	container := container.NewHBox(
	// widget.NewIcon()MenuExpandIcon
	)

	return container
}

// get macro position from macro name, if not found return -1
func (eb *MacroEditBox) getIdxFromMacro(macroName string) config.BtnId {
	for idx, macro := range eb.Config.Macros {
		if macroName == macro.Name {
			return idx
		}
	}
	return -1
}
