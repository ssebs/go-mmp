package widgets

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/ssebs/go-mmp/config"
	"github.com/ssebs/go-mmp/macro"
	"github.com/ssebs/go-mmp/utils"
)

// Ensure interface implementation.
var _ fyne.Widget = (*ActionItemEditor)(nil)

type ActionItemEditor struct {
	widget.BaseWidget
	Config            *config.Config
	Action            map[string]string
	FuncName          string
	FuncParam         string
	content           *fyne.Container
	paramEntryBinding binding.String
	paramEntry        *widget.Entry
	funcSelect        *widget.Select
}

func NewActionItemEdtior(conf *config.Config, action map[string]string) *ActionItemEditor {
	funcName, funcParam := utils.GetKeyVal(action)
	ae := &ActionItemEditor{
		Config:            conf,
		Action:            action,
		FuncName:          funcName,
		FuncParam:         funcParam,
		content:           nil,
		paramEntryBinding: nil,
		paramEntry:        nil,
		funcSelect:        nil,
	}

	ae.paramEntryBinding = binding.NewString()
	ae.paramEntryBinding.Set(funcParam)

	ae.paramEntry = widget.NewEntryWithData(ae.paramEntryBinding)
	ae.paramEntry.Validator = nil

	ae.funcSelect = widget.NewSelect(macro.FunctionList, func(s string) {
		fmt.Println(s)
	})
	ae.funcSelect.SetSelected(funcName)

	ae.content = container.NewBorder(
		nil,
		nil,
		container.NewHBox(
			widget.NewIcon(theme.MenuIcon()),
			ae.funcSelect,
		),
		container.NewHBox(
			widget.NewIcon(theme.WindowCloseIcon()),
			layout.NewSpacer()),
		ae.paramEntry,
	)

	ae.ExtendBaseWidget(ae)
	return ae
}

func (ae *ActionItemEditor) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(ae.content)
}
