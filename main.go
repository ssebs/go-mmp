package main

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/ssebs/go-mmp/controllers"
	"github.com/ssebs/go-mmp/macro"
	"github.com/ssebs/go-mmp/models"
	"github.com/ssebs/go-mmp/serialdevice"
	"github.com/ssebs/go-mmp/utils"
	"github.com/ssebs/go-mmp/views"
)

func main() {
	cliFlags := models.ParseFlags()

	conf, err := models.NewConfigFromFile(cliFlags)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		views.ShowErrorDialogAndRun(err) // TODO: only if GUIMode is not set to daemon
	}

	macroMgr, err := macro.NewMacroManager(conf)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		views.ShowErrorDialogAndRun(err) // TODO: only if GUIMode is not set to daemon
	}

	if cliFlags.GUIMode != models.NOTSET {
		macroMgr.Config.GUIMode = cliFlags.GUIMode
	}

	// TODO: refactor this section to support daemon / CLI only

	// TODO: move creating ui to func
	mmpApp := app.New()
	rootWin := mmpApp.NewWindow(utils.ProjectName)
	rootWin.Resize(fyne.NewSize(400, 400))
	rootWin.CenterOnScreen()

	mainGUI := views.NewMacroRunnerView(conf.Columns, rootWin)
	mainGUIController := controllers.NewMacroRunnerController(conf, mainGUI, macroMgr)
	rootWin.SetContent(mainGUIController.MacroRunnerView)

	// If GUI only mode, ShowAndRun instead of continuing with serial stuff.
	// This will "block" until the window is closed, then exit
	if conf.GUIMode == models.GUIOnly {
		rootWin.SetOnClosed(func() {
			fmt.Println("gui only closed")
			os.Exit(0)
		})
		rootWin.ShowAndRun()
		return
	}

	// Else, do the serial stuff

	// Connect Serial Device from the config
	arduino, err := serialdevice.NewSerialDeviceFromConfig(conf, time.Millisecond*20)
	if err != nil {
		views.ShowErrorDialogAndRunWithLink(err, conf.ConfigFullPath)
		// TODO: show list of devices to select from and update config
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
	go mainGUIController.ListenForDisplayButtonPress(displayBtnch, quitch)

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
					err := macroMgr.RunMacroById(btn)
					if err != nil {
						slog.Warn(err.Error())
					}
				}
			case <-quitch:
				break free
			}
		}
		mmpApp.Quit()
	}()

	// Finally, display the GUI once everything is loaded & loop
	rootWin.ShowAndRun()
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
