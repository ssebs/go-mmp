package widgets

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/ssebs/go-mmp/config"
)

// Ensure interface implementation.
var _ fyne.Widget = (*ActionEditor)(nil)

type ActionEditor struct {
	widget.BaseWidget
	Config           *config.Config
	Macro            config.Macro
	app              fyne.App
	content          *fyne.Container
	nameEntryBinding binding.String
	nameEntry        *widget.Entry
	actionsScroll    *container.Scroll
}

func NewActionEdtior(app fyne.App, conf *config.Config, macro config.Macro) *ActionEditor {
	ae := &ActionEditor{
		app:              app,
		Config:           conf,
		Macro:            macro,
		content:          nil,
		nameEntryBinding: nil,
		nameEntry:        nil,
		actionsScroll:    nil,
	}

	// Action Name
	ae.nameEntryBinding = binding.NewString()
	ae.nameEntryBinding.Set(ae.Macro.Name)

	ae.nameEntry = widget.NewEntryWithData(ae.nameEntryBinding)
	ae.nameEntry.Validator = nil
	ae.nameEntry.OnChanged = func(s string) {
		ae.nameEntryBinding.Set(s)
	}

	// Scrollbox for actions
	ae.actionsScroll = container.NewVScroll(container.NewVBox())
	ae.actionsScroll.Resize(ae.actionsScroll.Size().AddWidthHeight(0, 400))

	// Add actions to actionsScroll's VBox
	for _, action := range ae.Macro.Actions {
		ae.actionsScroll.Content.(*fyne.Container).Add(
			NewActionItemEdtior(conf, action),
		)
	}

	ae.content = container.NewBorder(
		container.NewVBox(
			widget.NewLabelWithStyle(
				fmt.Sprintf("Edit %s", ae.nameEntry.Text),
				fyne.TextAlignCenter,
				fyne.TextStyle{Bold: true},
			),
			widget.NewForm(
				widget.NewFormItem("Name/Title:", ae.nameEntry),
				widget.NewFormItem("Actions", layout.NewSpacer()),
				widget.NewFormItem("", layout.NewSpacer()),
			),
		),
		container.NewHBox(
			widget.NewButton("Close", func() {
				fmt.Println("CLOSE WINDOW")
			}),
			widget.NewButton("+ Add Action", func() {
				fmt.Println("ADD ACTION")
			}),
			widget.NewButton("Save", func() {
				msg := fmt.Sprintf("Saved %s", ae.nameEntry.Text)
				fmt.Println(msg)
				// app.SendNotification(fyne.NewNotification("Go-MMP", msg))
				// NewToast(ae.app, "Mini Macro Pad", msg, 1*time.Second).Show()
			}),
		),
		nil, nil,
		ae.actionsScroll,
	)

	ae.ExtendBaseWidget(ae)
	return ae
}

// func (ae *ActionEditor) newActionItemEditor(action map[string]string) *fyne.Container {
// 	// Get the key/vals from the action
// 	funcName, funcParam := utils.GetKeyVal(action)

// 	paramEntryBinding := binding.NewString()
// 	paramEntryBinding.Set(funcParam)

// 	paramEntry := widget.NewEntryWithData(paramEntryBinding)
// 	paramEntry.Validator = nil

// 	funcSelect := widget.NewSelect(macro.FunctionList, func(s string) {
// 		fmt.Println(s)
// 	})
// 	funcSelect.SetSelected(funcName)

// 	container := container.NewBorder(
// 		nil,
// 		nil,
// 		container.NewHBox(
// 			widget.NewIcon(theme.MenuIcon()),
// 			funcSelect,
// 		),
// 		container.NewHBox(
// 			widget.NewIcon(theme.WindowCloseIcon()),
// 			layout.NewSpacer()),
// 		paramEntry,
// 	)
// 	return container
// }

func (ae *ActionEditor) Show() {
	newWin := ae.app.NewWindow("Edit Actions")
	newWin.SetContent(ae.content)
	newWin.Resize(fyne.NewSquareSize(400))
	newWin.CenterOnScreen()
	newWin.Show()
}

func (ae *ActionEditor) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(ae.content)
}
