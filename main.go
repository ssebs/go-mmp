package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/ssebs/go-mmp/config"
	"github.com/ssebs/go-mmp/gui"
	"github.com/ssebs/go-mmp/macro"
	"github.com/ssebs/go-mmp/serialdevice"
	"github.com/ssebs/go-mmp/utils"
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
	// CLI flags
	cliFlags := parseFlags()

	// Init MacroManager & Load config
	macroMgr, err := macro.NewMacroManager(cliFlags.DoResetConfig)
	if err != nil {
		gui.ShowErrorDialogAndRun(err)
	}

	// Init GUI from macroMgr
	g := gui.NewGUI(macroMgr)

	// If GUI only mode, ShowAndRun. This will block until the window is closed.
	if cliFlags.IsGUIOnly {
		g.RootWin.SetOnClosed(func() {
			fmt.Println("gui only closed")
			os.Exit(0)
		})
		g.ShowAndRun()
	}

	// Otherwise, connect to the serial device & init the listeners

	// Connect Serial Device from the config
	arduino, err := serialdevice.NewSerialDeviceFromConfig(macroMgr.Config, time.Millisecond*20)
	if err != nil {
		path, _ := config.GetConfigFilePath()
		gui.ShowErrorDialogAndRunWithLink(err, path)
	}
	defer arduino.CloseConnection()

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

// CLI flag values will be stored in this
type CLIFlags struct {
	IsGUIOnly     bool
	DoResetConfig bool
}

// parseFlags will parse the CLI flags that may have been used.
// Useful for disabling the serial listening functionality
func parseFlags() CLIFlags {
	// TODO: Allow for GUI to pop up with a failed arduino connection... might not be possible with fyne
	// for now, let's use CLI flags
	cliFlags := CLIFlags{}
	flag.BoolVar(&cliFlags.IsGUIOnly, "gui-only", false, fmt.Sprintf("Open %s in GUI Only Mode. Useful if you don't have a working arduino.", utils.ProjectName))
	flag.BoolVar(&cliFlags.DoResetConfig, "reset-config", false, "If you want to reset your mmpConfig.yml file.")

	flag.Parse()
	return cliFlags
}
