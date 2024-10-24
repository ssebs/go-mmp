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
	latestItemIdx  int
	g              *fyne.Container
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
		latestItemIdx:  -1,
	}

	dbw.genGrid()
	dbw.genG()

	dbw.Title.Truncation = fyne.TextTruncateEllipsis
	dbw.ExtendBaseWidget(dbw)
	return dbw
}

// TODO: Implement fyne.WidgetRenderer
func (dbw *DragBoxWidget) CreateRenderer() fyne.WidgetRenderer {
	dbw.genG()
	c := container.NewStack(dbw.BGRect, dbw.g)
	return widget.NewSimpleRenderer(c)
}

func (dbw *DragBoxWidget) genG() {
	dbw.g = container.NewGridWithColumns(dbw.Cols)
	for _, item := range dbw.Grid {
		dbw.g.Add(item)
	}
	dbw.g.Refresh()
}

func (dbw *DragBoxWidget) genGrid() {
	// Fill the grid with widgets gen'd from Macros
	for pos := 0; pos < len(dbw.Config.Macros); pos++ {
		macro := dbw.Config.Macros[config.BtnId(pos+1)]
		dbw.Grid[pos] = container.NewStack(canvas.NewRectangle(color.Gray{0x20}), widget.NewLabel(macro.Name))
		// dbw.Grid[pos].Objects[1].Resize(fyne.NewSquareSize(64))
	}
}

func (dbw *DragBoxWidget) Tapped(e *fyne.PointEvent) {
	fmt.Println("tapped, e:", e.Position)
	hitItem := dbw.getItemInPosition(e.Position)
	if hitItem != -1 {
		fmt.Printf("hit the %s item\n", dbw.getMacroFromIdx(hitItem).Name)
	}
}

func (dbw *DragBoxWidget) Dragged(e *fyne.DragEvent) {
	// fmt.Println("dragged, epos:", e.Position)
	// fmt.Println("dragged, edrag:", e.Dragged)

	// FIX: only works on last element in list. 2nd elem can hover over 1st, 3rd can hover over 2nd and 1st

	// Use dbw.latestItemIdx for box being hovered over (update every time)
	dbw.latestItemIdx = dbw.getItemInPosition(e.Position) // slow
	// fmt.Println("lastItemIdx:", dbw.latestItemIdx)
	// if dbw.latestItemIdx != -1 {
	// 	fmt.Println("hovering over", dbw.getMacroFromIdx(dbw.latestItemIdx).Name)
	// }

	// Use dbw.draggedItemIdx for box being dragged (update only after letting go)
	if dbw.draggedItemIdx == -1 {
		if dbw.latestItemIdx != -1 {
			dbw.draggedItemIdx = dbw.latestItemIdx
			fmt.Println("dragging the", dbw.getMacroFromIdx(dbw.draggedItemIdx).Name, "item.")
			dbw.Grid[dbw.draggedItemIdx].Move(dbw.Grid[dbw.draggedItemIdx].Position().AddXY(e.Dragged.DX, e.Dragged.DY))
		}
	} else {
		dbw.Grid[dbw.draggedItemIdx].Move(dbw.Grid[dbw.draggedItemIdx].Position().AddXY(e.Dragged.DX, e.Dragged.DY))
	}
}

func (dbw *DragBoxWidget) DragEnd() {
	// fmt.Println("drag end")

	if dbw.latestItemIdx != -1 {
		fmt.Printf("released over %s\n", dbw.getMacroFromIdx(dbw.latestItemIdx).Name)

		if dbw.draggedItemIdx != dbw.latestItemIdx {
			dbw.swapMacros(dbw.draggedItemIdx, dbw.latestItemIdx)
		}
	}

	dbw.draggedItemIdx = -1
	dbw.latestItemIdx = -1
}

func (dbw *DragBoxWidget) swapMacros(first, second int) {
	// Update Macro data internally in config
	// TODO: Move to setter function
	tmp := dbw.Config.Macros[config.BtnId(first+1)]
	dbw.Config.Macros[config.BtnId(first+1)] = dbw.Config.Macros[config.BtnId(second+1)]
	dbw.Config.Macros[config.BtnId(second+1)] = tmp
	// TODO: config.save()

	// dbw.genGrid() // <- breaks it!
	// Update UI

	// tmp2 := dbw.Grid[first]
	// dbw.Grid[first] = dbw.Grid[second]
	// dbw.Grid[first].Refresh()

	// dbw.Grid[second] = tmp2
	// dbw.Grid[second].Refresh()
	dbw.g.Refresh()

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
				if dbw.draggedItemIdx == i {
					continue
				}
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
