package controllers

import (
	"fmt"

	"github.com/beevik/guid"
)

type ActionController struct {
	*ControllerNotifier
	controllerID *guid.Guid
}

func NewActionController(controllerNotifier *ControllerNotifier) *ActionController {
	ac := &ActionController{
		ControllerNotifier: controllerNotifier,
		controllerID:       guid.New(),
	}

	return ac
}

// Notifier implementation
var _ Notifier = (*ActionController)(nil)

func (ac *ActionController) NotifyById(id *guid.Guid, payload interface{}) {
	ac.ControllerNotifier.NotifyById(id, payload)
}
func (ac *ActionController) NotifyAll(payload interface{}) {
	ac.ControllerNotifier.NotifyAll(payload)
}

func (ac *ActionController) ReceiveNotification(payload interface{}) {
	fmt.Println(payload)
}

func (ac *ActionController) Subscribe() {
	ac.ControllerNotifier.Subscribe(ac)
}
func (ac *ActionController) Unsubscribe() {}
func (ac *ActionController) GetID() *guid.Guid {
	return ac.controllerID
}
