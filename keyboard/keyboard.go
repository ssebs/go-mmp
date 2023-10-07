package keyboard

import (
	"time"

	"github.com/micmonay/keybd_event"
)

// HotKeys
type HotKey int

const (
	SHIFT HotKey = iota
	CTRL  HotKey = iota
	ALT   HotKey = iota
	SUPER HotKey = iota
)

// HotKeyModifiers
type HotKeyModifiers struct {
	shift   bool
	control bool
	alt     bool
	super   bool
}

// GetActiveModifiers
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

func (k Keyboard) PressHold(duration time.Duration, keys ...int) error {
	k.kb.SetKeys(keys...)
	err := k.kb.Press()
	if err != nil {
		return err
	}
	time.Sleep(duration)
	err = k.kb.Release()
	if err != nil {
		return err
	}
	return nil
}

func (k Keyboard) HotKey(duration time.Duration, mods HotKeyModifiers, keys ...int) error {
	k.kb.SetKeys(keys...)

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
	return nil
}
