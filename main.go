package main

import (
	"fmt"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/ssebs/go-mmp/mmp"
	"github.com/ssebs/go-mmp/serialdevice"
	"github.com/ssebs/go-mmp/utils"
)

const projName = "Go-MMP"

func runActionIDFromSerial(actionID string) (shouldBreak bool) {
	iActionID, err := utils.StringToInt(actionID)
	if err != nil {
		log.Println(err.Error())
		return true
	}
	switch iActionID {
	case 9:
		return true
	case 10:
		mmp.OpenTaskManager()
		fmt.Printf("pressed: %d\n", iActionID)
	default:
		fmt.Printf("pressed: %d\n", iActionID)
	}
	return false
}

func main() {
	arduino, err := serialdevice.NewSerialDevice("COM7", 9600, time.Millisecond*20)
	if err != nil {
		log.Fatal(err)
	}
	defer arduino.CloseConnection()
	go func() {
		shouldQuit := false
		for !shouldQuit {
			shouldQuit = arduino.ListenCallback(runActionIDFromSerial)
			fmt.Println("shouldquit: ", shouldQuit)
		}
		log.Println("No longer listening for serial data, leaving goroutine")
	}()

	app := app.New()
	win := app.NewWindow(projName)
	win.Resize(fyne.NewSize(300, 200))
	win.CenterOnScreen()

	container := container.NewVBox()

	// Create button to test CTRL + SHIFT + ESC hotkey
	tmBtn := widget.NewButton("Open Task Manager", mmp.OpenTaskManager)

	container.Add(tmBtn)

	win.SetContent(container)
	win.ShowAndRun()
}
