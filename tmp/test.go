package main

import (
	"fmt"
	"os"

	"github.com/ssebs/go-mmp/config"
	"github.com/ssebs/go-mmp/gui"
	"github.com/ssebs/go-mmp/macro"
)

func main() {
	cliFlags := config.ParseFlags()
	conf, err := config.NewConfig(cliFlags)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	macroMgr, err := macro.NewMacroManager(conf)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	if cliFlags.GUIMode != config.NOTSET {
		macroMgr.Config.GUIMode = cliFlags.GUIMode
	}

	g := gui.NewGUI(macroMgr)

	g.EditConfig()
	g.App.Run()
}

// import (
// 	"time"

// 	"github.com/go-vgo/robotgo"
// )

// func main() {
// 	// robotgo.KeySleep = 100
// 	// robotgo.KeyTap("ctrl", "shift", "esc")

// 	keys := []string{"cmd", "r"}
// 	// robotgo.KeyTap(keys[0], keys[1])
// 	shortcut(keys, 150*time.Millisecond)

// }

// func shortcut(keys []string, delay time.Duration) error {
// 	// Hold down all keys
// 	for _, key := range keys {
// 		if err := robotgo.KeyDown(key); err != nil {
// 			return err
// 		}
// 	}
// 	// Delay
// 	time.Sleep(delay)

// 	// Release all keys
// 	for _, key := range keys {
// 		if err := robotgo.KeyUp(key); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

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
