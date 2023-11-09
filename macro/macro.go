package macro

import (
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/micmonay/keybd_event"

	"github.com/ssebs/go-mmp/config"
	"github.com/ssebs/go-mmp/keyboard"
	"github.com/ssebs/go-mmp/utils"
)

// fn Function type, for use in functionMap
type fn func(string) error

// MacroManager
type MacroManager struct {
	Keeb        *keyboard.Keyboard
	Config      *config.Config
	functionMap map[string]fn
}

// Creates a new MacroManager struct
// Will load a Config from the configFilePath. If this is empty, load the default.
// Creates a Keyboard under Keeb
//
// To run a macro, use the RunActionFromID func
func NewMacroManager(configFilePath string) (*MacroManager, error) {
	// Create Keyboard
	kb, err := keyboard.NewKeyboard()
	if err != nil {
		return &MacroManager{}, err
	}
	// Create/Load Config
	if configFilePath == "" {
		configFilePath = "res/defaultConfig.yml"
	}
	conf, err := config.NewConfigFromFile(configFilePath)
	if err != nil {
		return &MacroManager{}, err
	}

	// Create the MacroManager
	// No reason for "4", just some rand size
	mgr := &MacroManager{Config: conf, Keeb: kb, functionMap: make(map[string]fn, 4)}
	mgr.initFunctionMap()
	return mgr, nil
}

// Create function map from string to actual method
func (mm *MacroManager) initFunctionMap() {
	mm.functionMap = map[string]fn{
		"TaskMgr":  mm.RunTaskManager,
		"PressKey": mm.PressKeyAction,
		"Shortcut": mm.RunShortcutAction,
	}
}

// Run the function from the functionMap if it exists
func (mm *MacroManager) runFuncFromMap(funcName string, funcParams string) error {
	_, ok := mm.functionMap[funcName]
	if !ok {
		return fmt.Errorf("could not find %s in mm.functionMap", funcName)
	}
	return mm.functionMap[funcName](funcParams)
}

// RunActionFromID - Run Actions from the matching ID in Config.Macros (loaded from yml)
// This converts the actionID to an int (if possible), if not then log the error
func (mm *MacroManager) RunActionFromID(actionID string) error {
	fmt.Printf("pressed: %s\n", actionID)
	iActionID, err := convertActionIDToInt(actionID)
	if err != nil {
		return err
	}

	// macro is a ptr to the config.Macros[iActionID] if it exists
	macro, ok := mm.Config.Macros[iActionID]
	if !ok {
		return ErrActionIDNotFoundInMacros{aID: iActionID, macros: mm.Config.Macros}
	}

	// For each action
	for _, action := range macro.Actions {
		// Get the key/vals from the action
		for funcName, funcParam := range action {
			// Try and run function
			err := mm.runFuncFromMap(funcName, funcParam)
			// Pass up error if there is one
			return err
		}
	}
	return nil
}

/*
 Below are the functions that provide the actual macro functionality
*/

// Open Task Manager by running CTRL + SHIFT + ESC
func (mm *MacroManager) RunTaskManager(param string) error {
	hkm := keyboard.HotKeyModifiers{Shift: true, Control: true}
	mm.Keeb.RunHotKey(10*time.Millisecond, hkm, keybd_event.VK_ESC)
	return nil
}

func (mm *MacroManager) RunShortcutAction(param string) error {
	// hkm keyboard.HotKeyModifiers, keys ...int
	// mm.Keeb.RunHotKey(10*time.Millisecond, hkm, keys...)
	return nil
}

// PressKeyAction
// keyName should be VK_ESC format
func (mm *MacroManager) PressKeyAction(keyName string) error {
	convertedName, err := keyboard.ConvertKeyName(keyName)
	if err != nil {
		return fmt.Errorf("could not press key: %s", keyName)
	}
	mm.Keeb.PressHold(mm.Config.Delay, convertedName)
	return nil
}

// PressKeysAction
// Needs a slice of keyName strings,
// keyNames should follow the same format at PressKeyAction
func (mm *MacroManager) PressKeysAction(keyNames []string) error {
	keys := make([]int, 0)
	for _, key := range keyNames {
		convertedName, err := keyboard.ConvertKeyName(key)
		if err != nil {
			fmt.Printf("could not press key: %s", key)
			continue
		}
		keys = append(keys, convertedName)
	}
	// TODO: Check if the key is a modifier key
	mm.Keeb.PressHold(mm.Config.Delay, keys...)
	return nil
}

/*
Below are helper functions
*/

// Do error handling if actionID is not actually an int
// checks for empty string error, returns -1 if there's an error.
func convertActionIDToInt(actionID string) (iActionID int, err error) {
	iActionID, err = utils.StringToInt(actionID)
	if errors.Is(err, &utils.ErrCannotParseIntFromEmptyString{}) {
		// do nothing if an empty string was passed
		return -1, err
	} else if err != nil {
		slog.Warn("convertActionIDToInt err: ", err)
		return -1, err
	}
	return iActionID, nil
}

/* Errors */
type ErrActionIDNotFoundInMacros struct {
	aID    int
	macros map[int]config.Macro
}

func (e ErrActionIDNotFoundInMacros) Error() string {
	return fmt.Sprintf("could not find actionID: %d in mm.Config.Macros %+v", e.aID, e.macros)
}
