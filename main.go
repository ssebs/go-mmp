package main

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/micmonay/keybd_event"
	"github.com/ssebs/go-mmp/keyboard"
	"github.com/ssebs/go-mmp/serialdevice"
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

func serialListen(s *serialdevice.SerialDevice) {
	// https://stackoverflow.com/a/50091168
	reader := bufio.NewScanner(s.Conn)

	for reader.Scan() {
		numString := reader.Text()
		fmt.Println("d: ", numString)
		num, err := strconv.ParseInt(numString, 10, 32)
		if err != nil {
			log.Printf("cannot convert %v to int", numString)
		}
		runActionIDFromSerial(int(num))
	}
}

func runActionIDFromSerial(actionID int) {
	switch actionID {
	case 10:
		openTaskManager()
	}
}

func main() {

	arduino, err := serialdevice.NewSerialDevice("COM7", 9600, time.Millisecond*20)
	if err != nil {
		log.Fatal(err)
	}
	go serialListen(&arduino)
	defer arduino.CloseConnection()

	app := app.New()
	win := app.NewWindow(projName)
	win.Resize(fyne.NewSize(300, 200))
	win.CenterOnScreen()

	container := container.NewVBox()

	// Create button to test CTRL + SHIFT + ESC hotkey
	tmBtn := widget.NewButton("Open Task Manager", openTaskManager)

	container.Add(tmBtn)

	win.SetContent(container)
	win.ShowAndRun()
}
