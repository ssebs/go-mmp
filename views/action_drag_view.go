package views

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Ensure interface implementation.
var _ fyne.Widget = (*ActionDragView)(nil)

type ActionDragView struct {
	widget.BaseWidget
	dragItems *fyne.Container
}

func NewActionDragView() *ActionDragView {
	view := &ActionDragView{
		dragItems: container.NewVBox(),
	}
	view.initTestItems()

	view.ExtendBaseWidget(view)
	return view
}

func (v *ActionDragView) initTestItems() {
	for i := 0; i < 3; i++ {
		v.dragItems.Add(container.NewHBox(
			widget.NewIcon(theme.MenuIcon()),
			widget.NewLabel(fmt.Sprint("Test Item ", i)),
		))
	}
}

/* Actions Drag and Drop funcs */
func (v *ActionDragView) Tapped(e *fyne.PointEvent) {
	fmt.Printf("tapped: x: %.1f y: %.1f\n", e.Position.X, e.Position.Y)
	v.getDragIconIdxAtPosition(e.Position)
}

// Return the idx of v.dragItems where the mousePos clicks.
// Only matches on the drag icon
// Returns -1 if there's no match
func (v *ActionDragView) getDragIconIdxAtPosition(mousePos fyne.Position) int {
	for idx, item := range v.dragItems.Objects {

		// Get the position of the drag icon, so we can only drag from that
		itemIcon := item.(*fyne.Container).Objects[0]
		globalItemPos := itemIcon.Position().Add(item.(*fyne.Container).Position())

		if withinBounds(mousePos, globalItemPos, itemIcon.Size()) {
			return idx
		}
	}

	return -1
}

// withinBounds checks if a point is within the bounds of a rectangle defined by position and size.
func withinBounds(point, pos fyne.Position, size fyne.Size) bool {
	return point.X >= pos.X && point.X <= pos.X+size.Width && point.Y >= pos.Y && point.Y <= pos.Y+size.Height
}

func (v *ActionDragView) Dragged(e *fyne.DragEvent) {
	fmt.Println(e)
}
func (v *ActionDragView) DragEnd() {
	fmt.Println("Done dragging")
}

func (v *ActionDragView) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(v.dragItems)
}
