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

/* How drag and drop should work:
- Remake grid? or track next/last in LL

*/

type DragBoxWidget struct {
	widget.BaseWidget
	Rect    *canvas.Rectangle
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
		Title:   widget.NewLabelWithStyle(title, fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		Rect:    canvas.NewRectangle(bgcolor),
		BGColor: bgcolor,
		FGColor: fgcolor,
		EditBtn: widget.NewButton("Edit", editCallback),
		Config:  conf,
		Cols:    conf.MacroLayout.SizeX,
		Grid:    make([]*fyne.Container, len(conf.Macros)),
	}

	for pos := 1; pos <= len(dbw.Config.Macros); pos++ {
		macro := dbw.Config.Macros[config.BtnId(pos)]
		dbw.Grid[pos-1] = container.NewStack(canvas.NewRectangle(color.Gray{0x20}), widget.NewLabel(macro.Name))
		dbw.Grid[pos-1].Objects[0].Resize(fyne.NewSquareSize(80))
		dbw.Grid[pos-1].Objects[1].Resize(fyne.NewSquareSize(64))
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
	c := container.NewStack(dbw.Rect, g)
	return widget.NewSimpleRenderer(c)
}

func (dbw *DragBoxWidget) Tapped(e *fyne.PointEvent) {
	fmt.Println("tapped, e:", e.Position)
}

func (dbw *DragBoxWidget) Dragged(e *fyne.DragEvent) {
	fmt.Println("dragged, epos:", e.Position)
	fmt.Println("dragged, edrag:", e.Dragged)

	// poz := make([]fyne.Position, 0, len(dbw.Grid.Objects))
	// // if e pos is over a grid item, then move positions
	// for _, item := range dbw.Grid.Objects {
	// 	poz = append(poz, item.Position())
	// }
	// fmt.Println("poz")
	// fmt.Println(poz)

	dbw.Move(dbw.Position().AddXY(e.Dragged.DX, e.Dragged.DY))
}
func (dbw *DragBoxWidget) DragEnd() {
	fmt.Println("drag end")
}
