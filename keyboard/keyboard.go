package keyboard

import (
	"time"

	"github.com/micmonay/keybd_event"
)

// Modifier Keys for HotKey
const (
	SHIFT HotKey = iota
	CTRL  HotKey = iota
	ALT   HotKey = iota
	SUPER HotKey = iota
)

// HotKey
type HotKey int

func (hk HotKey) String() string {
	return []string{"SHIFT", "CTRL", "ALT", "SUPER"}[hk]
}

// HotKeyModifiers
// When creating, set whatever modifier to true,
// then use GetActiveModifiers() to get a list of active HotKeys
type HotKeyModifiers struct {
	shift   bool
	control bool
	alt     bool
	super   bool
}

// GetActiveModifiers
// Return active modifier HotKeys from self
func (h *HotKeyModifiers) GetActiveModifiers() []HotKey {
	activeKeys := []HotKey{}
	if h.shift {
		activeKeys = append(activeKeys, SHIFT)
	}
	if h.control {
		activeKeys = append(activeKeys, CTRL)
	}
	if h.alt {
		activeKeys = append(activeKeys, ALT)
	}
	if h.super {
		activeKeys = append(activeKeys, SUPER)
	}
	return activeKeys
}

// Keyboard
type Keyboard struct {
	kb keybd_event.KeyBonding
}

func (k Keyboard) PressHold(delayDuration time.Duration, keys ...int) {
	// Set keys, press, sleep, release
	k.kb.SetKeys(keys...)
	k.kb.Press()
	time.Sleep(delayDuration)
	k.kb.Release()
}

func (k Keyboard) HotKey(delayDuration time.Duration, mods HotKeyModifiers, keys ...int) {
	// Press hotkey (combo of modifier key + other keys) e.g. CTRL+c
	k.kb.SetKeys(keys...)

	// Set the modifiers to press in addition to the keys
	for _, hotkey := range mods.GetActiveModifiers() {
		switch hotkey {
		case SHIFT:
			k.kb.HasSHIFT(true)
		case CTRL:
			k.kb.HasCTRL(true)
		case ALT:
			k.kb.HasALT(true)
		case SUPER:
			k.kb.HasSuper(true)
		}
	}
	k.PressHold(delayDuration)
}
