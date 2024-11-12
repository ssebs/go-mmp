package controllers

type Notifier interface {
	Notify(payload interface{})
	Subscribe()
}

type ControllerNotifier struct {
	Followers []Notifier
}

func (cn *ControllerNotifier) NotifyAll(payload interface{}) {
	for _, follower := range cn.Followers {
		follower.Notify(payload)
	}
}
func (cn *ControllerNotifier) Subscribe(n Notifier) {
	cn.Followers = append(cn.Followers, n)
}
