package macro

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
)

// macroaction.go is where the DoBlah functions are housed.
// macro.go was getting too hard to read, so I moved half of the methods here.

/*
 Below are the functions that provide the actual macro functionality
*/

// DoSendString will type a string that's passed
// param should be a string, this can be a letter, word, sentence, etc.
// e.g. DoSendTextAction("foo bar")
func (mm *MacroManager) DoSendTextAction(param string) error {
	// fmt.Println("DoSendTextAction:", param)
	robotgo.TypeStr(param)
	return nil
}

// DoPressReleaseAction will press and release a key or mouse button
// param should be a keyname or mouse btn name, see README.md
// e.g. DoPressReleaseAction("enter") or DoPressReleaseAction("RMB")
func (mm *MacroManager) DoPressReleaseAction(param string) error {
	// fmt.Println("DoPressReleaseAction:", param)
	// If it's a mouse button, pressMouse
	if isKeyNameMouseBtn(param) {
		return pressMouse(param)
	}
	// Otherwise, press the Key
	return pressKey(param)
}

// DoShortcutAction will type a shortcut out for you
// param should be formatted as: "SHIFT+ENTER+c"
// After keys are held down, there's a delay defined in mm.Config.Delay,
// then the keys will release
// This does NOT support mouse buttons
// e.g. DoShortcutAction("CTRL+SHIFT+ESC")
func (mm *MacroManager) DoShortcutAction(param string) error {
	// TODO: add option to delay between keydown/keyup
	keys := strings.Split(param, "+")

	// Hold down all keys
	for _, key := range keys {
		if err := robotgo.KeyDown(key); err != nil {
			return err
		}
	}
	// Delay
	time.Sleep(mm.Config.Delay)

	// Release all keys
	for _, key := range keys {
		if err := robotgo.KeyUp(key); err != nil {
			return err
		}
	}
	return nil
}

// DoRepeatAction will...
// param should be formatted as: "LMB+100ms"
// Only a single key and the delay between repeats should be in the string.
func (mm *MacroManager) DoRepeatAction(param string) error {
	// Generate keys from the string
	parts := strings.Split(param, "+")
	keyOrBtn := parts[0]

	// Assert that we have the correct params
	if len(parts) != 2 {
		// TODO: convert this to an error type
		return fmt.Errorf("config error: Repeat should only have 1 \"+\" between a keyname and the delay. expect format such as: \"LMB+100ms\", but got %s", param)
	}

	// Verify the delay is parsable, parse and save it
	repeatDelay, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("could not parse delay duration %q, err: %s", parts[1], err)
	}

	// If it's a mouse button, pressMouse
	if isKeyNameMouseBtn(keyOrBtn) {
		// goroutine will run until button is pressed again
		go repeatFunc(pressMouse, keyOrBtn, repeatDelay, mm.repeatStopCh)
	} else {
		// Otherwise, KeyPress
		go repeatFunc(pressKey, keyOrBtn, repeatDelay, mm.repeatStopCh)
	}

	// If isRepeating is set to true and this function is called again, close stopCh
	if mm.isRepeating {
		close(mm.repeatStopCh)
		mm.repeatStopCh = make(chan struct{})
	}
	mm.isRepeating = !mm.isRepeating
	return nil
}

// repeatFunc will run f() until stopCh is closed
// f is the function to run, param is the parameter to that function,
// repeatDelay is the delay between repeats in the loop,
// stopCh will break the loop when it's closed.
func repeatFunc(f fn, param string, repeatDelay time.Duration, stopCh chan struct{}) {
free:
	for {
		select {
		case <-stopCh:
			break free
		default:
			// Run fn over and over
			if err := f(param); err != nil {
				break free
			}
			// Sleep between repeats
			time.Sleep(repeatDelay)
		}
	}
}

// DoDelay will time.sleep for the delay if it can be parsed
func (mm *MacroManager) DoDelayAction(param string) error {
	delay, err := time.ParseDuration(param)
	if err != nil {
		return fmt.Errorf("could not parse delay duration %q, err: %s", param, err)
	}
	time.Sleep(delay)
	return nil
}

/* Helpers */

// pressKey will press and release a single key.
// If you want to press multiple, use robotgo.KeyPress
func pressKey(key string) error {
	return robotgo.KeyPress(key)
}

// pressMouse will press a mouse button down.
// "button" is the button to press.
// It can be either: LMB, RMB, or MMB
func pressMouse(button string) error {
	switch button {
	case "LMB", "LEFTMOUSE", "LEFTMOUSEBUTTON", "LEFTCLICK":
		robotgo.Click("left", false)
	case "RMB", "RIGHTMOUSE", "RIGHTMOUSEBUTTON", "RIGHTCLICK":
		robotgo.Click("right", false)
	case "MMB", "MIDDLEMOUSE", "MIDDLEMOUSEBUTTON", "MIDDLECLICK":
		robotgo.Click("center", false)
	}
	return nil
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
