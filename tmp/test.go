package main

import (
	"fmt"
	"log"

	"github.com/ssebs/go-mmp/serialdevice"
)

// test serial
func main() {
	arduino, err := serialdevice.NewSerialDevice("COM7", 9600)
	if err != nil {
		log.Fatal(err)
	}
	defer arduino.CloseConnection()

	buf := make([]byte, 1)

	for {
		// fmt.Println("starting scan")
		n, err := arduino.Conn.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		if n != 0 {
			str := fmt.Sprintf("- %s\n", buf[0:n])
			fmt.Println(str)
		}
	}

	// for {
	// 	fmt.Println("starting scan")
	// 	scanner := bufio.NewScanner(arduino.Conn)
	// 	for scanner.Scan() {
	// 		txt := scanner.Text()
	// 		fmt.Println(":", txt)
	// 	}
	// 	fmt.Println("out of scan")
	// }

}

// var TealColor = color.RGBA{0, 120, 120, 255}
// var GrayColor = color.RGBA{60, 60, 60, 255}

// func main() {
// 	testApp := app.New()
// 	win := testApp.NewWindow("TEST")

// 	testBox := views.NewColorBorderBox(12, TealColor, container.NewGridWithColumns(2,
// 		widget.NewButton("test", nil),
// 		widget.NewLabel("test"),
// 		widget.NewLabel("test2"),
// 	))

// 	testBox.Item.(*fyne.Container).Objects[0].(*widget.Button).OnTapped = func() {
// 		testBox.BGColor = color.RGBA{255, 0, 0, 255}
// 		testBox.PadWidth = 32
// 		testBox.Refresh()
// 	}

// 	win.SetContent(container.NewCenter(testBox))

// 	win.CenterOnScreen()
// 	win.Resize(fyne.NewSize(300, 300))
// 	win.ShowAndRun()
// }

// test medatata editor

// func main() {
// 	testApp := app.New()
// 	win := testApp.NewWindow("TEST")

// 	mm := models.NewDefaultMetadata()
// 	mv := views.NewMetadataEditorView()
// 	mc := controllers.NewMetadataController(mm, mv)

// 	mc.UpdateMetadataView()

// 	win.SetContent(container.NewBorder(
// 		widget.NewSeparator(),
// 		widget.NewSeparator(),
// 		widget.NewSeparator(),
// 		widget.NewSeparator(),
// 		mv,
// 	))

// 	win.CenterOnScreen()
// 	win.Resize(fyne.NewSize(300, 500))
// 	win.ShowAndRun()

// }
// serialdevice
// test MacroEditorView

// func main() {
// 	testApp := app.New()
// 	win := testApp.NewWindow("TEST")

// 	mm := models.NewMacro("TestMacro", []*models.Action{
// 		models.NewAction("PressRelease", "ENTER"),
// 		models.NewAction("Delay", "200ms"),
// 		models.NewAction("SendText", "GG"),
// 		models.NewAction("PressRelease", "ENTER"),
// 	})
// 	mv := views.NewMacroEditorView(win)
// 	mc := controllers.NewMacroController(mm, mv)
// 	mc.UpdateMacroView()

// 	win.SetContent(container.NewBorder(
// 		widget.NewSeparator(),
// 		widget.NewSeparator(),
// 		widget.NewSeparator(),
// 		widget.NewSeparator(),
// 		mc.MacroEditorView,
// 	))

// 	win.CenterOnScreen()
// 	win.Resize(fyne.NewSize(300, 500))
// 	win.ShowAndRun()
// }

// func main() {
// 	cliFlags := config.ParseFlags()
// 	conf, err := config.NewConfig(cliFlags)
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, err.Error())
// 	}

// 	macroMgr, err := macro.NewMacroManager(conf)
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, err.Error())
// 	}

// 	if cliFlags.GUIMode != models.NOTSET {
// 		macroMgr.Config.GUIMode = cliFlags.GUIMode
// 	}

// 	g := gui.NewGUI(macroMgr)

// 	g.EditConfig()
// 	g.App.Run()
// }
