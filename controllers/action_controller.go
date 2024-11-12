package controllers

// Ensure interface implementation.
var _ Notifier = (*ActionController)(nil)

type ActionController struct {
	ControllerNotifier
}

func (ac *ActionController) Notify(payload interface{}) {}

func (ac *ActionController) Subscribe() {}
