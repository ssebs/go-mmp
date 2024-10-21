package gui

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/ssebs/go-mmp/config"
)

type DragBoxWidget struct {
	widget.BaseWidget
	BGRect  *canvas.Rectangle
	BGColor color.Color
	FGColor color.Color
	Title   *widget.Label
	EditBtn *widget.Button
	Config  *config.Config
	Cols    int
	Grid    []*fyne.Container
}

func NewDragBoxWidget(title string, conf *config.Config, bgcolor, fgcolor color.Color, editCallback func()) *DragBoxWidget {
	dbw := &DragBoxWidget{
		BGColor: bgcolor,
		FGColor: fgcolor,
		BGRect:  canvas.NewRectangle(bgcolor),
		Title:   widget.NewLabelWithStyle(title, fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		EditBtn: widget.NewButton("Edit", editCallback),
		Config:  conf,
		Cols:    conf.MacroLayout.SizeX,
		Grid:    make([]*fyne.Container, len(conf.Macros)),
	}

	// Fill the grid with widgets gen'd from Macros
	for pos := 0; pos < len(dbw.Config.Macros); pos++ {
		macro := dbw.Config.Macros[config.BtnId(pos+1)]
		dbw.Grid[pos] = container.NewStack(canvas.NewRectangle(color.Gray{0x20}), widget.NewLabel(macro.Name))
		dbw.Grid[pos].Resize(fyne.NewSquareSize(500))
		dbw.Grid[pos].Objects[1].Resize(fyne.NewSquareSize(64))
	}

	dbw.Title.Truncation = fyne.TextTruncateEllipsis
	dbw.ExtendBaseWidget(dbw)
	return dbw
}

func (dbw *DragBoxWidget) CreateRenderer() fyne.WidgetRenderer {
	// vb := container.NewVBox(dbw.Title, dbw.EditBtn)
	g := container.NewGridWithColumns(dbw.Cols)
	for _, item := range dbw.Grid {
		g.Add(item)
	}
	c := container.NewStack(dbw.BGRect, g)
	c.Resize(g.Size().AddWidthHeight(120, 120))
	return widget.NewSimpleRenderer(c)
}

func (dbw *DragBoxWidget) Tapped(e *fyne.PointEvent) {
	fmt.Println("tapped, e:", e.Position)
}

func (dbw *DragBoxWidget) Dragged(e *fyne.DragEvent) {
	fmt.Println("dragged, epos:", e.Position)
	// fmt.Println("dragged, edrag:", e.Dragged)

	// if e pos is over a grid item (loop thru items and check item.pos + item.size..)
	// check if e.pos is

	dbw.Move(dbw.Position().AddXY(e.Dragged.DX, e.Dragged.DY))
}
func (dbw *DragBoxWidget) DragEnd() {
	fmt.Println("drag end")
}
