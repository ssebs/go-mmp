package macro

import (
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/micmonay/keybd_event"

	"github.com/ssebs/go-mmp/config"
	"github.com/ssebs/go-mmp/keyboard"
	"github.com/ssebs/go-mmp/utils"
)

// For testing: Check out https://keyboard-test.space/

// fn Function type, for use in functionMap
type fn func(string) error

// MacroManager
type MacroManager struct {
	Keeb        *keyboard.Keyboard
	Config      *config.Config
	functionMap map[string]fn
}

// NewMacroManager Creates a MacroManager with the given configuration file path.
// It initializes the keyboard and loads the configuration. If the configuration file
// path is empty, it uses the default configuration file.

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
			// fmt.Printf("funcName: %s, funcParam: %s\n", funcName, funcParam)
			if err := mm.runFuncFromMap(funcName, funcParam); err != nil {
				// Pass up error if there is one
				return err
			}
		}
	}

	//TODO: Show button pressed on gui
	return nil
}

// Create function map from string to actual method
func (mm *MacroManager) initFunctionMap() {
	mm.functionMap = map[string]fn{
		"TaskMgr":    mm.DoTaskManager,
		"Shortcut":   mm.DoShortcutAction,
		"SendString": mm.DoSendString,
		"Delay":      mm.DoDelay,
		"PressKey":   mm.DoPressKeyAction,
	}
}

/*
 Below are the functions that provide the actual macro functionality
*/

// Open Task Manager by running CTRL + SHIFT + ESC
func (mm *MacroManager) DoTaskManager(param string) error {
	hkm := &keyboard.HotKeyModifiers{Shift: true, Control: true}
	mm.Keeb.RunHotKey(mm.Config.Delay, hkm, keybd_event.VK_ESC)
	return nil
}

// DoShortcutAction will type a shortcut
// param should be formatted as: "SHIFT+ENTER+c"
func (mm *MacroManager) DoShortcutAction(param string) error {
	keymods := &keyboard.HotKeyModifiers{}
	keys := make([]int, 0)
	// Generate HotKeyModifiers from the string
	for _, word := range strings.Split(param, "+") {
		switch word {
		case "SHIFT":
			keymods.Shift = true
		case "CTRL":
			keymods.Control = true
		case "ALT":
			keymods.Alt = true
		case "SUPER":
			keymods.Super = true
		default:
			iKey, err := keyboard.ConvertKeyName(word)
			if err != nil {
				return fmt.Errorf("could not convert %s to keyboard int", word)
			}
			keys = append(keys, iKey)
		}
	}

	// Run the macro
	mm.Keeb.RunHotKey(mm.Config.Delay, keymods, keys...)
	return nil
}

// DoSendString will type a string that's passed
func (mm *MacroManager) DoSendString(param string) error {
	fmt.Println("RunSendString, ", param)
	return mm.Keeb.RunSendString(param)

}

// DoDelay will time.sleep for the delay given if it can be parsed
func (mm *MacroManager) DoDelay(param string) error {
	delay, err := time.ParseDuration(param)
	if err != nil {
		return fmt.Errorf("could not parse delay duration %q, err: %s", param, err)
	}
	time.Sleep(delay)
	return nil
}

// PressKeyAction converts the keyName & will press&hold it with mm.Config.Delay
// keyName should be found in KeyMap
func (mm *MacroManager) DoPressKeyAction(keyName string) error {
	convertedName, err := keyboard.ConvertKeyName(keyName)
	if err != nil {
		return fmt.Errorf("could not press key: %s", keyName)
	}

	mm.Keeb.PressHold(mm.Config.Delay, convertedName)
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
