package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/ssebs/go-mmp/controllers"
	"github.com/ssebs/go-mmp/models"
)

type ActionView struct {
	label      *widget.Label
	controller *controllers.ActionController
}

func NewActionView(controller *controllers.ActionController) *ActionView {
	view := &ActionView{
		label:      widget.NewLabel("DEFAULT"),
		controller: controller,
	}

	return view
}

// updateLabel is called when the model updates, updating the label text.
func (v *ActionView) updateLabel(action models.Action) {
	v.label.SetText(action.String())
	v.label.Refresh()
}

// Widget returns the fyne.Widget representing this view.
func (v *ActionView) Widget() fyne.CanvasObject {
	return v.label
}
