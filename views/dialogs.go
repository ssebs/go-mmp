package views

import (
	"fmt"
	"log"
	"net/url"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// /* Dialogs */

// ShowErrorDialogAndRunWithLink will create a new error window displaying the text of the error.
// Takes in an error, and an optional link. If the link is added, a hyperlink will be created
// at the bottom so the user can click on it.
func ShowErrorDialogAndRunWithLink(err error, link string) {
	curApp := fyne.CurrentApp()
	if curApp == nil {
		curApp = app.New()
	}

	w := curApp.NewWindow("Error!")

	// What to do if the button / close btn are pressed
	errFunc := func() {
		log.Fatal("error", err.Error())
	}
	// Container for the dialog stuff
	container := container.NewVBox()

	// Add widgets to it
	lbl := widget.NewLabel(fmt.Sprintf("Error: %s", err.Error()))
	btn := widget.NewButton("OK", errFunc)
	container.Add(lbl)

	// Add link if it's not empty
	if link != "" {
		link = filepath.ToSlash(link)
		filePrefix := "file://"
		if !strings.HasPrefix(filePrefix, link) {
			link = filePrefix + link
		}
		if url, err := url.Parse(link); err == nil {
			fmt.Println("url:", url)
			container.Add(widget.NewHyperlink(link[len(filePrefix):], url))
		} else {
			fmt.Println(err)
		}
	}

	// Add button at the bottom
	container.Add(btn)

	w.SetContent(container)
	w.SetOnClosed(errFunc)
	w.CenterOnScreen()
	// TODO: Make this work after gui is initialized
	w.ShowAndRun()
}

// ShowErrorDialogAndRun will create a new error window displaying the text of the error.
// Takes in an error, and an optional link. If the link is added, a hyperlink will be created
// at the bottom so the user can click on it.
func ShowErrorDialogAndRun(err error) {
	ShowErrorDialogAndRunWithLink(err, "")
}
