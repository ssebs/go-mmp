package main

import (
	"errors"
	"fmt"
	"log"
	"time"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/ssebs/go-mmp/gui"
	"github.com/ssebs/go-mmp/macro"
	"github.com/ssebs/go-mmp/serialdevice"
	"github.com/ssebs/go-mmp/utils"
)

func runActionIDFromSerial(actionID string) (shouldBreak bool) {
	iActionID, err := utils.StringToInt(actionID)
	if errors.Is(err, utils.ErrCannotParseIntFromEmptyString{}) {
		log.Println(err)
	} else if err != nil {
		log.Println(err)
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
	// gui.ShowDialog("title", "some contents")
	// gui.ShowErrorDialog(errors.New("test"))

	macroMgr, err := macro.NewMacroManager("") // replace "" with path to config
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("Config: %s", macroMgr.Config)
	g := gui.NewGUI(macroMgr.Config, 1)
	arduino, err := serialdevice.NewSerialDeviceFromConfig(macroMgr.Config, time.Millisecond*20)

	// Show error dialog
	if err != nil {
		gui.ShowErrorDialog(err)
	}
	defer arduino.CloseConnection()

	// GUI container
	container := container.NewVBox()
	// Display button pressed
	pressedLabel := widget.NewLabel("Button Pressed: ")

	// Run listener
	go listener(pressedLabel, &arduino)

	// Create button to test CTRL + SHIFT + ESC hotkey
	tmBtn := widget.NewButton("Open Task Manager", macro.OpenTaskManager)

	container.Add(pressedLabel)
	container.Add(tmBtn)

	g.SetContent(container)
	g.ShowAndRun()
}
