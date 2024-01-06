package main

import (
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	// robotgo.KeySleep = 100
	// robotgo.KeyTap("ctrl", "shift", "esc")

	keys := []string{"cmd", "r"}
	// robotgo.KeyTap(keys[0], keys[1])
	shortcut(keys, 150*time.Millisecond)

}

func shortcut(keys []string, delay time.Duration) error {
	// Hold down all keys
	for _, key := range keys {
		if err := robotgo.KeyDown(key); err != nil {
			return err
		}
	}
	// Delay
	time.Sleep(delay)

	// Release all keys
	for _, key := range keys {
		if err := robotgo.KeyUp(key); err != nil {
			return err
		}
	}
	return nil
}

// import (
// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/container"
// 	"fyne.io/fyne/v2/widget"
// )

// func main() {
// 	myApp := app.New()
// 	myWindow1 := myApp.NewWindow("Window 1")

// 	myWindow1.SetContent(container.NewVBox(
// 		widget.NewLabel("Hello from Window 1"),
// 		widget.NewButton("Open Window 2", func() {
// 			go func() {
// 				myWindow2 := myApp.NewWindow("Window 2")
// 				myWindow2.SetContent(container.NewVBox(
// 					widget.NewLabel("Hello from Window 2"),
// 				))
// 				myWindow2.Show()
// 			}()
// 		}),
// 	))

// 	myWindow1.ShowAndRun()
// }
