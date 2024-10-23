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
	BGRect         *canvas.Rectangle
	BGColor        color.Color
	FGColor        color.Color
	Title          *widget.Label
	EditBtn        *widget.Button
	Config         *config.Config
	Cols           int
	Grid           []*fyne.Container
	draggedItemIdx int
}

func NewDragBoxWidget(title string, conf *config.Config, bgcolor, fgcolor color.Color, editCallback func()) *DragBoxWidget {
	dbw := &DragBoxWidget{
		BGColor:        bgcolor,
		FGColor:        fgcolor,
		BGRect:         canvas.NewRectangle(bgcolor),
		Title:          widget.NewLabelWithStyle(title, fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		EditBtn:        widget.NewButton("Edit", editCallback),
		Config:         conf,
		Cols:           conf.MacroLayout.SizeX,
		Grid:           make([]*fyne.Container, len(conf.Macros)),
		draggedItemIdx: -1,
	}

	// Fill the grid with widgets gen'd from Macros
	for pos := 0; pos < len(dbw.Config.Macros); pos++ {
		macro := dbw.Config.Macros[config.BtnId(pos+1)]
		dbw.Grid[pos] = container.NewStack(canvas.NewRectangle(color.Gray{0x20}), widget.NewLabel(macro.Name))
		// dbw.Grid[pos].Objects[1].Resize(fyne.NewSquareSize(64))
	}

	dbw.Title.Truncation = fyne.TextTruncateEllipsis
	dbw.ExtendBaseWidget(dbw)
	return dbw
}

func (dbw *DragBoxWidget) CreateRenderer() fyne.WidgetRenderer {
	// vb := container.NewVBox(dbw.Title, dbw.EditBtn)
	// TODO: move to struct
	g := container.NewGridWithColumns(dbw.Cols)
	for _, item := range dbw.Grid {
		g.Add(item)
	}
	c := container.NewStack(dbw.BGRect, g)
	return widget.NewSimpleRenderer(c)
}

func (dbw *DragBoxWidget) Tapped(e *fyne.PointEvent) {
	fmt.Println("tapped, e:", e.Position)
	dbw.draggedItemIdx = dbw.getItemInPosition(e.Position)
	if dbw.draggedItemIdx != -1 {
		fmt.Printf("hit the %s item\n", dbw.getMacroFromIdx(dbw.draggedItemIdx).Name)
	}
}

func (dbw *DragBoxWidget) Dragged(e *fyne.DragEvent) {
	fmt.Println("dragged, epos:", e.Position)
	// fmt.Println("dragged, edrag:", e.Dragged)

	mousePosX := e.Position.X
	mousePosY := e.Position.Y

	if dbw.draggedItemIdx != -1 {

		// See if we're over another box

		// Move dragged item
		dbw.Grid[dbw.draggedItemIdx].Move(dbw.Grid[dbw.draggedItemIdx].Position().AddXY(e.Dragged.DX, e.Dragged.DY))
		return
	}

	// find which item we're clicking
	for i, item := range dbw.Grid {
		itemStartPosX := item.Position().X
		itemStartPosY := item.Position().Y
		itemEndPosX := itemStartPosX + item.Size().Width
		itemEndPosY := itemStartPosY + item.Size().Height

		if mousePosX >= itemStartPosX && mousePosX <= itemEndPosX {
			if mousePosY >= itemStartPosY && mousePosY <= itemEndPosY {
				fmt.Println("Hovering over ", dbw.Config.Macros[config.BtnId(i+1)].Name)
				dbw.draggedItemIdx = i
				dbw.Grid[i].Move(dbw.Grid[i].Position().AddXY(e.Dragged.DX, e.Dragged.DY))
				return
			}
		}

	}

	// if e pos is over a grid item (loop thru items and check item.pos + item.size..)

	// swap: move items in dbw.grid, and reset grid objects

	// dbw.Move(dbw.Position().AddXY(e.Dragged.DX, e.Dragged.DY))
}
func (dbw *DragBoxWidget) DragEnd() {
	fmt.Println("drag end")
	dbw.draggedItemIdx = -1
}

// return -1 if no match
func (dbw *DragBoxWidget) getItemInPosition(mousePos fyne.Position) int {
	// find which item we're clicking
	for i, item := range dbw.Grid {
		itemStartPosX := item.Position().X
		itemStartPosY := item.Position().Y
		itemEndPosX := itemStartPosX + item.Size().Width
		itemEndPosY := itemStartPosY + item.Size().Height

		if mousePos.X >= itemStartPosX && mousePos.X <= itemEndPosX {
			if mousePos.Y >= itemStartPosY && mousePos.Y <= itemEndPosY {
				return i
			}
		}
	}
	return -1
}

// idx is from 0, but the macro is from 1. Use the 0 as idx
func (dbw *DragBoxWidget) getMacroFromIdx(idx int) config.Macro {
	return dbw.Config.Macros[config.BtnId(idx+1)]
}
