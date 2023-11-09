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
	actionMap   map[int]config.Macro
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
	conf, err := config.NewConfigFromFile(configFilePath)
	if err != nil {
		log.Fatal(err)
	}

	// Create the MacroManager
	// No reason for "4", just some rand size
	mgr := &MacroManager{Config: conf, Keeb: kb, functionMap: make(map[string]fn, 4), actionMap: make(map[int]config.Macro)}

	mgr.initFunctionMap()
	mgr.initActionMap()
	return mgr, nil
}

// Create function map
func (mm *MacroManager) initFunctionMap() {
	mm.functionMap = map[string]fn{
		"TaskMgr":  mm.RunTaskManager,
		"PressKey": mm.PressKeyAction,
		"Shortcut": mm.RunShortcutAction,
	}
}

// Convert config Macro list into map based on the actionid
func (mm *MacroManager) initActionMap() {
	for _, macro := range mm.Config.Macros {
		aID := macro.ActionID
		mm.actionMap[aID] = macro
	}
}

// Run the function from the functionMap
func (mm *MacroManager) runFuncFromMap(funcName string, funcParams string) error {
	_, ok := mm.functionMap[funcName]
	if !ok {
		return fmt.Errorf("could not find %s in mm.functionMap", funcName)
	}
	return mm.functionMap[funcName](funcParams)
}

// Get the action from the actionMap
func (mm *MacroManager) getActionFromMap(actionID int) (*config.Macro, error) {
	action, ok := mm.actionMap[actionID]
	if !ok {
		return nil, fmt.Errorf("could not find actionID: %d in mm.actionMap %v+", actionID, mm.actionMap)
	}
	return &action, nil
}

// RunActionFromID - the thing that runs the thing
// This converts the actionID to an int (if possible),
// and runs a macro based on the actionID=> action mapping from the config
//
// TODO: Fix this comment + rename
func (mm *MacroManager) RunActionFromID(actionID string, quitch chan struct{}) {
	fmt.Printf("pressed: %s\n", actionID)
	iActionID, ok := convertActionIDToInt(actionID, quitch)
	if !ok {
		return
	}

	// macro is a ptr to the config.Macros[iActionID] if it exists
	macro, err := mm.getActionFromMap(iActionID)
	fmt.Println(macro)
	if err != nil {
		slog.Debug(err.Error())
		return
	}

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
