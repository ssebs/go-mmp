package controllers

import (
	"github.com/beevik/guid"
)

// A Notifier is something that can send and recieve a payload, as well as
// be sub/unsubscribed to a ControllerNotifier
// GetID must be a unique *guid.Guid. TODO: GUID
type Notifier interface {
	NotifyById(id *guid.Guid, payload interface{})
	NotifyAll(payload interface{})
	ReceiveNotification(payload interface{})
	Subscribe()
	Unsubscribe()
	GetID() *guid.Guid
}

// ControllerNotifier keeps track of Followers that can be notified.
// Followers is a map of Notifier ID to Notifiers
type ControllerNotifier struct {
	Followers map[*guid.Guid]Notifier
}

func NewControllerNotifier() *ControllerNotifier {
	return &ControllerNotifier{
		Followers: make(map[*guid.Guid]Notifier),
	}
}

func (cn *ControllerNotifier) NotifyAll(payload interface{}) {
	for _, follower := range cn.Followers {
		follower.ReceiveNotification(payload)
	}
}

func (cn *ControllerNotifier) NotifyById(id *guid.Guid, payload interface{}) {
	if follower, ok := cn.Followers[id]; ok {
		follower.ReceiveNotification(payload)
	}
}

func (cn *ControllerNotifier) Subscribe(n Notifier) {
	cn.Followers[n.GetID()] = n
}

func (cn *ControllerNotifier) Unsubscribe(n Notifier) {
	delete(cn.Followers, n.GetID())
}
