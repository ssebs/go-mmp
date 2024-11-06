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
	nameEntryBinding binding.String
	nameEntry        *widget.Entry
	actionsScroll    *container.Scroll
}

func NewActionEdtior(app fyne.App, conf *config.Config, macro config.Macro) *ActionEditor {
	ae := &ActionEditor{
		app:              app,
		Config:           conf,
		Macro:            macro,
		nameEntryBinding: nil,
		nameEntry:        nil,
		actionsScroll:    nil,
	}

	// Action Name
	ae.nameEntryBinding = binding.NewString()
	ae.nameEntryBinding.Set(ae.Macro.Name)

	ae.nameEntry = widget.NewEntryWithData(ae.nameEntryBinding)
	ae.nameEntry.Validator = nil
	ae.nameEntry.OnChanged = ae.NameChanged

	// Scrollbox for actions
	ae.actionsScroll = container.NewVScroll(container.NewVBox())
	ae.actionsScroll.Resize(ae.actionsScroll.Size().AddWidthHeight(0, 400))

	// Add actions to actionsScroll's VBox
	for _, action := range ae.Macro.Actions {
		ae.actionsScroll.Content.(*fyne.Container).Add(
			NewActionItemEdtior(conf, action),
		)
	}

	ae.ExtendBaseWidget(ae)
	return ae
}

func (ae *ActionEditor) NameChanged(s string) {
	ae.nameEntryBinding.Set(s)
}

func (ae *ActionEditor) ShowNewWindow() {
	newWin := ae.app.NewWindow("Edit Actions")
	newWin.SetContent(ae)
	newWin.Resize(fyne.NewSquareSize(400))
	newWin.CenterOnScreen()
	newWin.Show()
}

func (ae *ActionEditor) CreateRenderer() fyne.WidgetRenderer {
	c := container.NewBorder(
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
	return widget.NewSimpleRenderer(c)
}
