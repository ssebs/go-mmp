package main

import (
	"fmt"
	"log/slog"
	"time"

	"fyne.io/fyne/v2/widget"
	"github.com/ssebs/go-mmp/config"
	"github.com/ssebs/go-mmp/gui"
	"github.com/ssebs/go-mmp/macro"
	"github.com/ssebs/go-mmp/serialdevice"
)

// Listen for data from a *SerialDevice, to be used in a goroutine
// Takes in a btnch to send data to when the serial connection gets something,
// and a quitch if we need to stop the goroutine
func Listen(btnch chan string, quitch chan struct{}, sd *serialdevice.SerialDevice) {
free:
	// Keep looping since sd.Listen() will return if no data is sent
	for {
		select {
		case <-quitch:
			break free
		default:
			// If we get data, send to chan
			actionID, err := sd.Listen()
			if err != nil {
				slog.Debug("Listen err: ", err)
			}
			btnch <- actionID
		}
	}
}

func main() {
	// Init MacroManager & Load config
	macroMgr, err := macro.NewMacroManager()
	if err != nil {
		gui.ShowErrorDialogAndRun(err)
	}
	// fmt.Printf("Config: %s", macroMgr.Config)

	// Init GUI from macroMgr
	g := gui.NewGUI(macroMgr)

	// TODO: Allow for GUI to pop up with a failed arduino connection

	// Connect Serial Device from the config
	arduino, err := serialdevice.NewSerialDeviceFromConfig(macroMgr.Config, time.Millisecond*20)
	if err != nil {
		path, _ := config.GetConfigFilePath()
		gui.ShowErrorDialogAndRunWithLink(err, path)
	}
	defer arduino.CloseConnection()

	// Display button pressed
	pressedLabel := widget.NewLabel("Button Pressed: ")

	// Chans for listeners
	btnch := make(chan string, 2)
	quitch := make(chan struct{})
	displayBtnch := make(chan string, 1)

	// Run Serial Listener
	go Listen(btnch, quitch, arduino)

	// Visible button press listener
	go g.ListenForDisplayButtonPress(displayBtnch, quitch)

	// Do something when btnch gets data
	go func() {
	free:
		for {
			select {
			case btn := <-btnch:
				// Set the button pressed to blank if we get blank data
				pressedLabel.SetText(fmt.Sprintf("Button Pressed: %s", btn))
				// Only run the function if it's not blank, tho
				if btn != "" {
					// send btn id to show the btn press
					displayBtnch <- btn
					// Run the action from the btn id
					err := macroMgr.RunActionFromID(btn)
					if err != nil {
						slog.Warn(err.Error())
						// close(quitch)
					}
				}
			case <-quitch:
				break free
			}
		}
		g.App.Quit()
	}()

	// Finally, display the GUI once everything is loaded & loop
	g.ShowAndRun()
}
