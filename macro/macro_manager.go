package macro

import (
	"fmt"

	"github.com/ssebs/go-mmp/models"
	"github.com/ssebs/go-mmp/utils"
)

// For testing: Check out https://keyboard-test.space/

type MacroManager struct {
	*models.Config
	functionMap  map[string]func(string) error
	isRepeating  bool
	repeatStopCh chan struct{}
}

func NewMacroManager(conf *models.Config) (*MacroManager, error) {
	mgr := &MacroManager{
		Config:       conf,
		functionMap:  make(map[string]func(string) error),
		isRepeating:  false,
		repeatStopCh: make(chan struct{}),
	}

	// TODO: Organize with macro.FunctionList
	mgr.functionMap = map[string]func(string) error{
		"Delay":        mgr.DoDelayAction,
		"PressRelease": mgr.DoPressReleaseAction,
		"Press":        mgr.DoPressAction,
		"Release":      mgr.DoReleaseAction,
		"SendText":     mgr.DoSendTextAction,
		"Shortcut":     mgr.DoShortcutAction,
		"Repeat":       mgr.DoRepeatAction,
	}
	return mgr, nil
}

// Run the Actions in a Macro in order.
func (mm *MacroManager) RunMacro(macro *models.Macro) error {
	// Within a Macro, there's a list of Actions to run.
	// we want to run each one in order
	for _, action := range macro.Actions {
		if err := mm.runFuncFromMap(action.FuncName, action.FuncParam); err != nil {
			return fmt.Errorf("failed to run function from map, %s", err)
		}
	}
	return nil
}

// Run the Actions from the matching ID in Config.Macros in a Macro in order.
func (mm *MacroManager) RunMacroById(macroIdStr string) error {
	macroId, err := utils.StringToInt(macroIdStr)
	if err != nil {
		return fmt.Errorf("failed to convert macroID into int, %s", err)
	}
	// TODO: verbose mode
	fmt.Printf("Pressed: %d\n", macroId)

	matchedMacro, err := mm.Config.GetMacro(macroId)
	if err != nil {
		return fmt.Errorf("could not find macroId: %d in Macros %+v", macroId, mm.Config.Macros)
	}

	return mm.RunMacro(matchedMacro)
}

// runFuncFromMap runs the function from the functionMap if it exists, errors otherwise
func (mm *MacroManager) runFuncFromMap(funcName string, funcParams string) error {
	_, ok := mm.functionMap[funcName]
	if !ok {
		return fmt.Errorf("could not find %s in functionMap", funcName)
	}
	return mm.functionMap[funcName](funcParams)
}
