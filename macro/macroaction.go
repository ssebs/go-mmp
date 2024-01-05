package macro

import (
	"fmt"
	"strings"
	"time"

	"github.com/micmonay/keybd_event"
	"github.com/ssebs/go-mmp/keyboard"
)

// macroaction.go is where the DoBlah functions are housed.
// macro.go was getting too hard to read, so I moved half of the methods here.

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
	switch err.(type) {
	case nil:
		mm.Keeb.PressHold(mm.Config.Delay, convertedName)
	case keyboard.ErrKeyNameIsMouseButton:
		mm.Keeb.PressMouse(keyName, false)
	default:
		return fmt.Errorf("could not press key: %s", keyName)
	}

	return nil
}

// DoRepeatKey converts the keyName & will press & repeat it until the button is pressed again
// keyName should be found in KeyMap
func (mm *MacroManager) DoRepeatKey(param string) error {
	// Generate keys from the string
	words := strings.Split(param, "+")

	repeatDelay, err := time.ParseDuration(words[1])
	if err != nil {
		return fmt.Errorf("could not parse delay duration %q, err: %s", words[1], err)
	}

	// Convert key name and see if it should be a button press
	iKey, err := keyboard.ConvertKeyName(words[0])
	if err == nil {
		// Run the function async until isRepeating is true & this func is called again
		go mm.Keeb.PressRepeat(repeatDelay, mm.Config.Delay, mm.repeatStopCh, iKey)
	} else if iKey == -2 {
		// TODO: Fix the error handling above! Use errors.As()
		// Run the function async until isRepeating is true & this func is called again
		go mm.Keeb.PressRepeatMouse(repeatDelay, mm.repeatStopCh, words[0])
	} else {
		return fmt.Errorf("could not press key: %s", words[0])
	}

	// If isRepeating is set to true and this function is called again, close stopCh
	if mm.isRepeating {
		close(mm.repeatStopCh)
		mm.repeatStopCh = make(chan struct{})
	}
	mm.isRepeating = !mm.isRepeating
	return nil
}
