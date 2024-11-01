package widgets

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/ssebs/go-mmp/config"
)

// Ensure interface implementation.
var _ fyne.Widget = (*MacroEditBox)(nil)

type MacroEditBox struct {
	widget.BaseWidget
	Config *config.Config
	Macro  config.Macro
}

func NewEditBox(conf *config.Config, macro config.Macro) *MacroEditBox {
	eb := &MacroEditBox{
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
		}),
		nil, nil,
		widget.NewLabelWithStyle(eb.Macro.Name, fyne.TextAlignCenter, fyne.TextStyle{}),
	)

	outer := canvas.NewRectangle(color.RGBA{30, 30, 30, 255})

	// outer.Resize(c.Size().AddWidthHeight(20, 20)) // how to make this work??

	return widget.NewSimpleRenderer(container.NewStack(outer, c))
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