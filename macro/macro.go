package macro

import (
	"errors"
	"fmt"
	"log"
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
// Also creates a Keyboard
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
	config, err := config.NewConfigFromFile(configFilePath)
	if err != nil {
		log.Fatal(err)
	}

	// Create the MacroManager
	mgr := &MacroManager{Config: config, Keeb: kb, functionMap: make(map[string]fn, 4)}

	// TODO: something with this
	mgr.functionMap = map[string]fn{
		"TaskMgr":  mgr.RunTaskManager,
		"PressKey": mgr.PressKeyAction,
		"Shortcut": mgr.RunShortcutAction,
	}
	return mgr, nil
}

func (mm *MacroManager) runFuncFromMap(funcName string, funcParams string) error {
	_, ok := mm.functionMap[funcName]
	if !ok {
		return fmt.Errorf("could not find %s in mm.functionMap", funcName)
	}

	return mm.functionMap[funcName](funcParams)
}

// RunActionFromID - the thing that runs the thing
// This converts the actionID to an int (if possible),
// and runs a macro based on the actionID=> action mapping from the config
//
// TODO: Fix this comment + rename
func (mm *MacroManager) RunActionFromID(actionID string, quitch chan struct{}) {
	iActionID, ok := convertActionIDToInt(actionID, quitch)
	if !ok {
		return
	}

	// TODO: move this to an init function...
	// TODO: make Macros a map instead & use switch so we can default to log

	// for each macro
	for _, macro := range mm.Config.Macros {
		// fmt.Printf("macro id: %d, name:%s, actions: %v+\n", macro.ActionID, macro.Name, macro.Actions)

		// If the user hit a button that's in the Macro's ActionID
		if macro.ActionID == iActionID {
			// For each action
			for _, action := range macro.Actions {
				// Get the key/vals from the action
				for funcName, funcParam := range action {
					// try and run function
					err := mm.runFuncFromMap(funcName, funcParam)
					if err != nil {
						slog.Debug(err.Error())
					}
				}
			}
		}

	}

	// // do the mapping
	// switch iActionID {
	// case 9:
	// 	fmt.Printf("quitting app")
	// 	close(quitch)
	// case 10:
	// 	OpenTaskManager()
	// 	fmt.Printf("pressed: %d\n", iActionID)
	// default:
	// 	fmt.Printf("pressed: %d\n", iActionID)
	// }
	fmt.Printf("pressed: %d\n", iActionID)
}

/*
 Below are the functions that provide the actual macro functionality
*/

// Open Task Manager by running CTRL + SHIFT + ESC
func OpenTaskManager() {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		slog.Warn(err.Error())
	}
	keeb := keyboard.Keyboard{KeyBonding: &kb}
	hkm := keyboard.HotKeyModifiers{Shift: true, Control: true}
	keeb.RunHotKey(10*time.Millisecond, hkm, keybd_event.VK_ESC)
}

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
func convertActionIDToInt(actionID string, quitch chan struct{}) (iActionID int, ok bool) {
	iActionID, err := utils.StringToInt(actionID)
	if errors.Is(err, &utils.ErrCannotParseIntFromEmptyString{}) {
		// do nothing if an empty string was passed
		return -1, false
	} else if err != nil {
		slog.Warn("convertActionIDToInt err: ", err)
		close(quitch)
		return -1, false
	}
	return iActionID, true
}
