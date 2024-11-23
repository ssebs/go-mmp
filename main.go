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

/* To Fix before v2 release:
- Change GUIMode In ConfigEditor => Reconnect to serial device
- Cleanup printing config when saving
- Close when Saving?
- Don't open 3 windows for ConfigEditor
- Support 1 indexing

*/

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

	// listener channels
	btnch := make(chan string, 2)
	quitch := make(chan struct{})
	displayBtnch := make(chan string, 1)

	// Run Serial Listener
	go arduino.Listen(btnch, quitch)

	// Visible button press listener
	go mainGUIController.ListenForDisplayButtonPress(displayBtnch, quitch)

	// Do something when btnch gets data
	go RunMacroOnDataIn(btnch, displayBtnch, quitch, mmpApp, macroMgr)

	// Finally, display the GUI once everything is loaded & loop
	rootWin.ShowAndRun()
}

func RunMacroOnDataIn(btnch chan string, displayBtnch chan string, quitch chan struct{}, mmpApp fyne.App, macroMgr *macro.MacroManager) {
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
}
