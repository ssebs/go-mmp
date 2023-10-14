package main

import (
	"log"
	"time"

	// "fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/app"
	// "fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/widget"
	"github.com/micmonay/keybd_event"
	"github.com/ssebs/go-mmp/keyboard"
	"github.com/ssebs/go-mmp/mmp"
)

const projName = "Go-MMP"

func openTaskManager() {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}
	keeb := keyboard.Keyboard{KeyBonding: &kb}
	hkm := keyboard.HotKeyModifiers{Shift: true, Control: true}
	keeb.RunHotKey(10*time.Millisecond, hkm, keybd_event.VK_ESC)
}

// Update var in UI

func main() {

	arduino, err := mmp.NewMMPSerialDevice("COM7", 9600, time.Millisecond*20)
	if err != nil {
		log.Fatal(err)
	}
	defer arduino.Close()

	// app := app.New()
	// win := app.NewWindow(projName)
	// win.Resize(fyne.NewSize(300, 200))
	// win.CenterOnScreen()

	// container := container.NewVBox()

	// // Create button to test CTRL + SHIFT + ESC hotkey
	// tmBtn := widget.NewButton("Open Task Manager", openTaskManager)

	// container.Add(tmBtn)

	// win.SetContent(container)
	// win.ShowAndRun()
}
