package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/ssebs/go-mmp/macro"
	"github.com/ssebs/go-mmp/serialdevice"
	"github.com/ssebs/go-mmp/utils"
	"log"
	"time"
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

		macro.OpenTaskManager()
		fmt.Printf("pressed: %d\n", iActionID)
	default:
		fmt.Printf("pressed: %d\n", iActionID)
	}
	return false
}

func listener(btnLabel *widget.Label, sd *serialdevice.SerialDevice) {
	shouldQuit := false
	for !shouldQuit {
		// log.Println("shouldquit: ", shouldQuit)
		actionID, err := sd.Listen()
		if err != nil {
			log.Println("Listen err: ", err)
		}
		log.Println("actionID: ", actionID)
		btnLabel.SetText(fmt.Sprintf("Button Pressed: %s", actionID))
		shouldQuit = runActionIDFromSerial(actionID)
	}
	log.Println("Exiting listener")
}

func main() {
	app := app.New()
	win := app.NewWindow(projName)
	win.Resize(fyne.NewSize(300, 200))
	win.CenterOnScreen()
	macroMgr, err := macro.NewMacroManager("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Config: %s", macroMgr.Config)
	arduino, err := serialdevice.NewSerialDeviceFromConfig(*macroMgr.Config, time.Millisecond*20)

	// Show error dialog
	if err != nil {
		errDialog := dialog.NewError(err, win)
		errDialog.Show()
		errDialog.SetOnClosed(func() {
			log.Fatal(err)
		})
		win.ShowAndRun()
	}
	defer arduino.CloseConnection()

	// // GUI container
	// container := container.NewVBox()
	// // Display button pressed
	// pressedLabel := widget.NewLabel("Button Pressed: ")

	// // Run listener
	// go listener(pressedLabel, &arduino)

	// // Create button to test CTRL + SHIFT + ESC hotkey
	// tmBtn := widget.NewButton("Open Task Manager", macro.OpenTaskManager)

	// container.Add(pressedLabel)
	// container.Add(tmBtn)

	// win.SetContent(container)
	// win.ShowAndRun()
}
