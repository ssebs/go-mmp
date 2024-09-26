package main

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/ssebs/go-mmp/config"
	"github.com/ssebs/go-mmp/gui"
	"github.com/ssebs/go-mmp/macro"
	"github.com/ssebs/go-mmp/serialdevice"
)

func main() {
	cliFlags := config.ParseFlags()

	// TODO: create config, then pass ptr around instead of creating in NewMacroManager
	// Init MacroManager & Load config
	macroMgr, err := macro.NewMacroManager(cliFlags)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		gui.ShowErrorDialogAndRun(err)
	}

	if cliFlags.GUIMode != config.NOTSET {
		macroMgr.Config.GUIMode = cliFlags.GUIMode
	}

	// TODO: refactor this section to support daemon / CLI only
	g := gui.NewGUI(macroMgr)

	// If GUI only mode, ShowAndRun instead of continuing with serial stuff.
	// This will "block" until the window is closed, then exit
	if macroMgr.Config.GUIMode == config.GUIOnly {
		g.RootWin.SetOnClosed(func() {
			fmt.Println("gui only closed")
			os.Exit(0)
		})
		g.ShowAndRun()
		return
	}

	// Else, do the serial stuff

	// Connect Serial Device from the config
	arduino, err := serialdevice.NewSerialDeviceFromConfig(macroMgr.Config, time.Millisecond*20)
	if err != nil {
		path, _ := config.GetConfigFilePath()
		gui.ShowErrorDialogAndRunWithLink(err, path)
	}
	defer arduino.CloseConnection()

	// listeners
	btnch := make(chan string, 2)
	quitch := make(chan struct{})
	displayBtnch := make(chan string, 1)

	// Run Serial Listener
	// TODO: rename this
	go Listen(btnch, quitch, arduino)

	// Visible button press listener
	go g.ListenForDisplayButtonPress(displayBtnch, quitch)

	// Do something when btnch gets data
	// TODO: move to func
	go func() {
	free:
		for {
			select {
			case btn := <-btnch:
				// Only run the function if it's not blank, tho
				if btn != "" {
					// send btn id to show the btn press
					displayBtnch <- btn

					// Run the action from the btn id
					err := macroMgr.RunActionFromID(btn)
					if err != nil {
						slog.Warn(err.Error())
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
				slog.Debug(fmt.Sprint("Listen err: ", err))
			}
			btnch <- actionID
		}
	}
}
