package macro

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/micmonay/keybd_event"
	"github.com/ssebs/go-mmp/keyboard"
)

// macroaction.go is where the DoBlah functions are housed.
// macro.go was getting too hard to read, so I moved half of the methods here.

/*
 Below are the functions that provide the actual macro functionality
*/

// MIGRATED TO ROBOTGO
// DoSendString will type a string that's passed
func (mm *MacroManager) DoSendTextAction(param string) error {
	// fmt.Println("DoSendTextAction:", param)
	robotgo.TypeStr(param)
	return nil
}

// DoPressReleaseAction will press and release a key or mouse button
func (mm *MacroManager) DoPressReleaseAction(param string) error {
	// fmt.Println("DoPressReleaseAction:", param)
	// If it's a mouse button, pressMouse
	if isKeyNameMouseBtn(param) {
		pressMouse(param, false)
		return nil
	}
	// Otherwise, KeyPress
	return robotgo.KeyPress(param)
}

// DoShortcutAction will type a shortcut
// param should be formatted as: "SHIFT+ENTER+c"
func (mm *MacroManager) DoShortcutAction(param string) error {
	// keys := strings.Split(param, "+")

	// We can't pass just keys...
	// KeyTap expects 1 key + args
	robotgo.KeySleep = int(mm.Config.Delay.Milliseconds() * 2)
	// robotgo.KeyTap(keys[0], keys[1:])
	robotgo.KeyTap("ctrl", "c")

	// OLD below

	// keymods := &keyboard.HotKeyModifiers{}
	// keys := make([]int, 0)
	// // Generate HotKeyModifiers from the string
	// for _, word := range strings.Split(param, "+") {
	// 	switch word {
	// 	case "SHIFT":
	// 		keymods.Shift = true
	// 	case "CTRL":
	// 		keymods.Control = true
	// 	case "ALT":
	// 		keymods.Alt = true
	// 	case "SUPER":
	// 		keymods.Super = true
	// 	default:
	// 		iKey, err := keyboard.ConvertKeyName(word)
	// 		if err != nil {
	// 			return fmt.Errorf("could not convert %s to keyboard int", word)
	// 		}
	// 		keys = append(keys, iKey)
	// 	}
	// }

	// // Run the macro
	// mm.Keeb.RunHotKey(mm.Config.Delay, keymods, keys...)
	return nil
}

// DoRepeatAction will...
func (mm *MacroManager) DoRepeatAction(param string) error {
	return nil
}

// OLD TO MIGRATE

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

// // DoShortcutAction will type a shortcut
// // param should be formatted as: "SHIFT+ENTER+c"
// func (mm *MacroManager) DoShortcutAction(param string) error {
// 	keymods := &keyboard.HotKeyModifiers{}
// 	keys := make([]int, 0)
// 	// Generate HotKeyModifiers from the string
// 	for _, word := range strings.Split(param, "+") {
// 		switch word {
// 		case "SHIFT":
// 			keymods.Shift = true
// 		case "CTRL":
// 			keymods.Control = true
// 		case "ALT":
// 			keymods.Alt = true
// 		case "SUPER":
// 			keymods.Super = true
// 		default:
// 			iKey, err := keyboard.ConvertKeyName(word)
// 			if err != nil {
// 				return fmt.Errorf("could not convert %s to keyboard int", word)
// 			}
// 			keys = append(keys, iKey)
// 		}
// 	}

// 	// Run the macro
// 	mm.Keeb.RunHotKey(mm.Config.Delay, keymods, keys...)
// 	return nil
// }

// DoRepeatKey converts the keyName & will press & repeat it until the button is pressed again
// keyName should be found in KeyMap
func (mm *MacroManager) DoRepeatKeyAction(param string) error {
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

// DoTaskManager - Open Task Manager by running CTRL + SHIFT + ESC
// Deprecated, just create a shortcut instead
func (mm *MacroManager) DoTaskManager(param string) error {
	hkm := &keyboard.HotKeyModifiers{Shift: true, Control: true}
	mm.Keeb.RunHotKey(mm.Config.Delay, hkm, keybd_event.VK_ESC)
	return nil
}

// MISC

// DoDelay will time.sleep for the delay given if it can be parsed
func (mm *MacroManager) DoDelayAction(param string) error {
	delay, err := time.ParseDuration(param)
	if err != nil {
		return fmt.Errorf("could not parse delay duration %q, err: %s", param, err)
	}
	time.Sleep(delay)
	return nil
}

/* Helpers */

// pressMouse will press a mouse button down.
// "button" is the button to press.
// It can be either: LMB, RMB, or MMB
// isDouble is if it's a double click
func pressMouse(button string, isDouble bool) {
	switch button {
	case "LMB", "LEFTMOUSE", "LEFTMOUSEBUTTON", "LEFTCLICK":
		robotgo.Click("left", isDouble)
	case "RMB", "RIGHTMOUSE", "RIGHTMOUSEBUTTON", "RIGHTCLICK":
		robotgo.Click("right", isDouble)
	case "MMB", "MIDDLEMOUSE", "MIDDLEMOUSEBUTTON", "MIDDLECLICK":
		robotgo.Click("center", isDouble)
	}
}

// isKeyNameMouseBtn checks if the keyname is a mouse button or not.
func isKeyNameMouseBtn(keyName string) bool {
	switch keyName {
	case "LMB", "LEFTMOUSE", "LEFTMOUSEBUTTON", "LEFTCLICK":
		return true
	case "RMB", "RIGHTMOUSE", "RIGHTMOUSEBUTTON", "RIGHTCLICK":
		return true
	case "MMB", "MIDDLEMOUSE", "MIDDLEMOUSEBUTTON", "MIDDLECLICK":
		return true
	default:
		return false
	}
}
