package views

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type DragDirection int

const (
	DRAG_HORIZONTAL DragDirection = iota
	DRAG_VERTICAL
	DRAG_BOTH
)

// Ensure interface implementation.
var _ fyne.Widget = (*DragAndDropView)(nil)

type DragAndDropView struct {
	widget.BaseWidget
	dragItems        *fyne.Container
	draggedItemIdx   int
	latestDraggedIdx int
	dragDirection    DragDirection
	OnItemsSwapped   func(idx1, idx2 int)
}

func NewDragAndDropView(containingBox *fyne.Container, ddir DragDirection) *DragAndDropView {
	view := &DragAndDropView{
		dragItems:        containingBox,
		draggedItemIdx:   -1,
		latestDraggedIdx: -1,
		dragDirection:    ddir,
	}

	view.ExtendBaseWidget(view)
	return view
}

/* Callback overrides */
func (v *DragAndDropView) SetOnItemsSwapped(f func(idx1, idx2 int)) {
	v.OnItemsSwapped = f
}
func (v *DragAndDropView) SetDragItems(items []fyne.CanvasObject) {
	v.dragItems.RemoveAll()
	for _, item := range items {
		v.dragItems.Add(container.NewBorder(
			nil, nil,
			widget.NewIcon(theme.MenuIcon()),
			nil,
			item,
		))
	}
	v.dragItems.Refresh()
}

/* Actions Drag and Drop funcs */

// TODO: draw dragged item on top z idx
func (v *DragAndDropView) Dragged(e *fyne.DragEvent) {
	v.latestDraggedIdx = v.getDragIconIdxAtPosition(e.Position)

	// not currently dragging anything AND clicking over a non-nil item
	if v.draggedItemIdx == -1 && v.latestDraggedIdx != -1 {
		v.draggedItemIdx = v.latestDraggedIdx
		// fmt.Printf("dragging: %d\n", v.draggedItemIdx)
	}

	// currently dragging
	if v.draggedItemIdx != -1 {
		// Allow releasing over the whole item, not just the icon
		v.latestDraggedIdx = v.getDragItemIdxAtPosition(e.Position)

		newPos := v.dragItems.Objects[v.draggedItemIdx].Position()
		switch v.dragDirection {
		case DRAG_BOTH:
			newPos = newPos.AddXY(e.Dragged.DX, e.Dragged.DY)
		case DRAG_HORIZONTAL:
			newPos = newPos.AddXY(e.Dragged.DX, 0)
		case DRAG_VERTICAL:
			newPos = newPos.AddXY(0, e.Dragged.DY)
		}

		v.dragItems.Objects[v.draggedItemIdx].Move(newPos)
	}
}
func (v *DragAndDropView) DragEnd() {
	fmt.Printf("draggedIdx: %d, hoverIdx: %d\n", v.draggedItemIdx, v.latestDraggedIdx)

	if v.latestDraggedIdx != -1 && v.draggedItemIdx != v.latestDraggedIdx {
		if v.OnItemsSwapped != nil {
			v.OnItemsSwapped(v.draggedItemIdx, v.latestDraggedIdx)
		} else {
			fmt.Printf("released over: %d\n", v.latestDraggedIdx)
			fmt.Printf("swap %d and %d\n", v.draggedItemIdx, v.latestDraggedIdx)
		}
	}
	v.draggedItemIdx, v.latestDraggedIdx = -1, -1
	v.dragItems.Refresh()
}

// func (v *ActionDragView) Tapped(e *fyne.PointEvent) {
// 	fmt.Printf("tapped: x: %.1f y: %.1f\n", e.Position.X, e.Position.Y)
// 	v.getDragIconIdxAtPosition(e.Position)
// }

// Return the idx of v.dragItems where the mousePos clicks.
// Only matches on the drag icon
// Returns -1 if there's no match
func (v *DragAndDropView) getDragIconIdxAtPosition(mousePos fyne.Position) int {
	for idx, item := range v.dragItems.Objects {

		// Get the position of the drag icon, so we can only drag from that
		// since we're using a border, we need to get the 2nd idx item
		itemIcon := item.(*fyne.Container).Objects[1]
		globalItemPos := itemIcon.Position().Add(item.(*fyne.Container).Position())

		if withinBounds(mousePos, globalItemPos, itemIcon.Size()) && v.draggedItemIdx != idx {
			return idx
		}
	}

	return -1
}

// Return the idx of v.dragItems where the mousePos clicks.
// Returns -1 if there's no match
func (v *DragAndDropView) getDragItemIdxAtPosition(mousePos fyne.Position) int {
	for idx, item := range v.dragItems.Objects {
		if withinBounds(mousePos, item.Position(), item.Size()) && v.draggedItemIdx != idx {
			return idx
		}
	}

	return -1
}

// withinBounds checks if a point is within the bounds of a rectangle defined by position and size.
func withinBounds(point, pos fyne.Position, size fyne.Size) bool {
	return point.X >= pos.X && point.X <= pos.X+size.Width && point.Y >= pos.Y && point.Y <= pos.Y+size.Height
}

func (v *DragAndDropView) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(v.dragItems)
}