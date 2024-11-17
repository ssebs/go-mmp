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
	fmt.Println(e)
}
func (v *ActionDragView) Dragged(e *fyne.DragEvent) {
	fmt.Println(e)
}
func (v *ActionDragView) DragEnd() {
	fmt.Println("Done dragging")
}

func (v *ActionDragView) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(container.NewBorder(
		widget.NewLabel("Test ActionDragView:"),
		nil, nil, nil,
		v.dragItems,
	))
}
