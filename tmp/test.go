package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var _ fyne.Widget = (*ColorBorderBox)(nil)

var TealColor = color.RGBA{0, 120, 120, 255}
var GrayColor = color.RGBA{60, 60, 60, 255}

// // Make "border" layout
// // Stacks the first item with pad
// type SebsCoolLayout struct{}

// func (d *SebsCoolLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
// 	w, h := float32(0), float32(0)
// 	for _, o := range objects {
// 		childSize := o.MinSize()

// 		w += childSize.Width
// 		h += childSize.Height
// 	}
// 	return fyne.NewSize(w, h)
// }

// func (d *SebsCoolLayout) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
// 	pos := fyne.NewPos(0, containerSize.Height-d.MinSize(objects).Height)
// 	for _, o := range objects {
// 		size := o.MinSize()
// 		o.Resize(size)
// 		o.Move(pos)

// 		pos = pos.Add(fyne.NewPos(size.Width, size.Height))
// 	}
// }

type ColorBorderBox struct {
	widget.BaseWidget
	ItemContainer *fyne.Container
	PadWidth      float32
	BGColor       color.Color
}

// Create div in a div, outer has padding and bg color
func NewTestBox(padWidth float32, bgColor color.Color, itemContainer *fyne.Container) *ColorBorderBox {
	box := &ColorBorderBox{
		ItemContainer: itemContainer,
		PadWidth:      float32(padWidth),
		BGColor:       TealColor,
	}
	box.ExtendBaseWidget(box)
	return box
}

func (box *ColorBorderBox) CreateRenderer() fyne.WidgetRenderer {
	c := container.NewStack(
		canvas.NewRectangle(box.BGColor),
		container.New(
			layout.NewCustomPaddedLayout(box.PadWidth, box.PadWidth, box.PadWidth, box.PadWidth),
			canvas.NewRectangle(theme.Color(theme.ColorNameBackground)),
			box.ItemContainer,
		),
	)
	return widget.NewSimpleRenderer(c)
}

func main() {
	testApp := app.New()
	win := testApp.NewWindow("TEST")

	testBox := NewTestBox(12, TealColor, container.NewGridWithColumns(2,
		widget.NewButton("test", nil),
		widget.NewLabel("test"),
		widget.NewLabel("test2"),
	))

	win.SetContent(container.NewCenter(testBox))

	win.CenterOnScreen()
	win.Resize(fyne.NewSize(300, 300))
	win.ShowAndRun()
}

// test medatata editor

// func main() {
// 	testApp := app.New()
// 	win := testApp.NewWindow("TEST")

// 	mm := models.NewDefaultMetadata()
// 	mv := views.NewMetadataEditorView()
// 	mc := controllers.NewMetadataController(mm, mv)

// 	mc.UpdateMetadataView()

// 	win.SetContent(container.NewBorder(
// 		widget.NewSeparator(),
// 		widget.NewSeparator(),
// 		widget.NewSeparator(),
// 		widget.NewSeparator(),
// 		mv,
// 	))

// 	win.CenterOnScreen()
// 	win.Resize(fyne.NewSize(300, 500))
// 	win.ShowAndRun()

// }

// test MacroEditorView

// func main() {
// 	testApp := app.New()
// 	win := testApp.NewWindow("TEST")

// 	mm := models.NewMacro("TestMacro", []*models.Action{
// 		models.NewAction("PressRelease", "ENTER"),
// 		models.NewAction("Delay", "200ms"),
// 		models.NewAction("SendText", "GG"),
// 		models.NewAction("PressRelease", "ENTER"),
// 	})
// 	mv := views.NewMacroEditorView()
// 	mc := controllers.NewMacroController(mm, mv)
// 	mc.UpdateMacroView()

// 	win.SetContent(container.NewBorder(
// 		widget.NewSeparator(),
// 		widget.NewSeparator(),
// 		widget.NewSeparator(),
// 		widget.NewSeparator(),
// 		mc.MacroEditorView,
// 	))

// 	win.CenterOnScreen()
// 	win.Resize(fyne.NewSize(300, 500))
// 	win.ShowAndRun()
// }

// func main() {
// 	cliFlags := config.ParseFlags()
// 	conf, err := config.NewConfig(cliFlags)
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, err.Error())
// 	}

// 	macroMgr, err := macro.NewMacroManager(conf)
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, err.Error())
// 	}

// 	if cliFlags.GUIMode != models.NOTSET {
// 		macroMgr.Config.GUIMode = cliFlags.GUIMode
// 	}

// 	g := gui.NewGUI(macroMgr)

// 	g.EditConfig()
// 	g.App.Run()
// }
