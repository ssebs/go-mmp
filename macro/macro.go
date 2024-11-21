package macro

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/ssebs/go-mmp/models"
	"github.com/ssebs/go-mmp/utils"
)

// macro.go is where the MacroManager is defined, and where the functionmap logic is setup.
// macroaction.go is where the actual macro functionality methods are housed.

// For testing: Check out https://keyboard-test.space/

// MacroManager creates a functionMap from the config, and is used to run Macros.
type MacroManager struct {
	*models.ConfigM
	functionMap  map[string]func(string) error
	isRepeating  bool
	repeatStopCh chan struct{}
}

// NewMacroManager creates a functionMap from the config, and is used to run Macros.
func NewMacroManager(conf *models.ConfigM) (*MacroManager, error) {
	mgr := &MacroManager{
		ConfigM:      conf,
		functionMap:  make(map[string]func(string) error),
		isRepeating:  false,
		repeatStopCh: make(chan struct{}),
	}

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

// // TODO: Merge with macro.FunctionList
// func (mm *MacroManager) GetFunctionMapActions() []string {
// 	keys := make([]string, 0, len(mm.functionMap))
// 	for k := range mm.functionMap {
// 		keys = append(keys, k)
// 	}

// 	return keys
// }

func (mm *MacroManager) RunActionFromStrID(actionID string) error {
	fmt.Printf("Pressed: %s\n", actionID)

	// Convert the button id to an int
	iActionID, err := convertActionID(actionID)
	if err != nil {
		return err
	}
	return mm.RunActionFromID(iActionID)
}

// RunActionFromID - Run Actions from the matching ID in Config.Macros (loaded from yml)
// This is the method to call a macro, if you want to *do* something, call this method.
// The macro must exist in the config, and the name must match the key in the function map.
// Use GetFunctionMapActions() to get a slice of function key names
//
// This converts the actionID to an int (if possible), if not then log the error
func (mm *MacroManager) RunActionFromID(actionID int) error {
	fmt.Printf("Pressed: %d\n", actionID)

	// matchedMacro is a ptr to the config.Macros[actionID] if it exists
	// this will have relevant info to call a method from.
	matchedMacro, err := mm.ConfigM.GetMacro(actionID)
	if err != nil {
		return fmt.Errorf("could not find actionID: %d in Macros %+v", actionID, mm.ConfigM.Macros)
	}

	// Within a Macro, there's a list of Actions to run.
	// we want to run each one in order here
	for _, action := range matchedMacro.Actions {
		if err := mm.runFuncFromMap(action.FuncName, action.FuncParam); err != nil {
			return fmt.Errorf("failed to run function from map, %s", err)
		}
	}
	return nil
}

// runFuncFromMap runs the function from the functionMap if it exists, errors otherwise
func (mm *MacroManager) runFuncFromMap(funcName string, funcParams string) error {
	_, ok := mm.functionMap[funcName]
	if !ok {
		return fmt.Errorf("could not find %s in mm.functionMap", funcName)
	}
	return mm.functionMap[funcName](funcParams)
}

// convertActionID converts a string to a BtnId (int)
// checks for empty string error, returns -1 if there's an error.
func convertActionID(actionID string) (int, error) {
	iActionID, err := utils.StringToInt(actionID)
	if errors.Is(err, &utils.ErrCannotParseIntFromEmptyString{}) {
		// do nothing if an empty string was passed
		return -1, err
	} else if err != nil {
		slog.Warn(fmt.Sprint("convertActionIDToInt err: ", err))
		return -1, err
	}
	return iActionID, nil
}
