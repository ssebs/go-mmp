package widgets

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/ssebs/go-mmp/config"
)

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
	delBtn := widget.NewButtonWithIcon("",
		theme.NewErrorThemedResource(theme.WindowCloseIcon()), func() {
			fmt.Printf("Delete %s\n", eb.Macro.Name)
			eb.Config.DelMacro(eb.Macro)
			eb.Refresh()
		})

	editBtn := widget.NewButton("Edit", func() {
		fmt.Printf("Edit %s, id:%d\n", eb.Macro.Name, eb.Config.GetIdxFromMacro(eb.Macro))
		NewActionEdtior(eb.app, eb.Config, eb.Macro).ShowNewWindow()
	})

	// Label + Del/Edit btn
	c := container.NewBorder(
		nil,
		container.NewBorder(
			nil, nil,
			delBtn,
			nil,
			editBtn,
		),
		nil, nil,
		widget.NewLabelWithStyle(eb.Macro.Name, fyne.TextAlignCenter, fyne.TextStyle{}),
	)

	outer := canvas.NewRectangle(color.RGBA{30, 30, 30, 255})

	return widget.NewSimpleRenderer(container.NewStack(outer, c))
}
