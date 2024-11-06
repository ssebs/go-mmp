package widgets

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Toast is a simple toast notification widget.
type Toast struct {
	widget.BaseWidget
	title   string
	content string
	app     fyne.App
	Delay   time.Duration
}

// NewToast creates a new Toast with a title and content, and an app reference.
func NewToast(app fyne.App, title, content string, delay time.Duration) *Toast {
	t := &Toast{title: title, content: content, app: app, Delay: delay}
	t.ExtendBaseWidget(t)
	return t
}

// Show displays the toast for a limited time (3 seconds in this example).
func (t *Toast) Show() {
	window := t.app.NewWindow("")
	window.SetContent(container.NewVBox(
		widget.NewLabel(t.title),
		widget.NewLabel(t.content),
	))
	window.Resize(fyne.NewSize(200, 100))
	window.SetFixedSize(true)
	window.CenterOnScreen()
	window.Show()

	go func() {
		time.Sleep(t.Delay) // Duration the toast is visible
		window.Close()
	}()
}

// CreateRenderer defines the appearance of the Toast widget.
func (t *Toast) CreateRenderer() fyne.WidgetRenderer {
	title := widget.NewLabel(t.title)
	content := widget.NewLabel(t.content)
	box := container.NewVBox(title, content)

	return widget.NewSimpleRenderer(box)
}
