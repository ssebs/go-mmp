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

// DragBoxWidget represents the draggable UI component.
type DragBoxWidget struct {
	widget.BaseWidget
	BGRect         *canvas.Rectangle
	Title          *widget.Label
	EditBtn        *widget.Button
	Config         *config.Config
	Cols           int
	Grid           []*fyne.Container
	draggedItemIdx int
	latestItemIdx  int
	g              *fyne.Container
	BGColor        color.Color
	FGColor        color.Color
}

// Ensure interface implementation.
var _ fyne.Widget = (*DragBoxWidget)(nil)
var _ fyne.WidgetRenderer = (*DragBoxRenderer)(nil)

// NewDragBoxWidget creates a new DragBoxWidget with the specified configuration.
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

	dbw.initializeGrid()
	dbw.initializeContainer()
	dbw.ExtendBaseWidget(dbw)
	return dbw
}

// CreateRenderer sets up the widget's custom renderer.
func (dbw *DragBoxWidget) CreateRenderer() fyne.WidgetRenderer {
	return &DragBoxRenderer{
		dbw:     dbw,
		objects: []fyne.CanvasObject{dbw.BGRect, dbw.g},
	}
}

// initializeGrid initializes the Grid field with containers for each macro.
func (dbw *DragBoxWidget) initializeGrid() {
	for pos := 0; pos < len(dbw.Config.Macros); pos++ {
		macro := dbw.Config.Macros[config.BtnId(pos+1)]
		dbw.Grid[pos] = container.NewStack(canvas.NewRectangle(color.Gray{0x20}), widget.NewLabel(macro.Name))
	}
}

// initializeContainer arranges grid items into a column-based container.
func (dbw *DragBoxWidget) initializeContainer() {
	dbw.g = container.NewGridWithColumns(dbw.Cols)
	for _, item := range dbw.Grid {
		dbw.g.Add(item)
	}
}

// updateGrid swaps and refreshes grid items.
func (dbw *DragBoxWidget) updateGrid() {
	dbw.initializeContainer()
	dbw.g.Refresh()

	currentSize := dbw.Size()
	dbw.Resize(currentSize.Add(fyne.NewSize(1, 1))) // Force a slight resize to refresh layout.
	dbw.Resize(currentSize)
}

// swapMacros swaps macros in the configuration and updates the UI.
func (dbw *DragBoxWidget) swapMacros(first, second int) {
	tmp := dbw.Config.Macros[config.BtnId(first+1)]
	dbw.Config.Macros[config.BtnId(first+1)] = dbw.Config.Macros[config.BtnId(second+1)]
	dbw.Config.Macros[config.BtnId(second+1)] = tmp

	dbw.Grid[first], dbw.Grid[second] = dbw.Grid[second], dbw.Grid[first]
	dbw.updateGrid()
}

// getItemInPosition returns the index of the item at a given mouse position or -1 if none.
func (dbw *DragBoxWidget) getItemInPosition(mousePos fyne.Position) int {
	for i, item := range dbw.Grid {
		if withinBounds(mousePos, item.Position(), item.Size()) && dbw.draggedItemIdx != i {
			return i
		}
	}
	return -1
}

// withinBounds checks if a point is within the bounds of a rectangle defined by position and size.
func withinBounds(point, pos fyne.Position, size fyne.Size) bool {
	return point.X >= pos.X && point.X <= pos.X+size.Width && point.Y >= pos.Y && point.Y <= pos.Y+size.Height
}

// getMacroFromIdx retrieves the macro at the specified index.
func (dbw *DragBoxWidget) getMacroFromIdx(idx int) config.Macro {
	return dbw.Config.Macros[config.BtnId(idx+1)]
}

// Tapped is triggered when the widget is tapped.
func (dbw *DragBoxWidget) Tapped(e *fyne.PointEvent) {
	if hitItem := dbw.getItemInPosition(e.Position); hitItem != -1 {
		fmt.Printf("Tapped item: %s\n", dbw.getMacroFromIdx(hitItem).Name)
	}
}

// Dragged handles dragging events within the widget.
func (dbw *DragBoxWidget) Dragged(e *fyne.DragEvent) {
	dbw.latestItemIdx = dbw.getItemInPosition(e.Position)

	if dbw.draggedItemIdx == -1 && dbw.latestItemIdx != -1 {
		dbw.draggedItemIdx = dbw.latestItemIdx
		fmt.Printf("Dragging item: %s\n", dbw.getMacroFromIdx(dbw.draggedItemIdx).Name)
	}

	if dbw.draggedItemIdx != -1 {
		dbw.Grid[dbw.draggedItemIdx].Move(dbw.Grid[dbw.draggedItemIdx].Position().AddXY(e.Dragged.DX, e.Dragged.DY))
	}
}

// DragEnd finalizes a drag operation, swapping items if necessary.
func (dbw *DragBoxWidget) DragEnd() {
	if dbw.latestItemIdx != -1 && dbw.draggedItemIdx != dbw.latestItemIdx {
		fmt.Printf("Released over %s\n", dbw.getMacroFromIdx(dbw.latestItemIdx).Name)
		dbw.swapMacros(dbw.draggedItemIdx, dbw.latestItemIdx)
	}
	dbw.draggedItemIdx, dbw.latestItemIdx = -1, -1
	dbw.updateGrid()
}

// DragBoxRenderer manages rendering for DragBoxWidget.
type DragBoxRenderer struct {
	dbw     *DragBoxWidget
	objects []fyne.CanvasObject
}

func (r *DragBoxRenderer) Layout(size fyne.Size) {
	r.dbw.BGRect.Resize(size)
	r.dbw.g.Resize(size)
}

func (r *DragBoxRenderer) MinSize() fyne.Size {
	return r.dbw.g.MinSize()
}

func (r *DragBoxRenderer) Refresh() {
	r.dbw.BGRect.FillColor = r.dbw.BGColor
	r.dbw.BGRect.Refresh()
	for _, item := range r.dbw.Grid {
		item.Refresh()
	}
	r.dbw.g.Refresh()
}

func (r *DragBoxRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *DragBoxRenderer) Destroy() {}
