package widgets

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/ssebs/go-mmp/config"
	"github.com/ssebs/go-mmp/macro"
	"github.com/ssebs/go-mmp/utils"
)

// Ensure interface implementation.
var _ fyne.Widget = (*ActionEditor)(nil)

type ActionEditor struct {
	widget.BaseWidget
	Config *config.Config
	Macro  config.Macro
	app    fyne.App
}

/*

TODO: IMPLEMENT THIS

*/

func NewActionEdtior(app fyne.App, conf *config.Config, macro config.Macro) *ActionEditor {
	ae := &ActionEditor{
		app:    app,
		Config: conf,
		Macro:  macro,
	}
	ae.ExtendBaseWidget(ae)
	return ae
}

func (ae *ActionEditor) Run() {
	fmt.Println("SHOW AND RUN WINDOW")
}

func (ae *ActionEditor) CreateRenderer() fyne.WidgetRenderer {
	// Label + Edit btn
	c := container.NewBorder(
		nil,
		widget.NewButton("Edit", func() {
			fmt.Println("replace_me")
		}),
		nil, nil,
		widget.NewLabelWithStyle(ae.Macro.Name, fyne.TextAlignCenter, fyne.TextStyle{}),
	)

	return widget.NewSimpleRenderer(c)
}

func (ae *ActionEditor) newActionItemEditor(action map[string]string) *fyne.Container {
	// Get the key/vals from the action
	funcName, funcParam := utils.GetKeyVal(action)

	paramEntryBinding := binding.NewString()
	paramEntryBinding.Set(funcParam)

	paramEntry := widget.NewEntryWithData(paramEntryBinding)
	paramEntry.Validator = nil

	funcSelect := widget.NewSelect(macro.FunctionList, func(s string) {
		fmt.Println(s)
	})
	funcSelect.SetSelected(funcName)

	container := container.NewBorder(
		nil,
		nil,
		container.NewHBox(
			widget.NewIcon(theme.MenuIcon()),
			funcSelect,
		),
		container.NewHBox(
			widget.NewIcon(theme.WindowCloseIcon()),
			layout.NewSpacer()),
		paramEntry,
	)
	return container
}
