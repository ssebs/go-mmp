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

// MacroManager
type MacroManager struct {
	Keeb        *keyboard.Keyboard
	Config      *config.Config
	functionMap map[string]interface{}
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
	mgr := &MacroManager{Config: config, Keeb: kb, functionMap: make(map[string]interface{}, 4)}

	// TODO: something with this
	mgr.functionMap = map[string]interface{}{
		"TaskMgr":   mgr.RunTaskManager,
		"PressKey":  mgr.PressKeyAction,
		"PressKeys": mgr.PressKeysAction,
		"Shortcut":  mgr.RunShortcutAction,
	}
	return mgr, nil
}

// RunActionFromID - the thing that runs the thing
// This converts the actionID to an int (if possible),
// and runs a macro based on the actionID=> action mapping from the config
//
// TODO: Fix this comment
func (m *MacroManager) RunActionFromID(actionID string, quitch chan struct{}) {
	iActionID, ok := convertActionIDToInt(actionID, quitch)
	if !ok {
		return
	}

	// for each macro
	for macroName, macro := range m.Config.Macros {
		// for each action in macro
		for _, actions := range macro.Actions {
			// for each func/param in actions
			for funcName, funcParam := range actions {
				fmt.Printf("macroName: %s, macroID %d, func: %s(%s)\n", macroName, macro.ActionID, funcName, funcParam)
			}
		}

	}

	// do the mapping
	switch iActionID {
	case 9:
		fmt.Printf("quitting app")
		close(quitch)
	case 10:
		OpenTaskManager()
		fmt.Printf("pressed: %d\n", iActionID)
	default:
		fmt.Printf("pressed: %d\n", iActionID)
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

func (mm *MacroManager) RunTaskManager() {
	hkm := keyboard.HotKeyModifiers{Shift: true, Control: true}
	mm.Keeb.RunHotKey(10*time.Millisecond, hkm, keybd_event.VK_ESC)
}

func (mm *MacroManager) RunShortcutAction(hkm keyboard.HotKeyModifiers, keys ...int) {
	mm.Keeb.RunHotKey(10*time.Millisecond, hkm, keys...)
}

// PressKeyAction
// keyName should be VK_ESC format
func (mm *MacroManager) PressKeyAction(keyName string) {
	convertedName, err := keyboard.ConvertKeyName(keyName)
	if err != nil {
		fmt.Println("could not press key:", keyName)
		return
	}
	mm.Keeb.PressHold(mm.Config.Delay, convertedName)
}

// PressKeysAction
// Needs a slice of keyName strings,
// keyNames should follow the same format at PressKeyAction
func (mm *MacroManager) PressKeysAction(keyNames []string) {
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
