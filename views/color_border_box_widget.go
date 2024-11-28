package views

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var _ fyne.Widget = (*ColorBorderBox)(nil)

type ColorBorderBox struct {
	widget.BaseWidget
	Item         fyne.CanvasObject
	PadWidth     float32
	BGColor      color.Color
	FGColor      color.Color
	padContainer *fyne.Container
	bgRect       *canvas.Rectangle
}

func NewColorBorderBox(padWidth float32, bgColor color.Color, item fyne.CanvasObject) *ColorBorderBox {
	box := &ColorBorderBox{
		Item:     item,
		PadWidth: float32(padWidth),
		BGColor:  bgColor,
		FGColor:  color.RGBA{0, 0, 0, 0},
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
		canvas.NewRectangle(box.FGColor),
		box.Item,
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
