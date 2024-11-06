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
var _ fyne.Widget = (*DragBox)(nil)

// DragBox represents the draggable UI component.
type DragBox struct {
	widget.BaseWidget
	Config         *config.Config
	fgColor        color.Color
	bgColor        color.Color
	backgroundRect *canvas.Rectangle
	grid           *fyne.Container
	cols           int
	draggedItemIdx int
	latestItemIdx  int
	app            fyne.App
}

// TODO: support changing the item container type (from editbox)
func NewDragBox(app fyne.App, conf *config.Config, bgcolor, fgcolor color.Color) *DragBox {
	dbw := &DragBox{
		app:            app,
		Config:         conf,
		cols:           conf.MacroLayout.SizeX,
		grid:           nil,
		draggedItemIdx: -1,
		latestItemIdx:  -1,
		bgColor:        bgcolor,
		fgColor:        fgcolor,
		backgroundRect: canvas.NewRectangle(bgcolor),
	}

	dbw.initGrid()
	dbw.ExtendBaseWidget(dbw)
	return dbw
}

func (dbw *DragBox) initGrid() {
	dbw.grid = container.NewGridWithColumns(dbw.cols)

	for pos := 0; pos < len(dbw.Config.Macros); pos++ {
		macroPos := config.BtnId(pos + 1)
		dbw.grid.Add(
			NewEditBox(dbw.app, dbw.Config, dbw.Config.Macros[macroPos]),
		)
	}
}

// swapMacros swaps macros in the config by grid idx and updates the UI.
func (dbw *DragBox) swapMacros(first, second int) {
	fmt.Printf("Swapping %q and %q\n", dbw.getMacroFromIdx(first).Name, dbw.getMacroFromIdx(second).Name)
	// TODO: Don't change config.macros directly!
	dbw.Config.Macros[config.BtnId(first+1)], dbw.Config.Macros[config.BtnId(second+1)] =
		dbw.Config.Macros[config.BtnId(second+1)], dbw.Config.Macros[config.BtnId(first+1)]

	dbw.grid.Objects[first], dbw.grid.Objects[second] =
		dbw.grid.Objects[second], dbw.grid.Objects[first]

	dbw.grid.Refresh()
}

// Tapped is triggered when the widget is tapped.
func (dbw *DragBox) Tapped(e *fyne.PointEvent) {
	if hitItem := dbw.getItemIdxAtPosition(e.Position); hitItem != -1 {
		fmt.Printf("Tapped %q\n", dbw.getMacroFromIdx(hitItem).Name)
	}
}

// Dragged handles dragging events within the widget.
func (dbw *DragBox) Dragged(e *fyne.DragEvent) {
	dbw.latestItemIdx = dbw.getItemIdxAtPosition(e.Position)

	// not currently dragging anything AND clicking over a non-nil item
	if dbw.draggedItemIdx == -1 && dbw.latestItemIdx != -1 {
		dbw.draggedItemIdx = dbw.latestItemIdx
		fmt.Printf("Dragging %q", dbw.getMacroFromIdx(dbw.draggedItemIdx).Name)
	}

	// currently dragging
	if dbw.draggedItemIdx != -1 {
		dbw.grid.Objects[dbw.draggedItemIdx].Move(dbw.grid.Objects[dbw.draggedItemIdx].Position().AddXY(e.Dragged.DX, e.Dragged.DY))
	}
}

// DragEnd finalizes a drag operation, swapping items if necessary.
func (dbw *DragBox) DragEnd() {
	if dbw.latestItemIdx != -1 && dbw.draggedItemIdx != dbw.latestItemIdx {
		fmt.Printf(" ...released over %q\n", dbw.getMacroFromIdx(dbw.latestItemIdx).Name)
		dbw.swapMacros(dbw.draggedItemIdx, dbw.latestItemIdx)
	}
	dbw.draggedItemIdx, dbw.latestItemIdx = -1, -1
	dbw.grid.Refresh()
}

// CreateRenderer sets up the widget's custom renderer.
func (dbw *DragBox) CreateRenderer() fyne.WidgetRenderer {
	c := container.NewStack(dbw.backgroundRect, dbw.grid)
	return widget.NewSimpleRenderer(c)
}

// withinBounds checks if a point is within the bounds of a rectangle defined by position and size.
func withinBounds(point, pos fyne.Position, size fyne.Size) bool {
	return point.X >= pos.X && point.X <= pos.X+size.Width && point.Y >= pos.Y && point.Y <= pos.Y+size.Height
}

// getMacroFromIdx retrieves the macro at the specified index.
func (dbw *DragBox) getMacroFromIdx(idx int) config.Macro {
	return dbw.Config.Macros[config.BtnId(idx+1)]
}

// getItemIdxAtPosition returns the index of the item at a given mouse position or -1 if none.
func (dbw *DragBox) getItemIdxAtPosition(mousePos fyne.Position) int {
	for i, item := range dbw.grid.Objects {
		if withinBounds(mousePos, item.Position(), item.Size()) && dbw.draggedItemIdx != i {
			return i
		}
	}
	return -1
}
