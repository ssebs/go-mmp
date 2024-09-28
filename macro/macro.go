package macro

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/ssebs/go-mmp/config"
	"github.com/ssebs/go-mmp/utils"
)

// macro.go is where the MacroManager is defined, and where the functionmap logic is setup.
// macroaction.go is where the actual macro functionality methods are housed.

// For testing: Check out https://keyboard-test.space/

// MacroManager creates a functionMap from the config, and is used to run Macros.
// Use NewMacroManager to init!
type MacroManager struct {
	Config       *config.Config
	functionMap  map[string]fn
	isRepeating  bool
	repeatStopCh chan struct{}
}

// NewMacroManager Uses CLIFlags to load / reset a Config.
// If the config file is missing, copy the default one there.
func NewMacroManager(conf *config.Config) (*MacroManager, error) {
	// Create the MacroManager
	mgr := &MacroManager{
		Config:       conf,
		functionMap:  make(map[string]fn, 4), // No reason for "4", I rolled a dice
		isRepeating:  false,
		repeatStopCh: make(chan struct{}),
	}

	mgr.initFunctionMap()
	return mgr, nil
}

// initFunctionMap will create function map from string to actual method
// This is needed for RunActionFromID to work
// If you want to add a new macro type to use in the config, add it here too.
func (mm *MacroManager) initFunctionMap() {
	mm.functionMap = map[string]fn{
		"Delay":        mm.DoDelayAction,
		"PressRelease": mm.DoPressReleaseAction,
		"Press":        mm.DoPressAction,
		"Release":      mm.DoReleaseAction,
		"SendText":     mm.DoSendTextAction,
		"Shortcut":     mm.DoShortcutAction,
		"Repeat":       mm.DoRepeatAction,
	}
}

// RunActionFromID - Run Actions from the matching ID in Config.Macros (loaded from yml)
// This is the method to call a macro, if you want to *do* something, call this method.
// The macro must exist in the config, and the name must match the key in the function map.
// This converts the actionID to an int (if possible), if not then log the error
func (mm *MacroManager) RunActionFromID(actionID string) error {
	fmt.Printf("Pressed: %s\n", actionID)

	// Convert the button id to an int
	iActionID, err := convertActionIDToInt(actionID)
	if err != nil {
		return err
	}

	// matchedMacro is a ptr to the config.Macros[iActionID] if it exists
	// this will have relevant info to call a method from.
	matchedMacro, ok := mm.Config.Macros[iActionID]
	if !ok {
		return ErrActionIDNotFoundInMacros{aID: iActionID, macros: mm.Config.Macros}
	}

	// Within a Macro, there's a list of Actions to run.
	// we want to run each one in order here
	for _, action := range matchedMacro.Actions {
		// Get the key/vals from the action
		for funcName, funcParam := range action {
			// Run the function that was mapped, with the params given as a string
			if err := mm.runFuncFromMap(funcName, funcParam); err != nil {
				// Pass up error if there is one
				return err
			}
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

// convertActionIDToInt converts a string to an int
// checks for empty string error, returns -1 if there's an error.
func convertActionIDToInt(actionID string) (iActionID int, err error) {
	iActionID, err = utils.StringToInt(actionID)
	if errors.Is(err, &utils.ErrCannotParseIntFromEmptyString{}) {
		// do nothing if an empty string was passed
		return -1, err
	} else if err != nil {
		slog.Warn(fmt.Sprint("convertActionIDToInt err: ", err))
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

// fn Function type, for use in functionMap
type fn func(string) error
