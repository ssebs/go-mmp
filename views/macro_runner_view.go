package views

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/ssebs/go-mmp/models"
)

// Ensure interface implementation.
var _ fyne.Widget = (*MacroRunnerView)(nil)

type MacroRunnerView struct {
	widget.BaseWidget
	cols            int
	macrosContainer *fyne.Container
	OnMacroTapped   func(m *models.Macro)
	OnOpenConfig    func()
	OnQuit          func()
	OnEditConfig    func()

	mainMenu *fyne.MainMenu
	rootWin  fyne.Window
}

func NewMacroRunnerView(cols int, rootWin fyne.Window) *MacroRunnerView {
	view := &MacroRunnerView{
		cols:            cols,
		macrosContainer: container.NewGridWithColumns(cols),
		mainMenu:        nil,
		rootWin:         rootWin,
	}

	view.mainMenu = fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("Open Config", func() { view.OnOpenConfig() }),
			fyne.NewMenuItem("Reset Config", func() { fmt.Println("reset config") }),
			fyne.NewMenuItemSeparator(),
			fyne.NewMenuItem("Quit", func() { view.OnQuit() }),
		),
		fyne.NewMenu("Edit",
			fyne.NewMenuItem("Edit Config", func() { view.OnEditConfig() }),
		),
	)

	rootWin.SetMainMenu(view.mainMenu)

	view.ExtendBaseWidget(view)
	return view
}

func (v *MacroRunnerView) ShowPressedAnimation(idx int, delay time.Duration) {
	go func(idx int, delay time.Duration, v *MacroRunnerView) {
		btn := v.macrosContainer.Objects[idx].(*widget.Button)

		btn.Importance = widget.HighImportance
		btn.Refresh()

		time.Sleep(delay)

		btn.Importance = widget.MediumImportance
		btn.Refresh()

	}(idx, delay, v)
}

// SEE: MacroEditorView.SetActions(...)
func (v *MacroRunnerView) SetMacros(macros []*models.Macro) {
	v.macrosContainer.RemoveAll()

	for _, macro := range macros {
		macroBtn := widget.NewButton(macro.Name, func() {
			v.OnMacroTapped(macro)
		})
		v.macrosContainer.Add(macroBtn)
	}
	v.macrosContainer.Refresh()
}

func (v *MacroRunnerView) SetOnMacroTapped(f func(*models.Macro)) {
	v.OnMacroTapped = f
}
func (v *MacroRunnerView) SetOnOpenConfig(f func()) {
	v.OnOpenConfig = f
}
func (v *MacroRunnerView) SetOnQuit(f func()) {
	v.OnQuit = f
}
func (v *MacroRunnerView) SetOnEditConfig(f func()) {
	v.OnEditConfig = f
}
func (v *MacroRunnerView) SetCols(c int) {
	v.cols = c
	v.macrosContainer.Layout = (container.NewGridWithColumns(c)).Layout
	v.macrosContainer.Refresh()
}

func (v *MacroRunnerView) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(v.macrosContainer)
}
