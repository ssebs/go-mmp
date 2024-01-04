package keyboard

import (
	"time"

	"git.tcp.direct/kayos/sendkeys"
	"github.com/go-vgo/robotgo"
	"github.com/micmonay/keybd_event"
)

// Keyboard
type Keyboard struct {
	KeyBonding *keybd_event.KeyBonding
	KBW        *sendkeys.KBWrap
}

// Create new Keyboard
func NewKeyboard() (*Keyboard, error) {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		return nil, err
	}
	wrap, err := sendkeys.NewKBWrapWithOptions()
	if err != nil {
		return nil, err
	}

	return &Keyboard{
		KeyBonding: &kb,
		KBW:        wrap,
	}, nil
}

// Press Mouse button. "button" is the button to press.
func (k *Keyboard) PressMouse(button string, isDouble bool) {
	switch button {
	case "LMB":
		robotgo.Click("left", isDouble)
	case "RMB":
		robotgo.Click("right", isDouble)
	case "MMB":
		robotgo.Click("center", isDouble)
	}
}

// Press and hold a set of keys, with a delayDuration between pressing and releasing
func (k *Keyboard) PressHold(delayDuration time.Duration, keys ...int) {
	// Set keys, press, sleep, release
	k.KeyBonding.SetKeys(keys...)
	k.KeyBonding.Press()
	time.Sleep(delayDuration)
	k.KeyBonding.Release()
	k.KeyBonding.Clear()
}

// PressRepeat will be called each time the button is pressed.
// After it's called, it should repeat pressing "keys" with delayDuration until it's pressed again.
func (k *Keyboard) PressRepeat(repeatDuration, delayDuration time.Duration, stopCh chan struct{}, keys ...int) {
	k.KeyBonding.SetKeys(keys...)

free:
	for {
		select {
		case <-stopCh:
			break free
		default:
			// Press the keys over and over
			k.KeyBonding.Press()
			time.Sleep(delayDuration)
			k.KeyBonding.Release()

			// Sleep between repeats
			time.Sleep(repeatDuration)
		}
	}
	k.KeyBonding.Clear()
}

// PressRepeatMouse will be called each time the button is pressed.
// After it's called, it should repeat pressing "keys" with delayDuration until it's pressed again.
func (k *Keyboard) PressRepeatMouse(repeatDuration time.Duration, stopCh chan struct{}, button string) {
free:
	for {
		select {
		case <-stopCh:
			break free
		default:
			// Press the button over and over
			// fmt.Println("Pressing", button)
			switch button {
			case "LMB":
				robotgo.Click("left", false)
			case "RMB":
				robotgo.Click("right", false)
			case "MMB":
				robotgo.Click("center", false)
			}
			// Sleep between repeats
			time.Sleep(repeatDuration)
		}
	}
}

// Press and hold a set of HotKeys, with a delayDuration between pressing and releasing
// Also supports adding other keys to press/release
func (k *Keyboard) RunHotKey(delayDuration time.Duration, mods *HotKeyModifiers, keys ...int) {
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

// RunSendString will type the keys with a random smol delay
func (k *Keyboard) RunSendString(keys string) error {
	return k.KBW.Type(keys)
}

/* HotKey + HotKeyMods stuff */

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

// NewHotKeyModifiers
func NewHotKeyModifiers(shift, control, alt, super bool) *HotKeyModifiers {
	h := &HotKeyModifiers{Shift: shift, Control: control, Alt: alt, Super: super}
	return h
}

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
