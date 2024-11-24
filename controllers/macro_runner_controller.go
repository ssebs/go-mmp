package controllers

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"github.com/ssebs/go-mmp/macro"
	"github.com/ssebs/go-mmp/models"
	"github.com/ssebs/go-mmp/serialdevice"
	"github.com/ssebs/go-mmp/utils"
	"github.com/ssebs/go-mmp/views"
)

type MacroRunnerController struct {
	*models.Config
	*views.MacroRunnerView
	*views.ConfigEditorView
	*ConfigController
	*macro.MacroManager
	*serialdevice.SerialDevice
}

func NewMacroRunnerController(m *models.Config, v *views.MacroRunnerView, mm *macro.MacroManager) *MacroRunnerController {
	cc := &MacroRunnerController{
		Config:           m,
		MacroRunnerView:  v,
		ConfigEditorView: views.NewConfigEditorView(m.Columns),
		MacroManager:     mm,
		SerialDevice:     nil,
	}
	cc.ConfigController = NewConfigController(cc.Config, cc.ConfigEditorView)

	cc.MacroRunnerView.SetOnMacroTapped(func(macro *models.Macro) {
		cc.MacroManager.RunMacro(macro)
	})

	cc.MacroRunnerView.SetOnEditConfig(func() {
		win := fyne.CurrentApp().NewWindow("Macro Editor")
		win.CenterOnScreen()
		win.Resize(fyne.NewSize(300, 500))

		win.SetContent(cc.ConfigController.ConfigEditorView)

		win.Show()
		win.SetOnClosed(func() {
			cc.UpdateConfigView()

			if err := cc.ReconnectSerialDevice(); err != nil {
				fmt.Fprint(os.Stderr, err)
			}
		})
	})

	cc.MacroRunnerView.SetOnOpenConfig(func() {
		yamlPath, err := utils.GetYAMLFilename(false)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
		}
		fmt.Println("Saving to", yamlPath)

		if err := cc.Config.OpenConfig(yamlPath); err != nil {
			fmt.Fprint(os.Stderr, err)
		}
		cc.ConfigController.UpdateConfigView()
		cc.UpdateConfigView()
	})

	cc.MacroRunnerView.SetOnQuit(func() {
		fyne.CurrentApp().Quit()
	})

	cc.UpdateConfigView()
	return cc
}

func (cc *MacroRunnerController) SetSerialDevice(s *serialdevice.SerialDevice) {
	cc.SerialDevice = s
}

// Reconnect Serial device if it's different
func (cc *MacroRunnerController) ReconnectSerialDevice() error {
	if cc.SerialDevice != nil &&
		(cc.SerialDevice.PortName != cc.Config.Metadata.SerialPortName) ||
		(cc.SerialDevice.Mode.BaudRate != cc.metaController.SerialBaudRate) {

		err := cc.SerialDevice.ChangePortAndReconnect(
			cc.Config.Metadata.SerialPortName,
			cc.Config.Metadata.SerialBaudRate,
		)
		return err
	}
	return fmt.Errorf("SerialDevice not set")
}

func (cc *MacroRunnerController) UpdateConfigView() {
	cc.MacroRunnerView.SetCols(cc.Config.Columns)
	cc.MacroRunnerView.SetMacros(cc.Config.Macros)
}

// ListenForDisplayButtonPress will listen for a button press then visibly update
// the button so it looks like it was pressed
func (cc *MacroRunnerController) ListenForDisplayButtonPress(displayBtnch chan string, quitch chan struct{}) {
free:
	for {
		select {
		case btnStr := <-displayBtnch:
			if iBtn, err := utils.StringToInt(btnStr); err == nil {
				cc.ShowPressedAnimation(iBtn, cc.Delay)
			}
		case <-quitch:
			break free
		}
	}
}
