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
	Shift   bool
	Control bool
	Alt     bool
	Super   bool
}

// Why doesn't intellisense sense this
// // NewHotKeyModifiers
// func (h *HotKeyModifiers) NewHotKeyModifiers(shift, control, alt, super bool) *HotKeyModifiers {
// 	h.shift = shift
// 	h.control = control
// 	h.alt = alt
// 	h.super = super
// 	return h
// }

// GetActiveModifiers
// Return active modifier HotKeys from self
func (h *HotKeyModifiers) GetActiveModifiers() []HotKey {
	activeKeys := []HotKey{}
	if h.Shift {
		activeKeys = append(activeKeys, SHIFT)
	}
	if h.Control {
		activeKeys = append(activeKeys, CTRL)
	}
	if h.Alt {
		activeKeys = append(activeKeys, ALT)
	}
	if h.Super {
		activeKeys = append(activeKeys, SUPER)
	}
	return activeKeys
}

// Keyboard
type Keyboard struct {
	KeyBonding *keybd_event.KeyBonding
}

// Create new Keyboard
func NewKeyboard() (*Keyboard, error) {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		return &Keyboard{}, err
	}
	return &Keyboard{&kb}, nil
}

// Press and hold a set of keys, with a delayDuration between pressing and releasing
func (k *Keyboard) PressHold(delayDuration time.Duration, keys ...int) {
	// Set keys, press, sleep, release
	k.KeyBonding.SetKeys(keys...)
	k.KeyBonding.Press()
	time.Sleep(delayDuration)
	k.KeyBonding.Release()
}

// Press and hold a set of HotKeys, with a delayDuration between pressing and releasing
// Also supports adding other keys to press/release
func (k *Keyboard) RunHotKey(delayDuration time.Duration, mods HotKeyModifiers, keys ...int) {
	// Press hotkey (combo of modifier key + other keys) e.g. CTRL+c

	// Set the modifiers to press in addition to the keys
	for _, hotkey := range mods.GetActiveModifiers() {
		switch hotkey {
		case SHIFT:
			k.KeyBonding.HasSHIFT(true)
		case CTRL:
			k.KeyBonding.HasCTRL(true)
		case ALT:
			k.KeyBonding.HasALT(true)
		case SUPER:
			k.KeyBonding.HasSuper(true)
		}
	}
	k.PressHold(delayDuration, keys...)
}
