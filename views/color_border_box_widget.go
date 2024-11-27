package views

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var _ fyne.Widget = (*ColorBorderBox)(nil)

type ColorBorderBox struct {
	widget.BaseWidget
	ItemContainer *fyne.Container
	PadWidth      float32
	BGColor       color.Color
	padContainer  *fyne.Container
	bgRect        *canvas.Rectangle
}

func NewColorBorderBox(padWidth float32, bgColor color.Color, itemContainer *fyne.Container) *ColorBorderBox {
	box := &ColorBorderBox{
		ItemContainer: itemContainer,
		PadWidth:      float32(padWidth),
		BGColor:       bgColor,
	}
	box.genPadContainer()
	box.genBGRect()

	box.ExtendBaseWidget(box)
	return box
}

func (box *ColorBorderBox) Refresh() {
	box.genBGRect()
	box.genPadContainer()
}

func (box *ColorBorderBox) genPadContainer() {
	box.padContainer = container.New(
		layout.NewCustomPaddedLayout(box.PadWidth, box.PadWidth, box.PadWidth, box.PadWidth),
		canvas.NewRectangle(theme.Color(theme.ColorNameBackground)),
		box.ItemContainer,
	)
	box.padContainer.Refresh()
}
func (box *ColorBorderBox) genBGRect() {
	box.bgRect = canvas.NewRectangle(box.BGColor)
	box.bgRect.Refresh()
}

func (box *ColorBorderBox) CreateRenderer() fyne.WidgetRenderer {
	c := container.NewStack(
		box.bgRect,
		box.padContainer,
	)
	return widget.NewSimpleRenderer(c)
}
