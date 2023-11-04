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
	Keeb   *keyboard.Keyboard
	Config *config.Config
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

	mgr := &MacroManager{Config: config, Keeb: kb}
	return mgr, nil
}

// RunActionFromID
// This converts the actionID to an int (if possible),
// and runs a macro based on the actionID=> action mapping from the config
//
// TODO: Fix this comment
func RunActionFromID(actionID string, quitch chan struct{}) {
	iActionID, ok := convertActionIDToInt(actionID, quitch)
	if !ok {
		return
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

func (mm *MacroManager) RunShortcut(hkm keyboard.HotKeyModifiers, keys ...int) {
	mm.Keeb.RunHotKey(10*time.Millisecond, hkm, keys...)
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
