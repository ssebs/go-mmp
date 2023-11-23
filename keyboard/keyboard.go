package keyboard

import (
	"fmt"
	"strings"
	"time"

	"git.tcp.direct/kayos/sendkeys"
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
	return &Keyboard{&kb, wrap}, nil
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

// TODO: implement sendstring
func (k *Keyboard) RunSendString(delayDuration time.Duration, keys string) error {
	fmt.Println("typing: ", keys)
	return k.KBW.Type(keys)
}

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

// ConvertKeyName takes a key name as a string and returns the corresponding
// integer value from the KeyMap. If the key name is not found in the KeyMap,
// it returns an error.
func ConvertKeyName(keyName string) (int, error) {
	keyName = strings.ToUpper(keyName)
	// check for space
	val, ok := KeyMap[keyName]
	if !ok {
		return -1, fmt.Errorf("could not convert %s to keybd_event.%s", keyName, keyName)
	}
	return val, nil
}

// Keymap for string => int for keybd_evt
// Copypasta'd from github.com/micmonay/keybd_event/keybd_windows.go
var KeyMap = map[string]int{
	"VK_SP1":                 keybd_event.VK_SP1,
	"VK_SP2":                 keybd_event.VK_SP2,
	"VK_SP3":                 keybd_event.VK_SP3,
	"VK_SP4":                 keybd_event.VK_SP4,
	"VK_SP5":                 keybd_event.VK_SP5,
	"VK_SP6":                 keybd_event.VK_SP6,
	"VK_SP7":                 keybd_event.VK_SP7,
	"VK_SP8":                 keybd_event.VK_SP8,
	"VK_SP9":                 keybd_event.VK_SP9,
	"VK_SP10":                keybd_event.VK_SP10,
	"VK_SP11":                keybd_event.VK_SP11,
	"VK_SP12":                keybd_event.VK_SP12,
	"VK_ESC":                 keybd_event.VK_ESC,
	"VK_1":                   keybd_event.VK_1,
	"VK_2":                   keybd_event.VK_2,
	"VK_3":                   keybd_event.VK_3,
	"VK_4":                   keybd_event.VK_4,
	"VK_5":                   keybd_event.VK_5,
	"VK_6":                   keybd_event.VK_6,
	"VK_7":                   keybd_event.VK_7,
	"VK_8":                   keybd_event.VK_8,
	"VK_9":                   keybd_event.VK_9,
	"VK_0":                   keybd_event.VK_0,
	"VK_Q":                   keybd_event.VK_Q,
	"VK_W":                   keybd_event.VK_W,
	"VK_E":                   keybd_event.VK_E,
	"VK_R":                   keybd_event.VK_R,
	"VK_T":                   keybd_event.VK_T,
	"VK_Y":                   keybd_event.VK_Y,
	"VK_U":                   keybd_event.VK_U,
	"VK_I":                   keybd_event.VK_I,
	"VK_O":                   keybd_event.VK_O,
	"VK_P":                   keybd_event.VK_P,
	"VK_A":                   keybd_event.VK_A,
	"VK_S":                   keybd_event.VK_S,
	"VK_D":                   keybd_event.VK_D,
	"VK_F":                   keybd_event.VK_F,
	"VK_G":                   keybd_event.VK_G,
	"VK_H":                   keybd_event.VK_H,
	"VK_J":                   keybd_event.VK_J,
	"VK_K":                   keybd_event.VK_K,
	"VK_L":                   keybd_event.VK_L,
	"VK_Z":                   keybd_event.VK_Z,
	"VK_X":                   keybd_event.VK_X,
	"VK_C":                   keybd_event.VK_C,
	"VK_V":                   keybd_event.VK_V,
	"VK_B":                   keybd_event.VK_B,
	"VK_N":                   keybd_event.VK_N,
	"VK_M":                   keybd_event.VK_M,
	"VK_F1":                  keybd_event.VK_F1,
	"VK_F2":                  keybd_event.VK_F2,
	"VK_F3":                  keybd_event.VK_F3,
	"VK_F4":                  keybd_event.VK_F4,
	"VK_F5":                  keybd_event.VK_F5,
	"VK_F6":                  keybd_event.VK_F6,
	"VK_F7":                  keybd_event.VK_F7,
	"VK_F8":                  keybd_event.VK_F8,
	"VK_F9":                  keybd_event.VK_F9,
	"VK_F10":                 keybd_event.VK_F10,
	"VK_F11":                 keybd_event.VK_F11,
	"VK_F12":                 keybd_event.VK_F12,
	"VK_F13":                 keybd_event.VK_F13,
	"VK_F14":                 keybd_event.VK_F14,
	"VK_F15":                 keybd_event.VK_F15,
	"VK_F16":                 keybd_event.VK_F16,
	"VK_F17":                 keybd_event.VK_F17,
	"VK_F18":                 keybd_event.VK_F18,
	"VK_F19":                 keybd_event.VK_F19,
	"VK_F20":                 keybd_event.VK_F20,
	"VK_F21":                 keybd_event.VK_F21,
	"VK_F22":                 keybd_event.VK_F22,
	"VK_F23":                 keybd_event.VK_F23,
	"VK_F24":                 keybd_event.VK_F24,
	"VK_NUMLOCK":             keybd_event.VK_NUMLOCK,
	"VK_SCROLLLOCK":          keybd_event.VK_SCROLLLOCK,
	"VK_RESERVED":            keybd_event.VK_RESERVED,
	"VK_MINUS":               keybd_event.VK_MINUS,
	"VK_EQUAL":               keybd_event.VK_EQUAL,
	"VK_BACKSPACE":           keybd_event.VK_BACKSPACE,
	"VK_TAB":                 keybd_event.VK_TAB,
	"VK_LEFTBRACE":           keybd_event.VK_LEFTBRACE,
	"VK_RIGHTBRACE":          keybd_event.VK_RIGHTBRACE,
	"VK_ENTER":               keybd_event.VK_ENTER,
	"VK_SEMICOLON":           keybd_event.VK_SEMICOLON,
	"VK_APOSTROPHE":          keybd_event.VK_APOSTROPHE,
	"VK_GRAVE":               keybd_event.VK_GRAVE,
	"VK_BACKSLASH":           keybd_event.VK_BACKSLASH,
	"VK_COMMA":               keybd_event.VK_COMMA,
	"VK_DOT":                 keybd_event.VK_DOT,
	"VK_SLASH":               keybd_event.VK_SLASH,
	"VK_KPASTERISK":          keybd_event.VK_KPASTERISK,
	"VK_SPACE":               keybd_event.VK_SPACE,
	"VK_CAPSLOCK":            keybd_event.VK_CAPSLOCK,
	"VK_KP0":                 keybd_event.VK_KP0,
	"VK_KP1":                 keybd_event.VK_KP1,
	"VK_KP2":                 keybd_event.VK_KP2,
	"VK_KP3":                 keybd_event.VK_KP3,
	"VK_KP4":                 keybd_event.VK_KP4,
	"VK_KP5":                 keybd_event.VK_KP5,
	"VK_KP6":                 keybd_event.VK_KP6,
	"VK_KP7":                 keybd_event.VK_KP7,
	"VK_KP8":                 keybd_event.VK_KP8,
	"VK_KP9":                 keybd_event.VK_KP9,
	"VK_KPMINUS":             keybd_event.VK_KPMINUS,
	"VK_KPPLUS":              keybd_event.VK_KPPLUS,
	"VK_KPDOT":               keybd_event.VK_KPDOT,
	"VK_LBUTTON":             keybd_event.VK_LBUTTON,
	"VK_RBUTTON":             keybd_event.VK_RBUTTON,
	"VK_CANCEL":              keybd_event.VK_CANCEL,
	"VK_MBUTTON":             keybd_event.VK_MBUTTON,
	"VK_XBUTTON1":            keybd_event.VK_XBUTTON1,
	"VK_XBUTTON2":            keybd_event.VK_XBUTTON2,
	"VK_BACK":                keybd_event.VK_BACK,
	"VK_CLEAR":               keybd_event.VK_CLEAR,
	"VK_PAUSE":               keybd_event.VK_PAUSE,
	"VK_CAPITAL":             keybd_event.VK_CAPITAL,
	"VK_KANA":                keybd_event.VK_KANA,
	"VK_HANGUEL":             keybd_event.VK_HANGUEL,
	"VK_HANGUL":              keybd_event.VK_HANGUL,
	"VK_JUNJA":               keybd_event.VK_JUNJA,
	"VK_FINAL":               keybd_event.VK_FINAL,
	"VK_HANJA":               keybd_event.VK_HANJA,
	"VK_KANJI":               keybd_event.VK_KANJI,
	"VK_CONVERT":             keybd_event.VK_CONVERT,
	"VK_NONCONVERT":          keybd_event.VK_NONCONVERT,
	"VK_ACCEPT":              keybd_event.VK_ACCEPT,
	"VK_MODECHANGE":          keybd_event.VK_MODECHANGE,
	"VK_PAGEUP":              keybd_event.VK_PAGEUP,
	"VK_PAGEDOWN":            keybd_event.VK_PAGEDOWN,
	"VK_END":                 keybd_event.VK_END,
	"VK_HOME":                keybd_event.VK_HOME,
	"VK_LEFT":                keybd_event.VK_LEFT,
	"VK_UP":                  keybd_event.VK_UP,
	"VK_RIGHT":               keybd_event.VK_RIGHT,
	"VK_DOWN":                keybd_event.VK_DOWN,
	"VK_SELECT":              keybd_event.VK_SELECT,
	"VK_PRINT":               keybd_event.VK_PRINT,
	"VK_EXECUTE":             keybd_event.VK_EXECUTE,
	"VK_SNAPSHOT":            keybd_event.VK_SNAPSHOT,
	"VK_INSERT":              keybd_event.VK_INSERT,
	"VK_DELETE":              keybd_event.VK_DELETE,
	"VK_HELP":                keybd_event.VK_HELP,
	"VK_SCROLL":              keybd_event.VK_SCROLL,
	"VK_LMENU":               keybd_event.VK_LMENU,
	"VK_RMENU":               keybd_event.VK_RMENU,
	"VK_BROWSER_BACK":        keybd_event.VK_BROWSER_BACK,
	"VK_BROWSER_FORWARD":     keybd_event.VK_BROWSER_FORWARD,
	"VK_BROWSER_REFRESH":     keybd_event.VK_BROWSER_REFRESH,
	"VK_BROWSER_STOP":        keybd_event.VK_BROWSER_STOP,
	"VK_BROWSER_SEARCH":      keybd_event.VK_BROWSER_SEARCH,
	"VK_BROWSER_FAVORITES":   keybd_event.VK_BROWSER_FAVORITES,
	"VK_BROWSER_HOME":        keybd_event.VK_BROWSER_HOME,
	"VK_VOLUME_MUTE":         keybd_event.VK_VOLUME_MUTE,
	"VK_VOLUME_DOWN":         keybd_event.VK_VOLUME_DOWN,
	"VK_VOLUME_UP":           keybd_event.VK_VOLUME_UP,
	"VK_MEDIA_NEXT_TRACK":    keybd_event.VK_MEDIA_NEXT_TRACK,
	"VK_MEDIA_PREV_TRACK":    keybd_event.VK_MEDIA_PREV_TRACK,
	"VK_MEDIA_STOP":          keybd_event.VK_MEDIA_STOP,
	"VK_MEDIA_PLAY_PAUSE":    keybd_event.VK_MEDIA_PLAY_PAUSE,
	"VK_LAUNCH_MAIL":         keybd_event.VK_LAUNCH_MAIL,
	"VK_LAUNCH_MEDIA_SELECT": keybd_event.VK_LAUNCH_MEDIA_SELECT,
	"VK_LAUNCH_APP1":         keybd_event.VK_LAUNCH_APP1,
	"VK_LAUNCH_APP2":         keybd_event.VK_LAUNCH_APP2,
	"VK_OEM_1":               keybd_event.VK_OEM_1,
	"VK_OEM_PLUS":            keybd_event.VK_OEM_PLUS,
	"VK_OEM_COMMA":           keybd_event.VK_OEM_COMMA,
	"VK_OEM_MINUS":           keybd_event.VK_OEM_MINUS,
	"VK_OEM_PERIOD":          keybd_event.VK_OEM_PERIOD,
	"VK_OEM_2":               keybd_event.VK_OEM_2,
	"VK_OEM_3":               keybd_event.VK_OEM_3,
	"VK_OEM_4":               keybd_event.VK_OEM_4,
	"VK_OEM_5":               keybd_event.VK_OEM_5,
	"VK_OEM_6":               keybd_event.VK_OEM_6,
	"VK_OEM_7":               keybd_event.VK_OEM_7,
	"VK_OEM_8":               keybd_event.VK_OEM_8,
	"VK_OEM_102":             keybd_event.VK_OEM_102,
	"VK_PROCESSKEY":          keybd_event.VK_PROCESSKEY,
	"VK_PACKET":              keybd_event.VK_PACKET,
	"VK_ATTN":                keybd_event.VK_ATTN,
	"VK_CRSEL":               keybd_event.VK_CRSEL,
	"VK_EXSEL":               keybd_event.VK_EXSEL,
	"VK_EREOF":               keybd_event.VK_EREOF,
	"VK_PLAY":                keybd_event.VK_PLAY,
	"VK_ZOOM":                keybd_event.VK_ZOOM,
	"VK_NONAME":              keybd_event.VK_NONAME,
	"VK_PA1":                 keybd_event.VK_PA1,
	"VK_OEM_CLEAR":           keybd_event.VK_OEM_CLEAR,

	"_VK_SHIFT":           0x10 + 0xFFF,
	"_VK_CTRL":            0x11 + 0xFFF,
	"_VK_ALT":             0x12 + 0xFFF,
	"_VK_LSHIFT":          0xA0 + 0xFFF,
	"_VK_RSHIFT":          0xA1 + 0xFFF,
	"_VK_LCONTROL":        0xA2 + 0xFFF,
	"_VK_RCONTROL":        0xA3 + 0xFFF,
	"_VK_LWIN":            0x5B + 0xFFF,
	"_VK_RWIN":            0x5C + 0xFFF,
	"_KEYEVENTF_KEYUP":    0x0002,
	"_KEYEVENTF_SCANCODE": 0x0008,

	"SHIFT":           0x10 + 0xFFF,
	"CTRL":            0x11 + 0xFFF,
	"ALT":             0x12 + 0xFFF,
	"LSHIFT":          0xA0 + 0xFFF,
	"RSHIFT":          0xA1 + 0xFFF,
	"LCONTROL":        0xA2 + 0xFFF,
	"RCONTROL":        0xA3 + 0xFFF,
	"LWIN":            0x5B + 0xFFF,
	"WIN":             0x5B + 0xFFF,
	"SUPER":           0x5B + 0xFFF,
	"RWIN":            0x5C + 0xFFF,
	"EVENTF_KEYUP":    0x0002,
	"EVENTF_SCANCODE": 0x0008,

	" ": keybd_event.VK_SPACE,
	// "!": keybd_event.VK_,
	// "@": keybd_event.VK_At,
	// "#": keybd_event.VK_HAS,
	// "$": keybd_event.VK_DO,
	// "%": keybd_event.VK_P,
	// "^": keybd_event.VCAR,
	// "&": keybd_event.VK_AMP,
	"*": keybd_event.VK_KPASTERISK,
	// "(": keybd_event.VK_PAR,
	// ")": keybd_event.VK_SPACE,
	// "-": keybd_event.VK_HYp,
	// "_": keybd_event.VK_UND,
	// "=": keybd_event.VK_SPACE,
	// "+": keybd_event.VK_SPACE,
	// "[": keybd_event.VK_SPACE,
	// "]": keybd_event.VK_SPACE,
	// "{": keybd_event.VK_SPACE,
	// "}": keybd_event.VK_SPACE,
	// ";": keybd_event.VK_SPACE,
	// ":": keybd_event.VK_CO,
	// "'": keybd_event.VK_SPACE,
	// "\"": keybd_event.VK_SPACE,
	",": keybd_event.VK_COMMA,
	// "<": keybd_event.VK_SPACE,
	".": keybd_event.VK_OEM_PERIOD,
	// ">": keybd_event.VK_SPACE,
	"/": keybd_event.VK_SLASH,
	// "?": keybd_event,
	"\\": keybd_event.VK_BACKSLASH,
	// "|": keybd_event.VK_SPACE,
	// "`": keybd_event.VK_SPACE,
	// "~": keybd_event.VK_SPACE,
	// "\t": keybd_event.VK_SPACE,
	"SP1":                 keybd_event.VK_SP1,
	"SP2":                 keybd_event.VK_SP2,
	"SP3":                 keybd_event.VK_SP3,
	"SP4":                 keybd_event.VK_SP4,
	"SP5":                 keybd_event.VK_SP5,
	"SP6":                 keybd_event.VK_SP6,
	"SP7":                 keybd_event.VK_SP7,
	"SP8":                 keybd_event.VK_SP8,
	"SP9":                 keybd_event.VK_SP9,
	"SP10":                keybd_event.VK_SP10,
	"SP11":                keybd_event.VK_SP11,
	"SP12":                keybd_event.VK_SP12,
	"ESC":                 keybd_event.VK_ESC,
	"1":                   keybd_event.VK_1,
	"2":                   keybd_event.VK_2,
	"3":                   keybd_event.VK_3,
	"4":                   keybd_event.VK_4,
	"5":                   keybd_event.VK_5,
	"6":                   keybd_event.VK_6,
	"7":                   keybd_event.VK_7,
	"8":                   keybd_event.VK_8,
	"9":                   keybd_event.VK_9,
	"0":                   keybd_event.VK_0,
	"Q":                   keybd_event.VK_Q,
	"W":                   keybd_event.VK_W,
	"E":                   keybd_event.VK_E,
	"R":                   keybd_event.VK_R,
	"T":                   keybd_event.VK_T,
	"Y":                   keybd_event.VK_Y,
	"U":                   keybd_event.VK_U,
	"I":                   keybd_event.VK_I,
	"O":                   keybd_event.VK_O,
	"P":                   keybd_event.VK_P,
	"A":                   keybd_event.VK_A,
	"S":                   keybd_event.VK_S,
	"D":                   keybd_event.VK_D,
	"F":                   keybd_event.VK_F,
	"G":                   keybd_event.VK_G,
	"H":                   keybd_event.VK_H,
	"J":                   keybd_event.VK_J,
	"K":                   keybd_event.VK_K,
	"L":                   keybd_event.VK_L,
	"Z":                   keybd_event.VK_Z,
	"X":                   keybd_event.VK_X,
	"C":                   keybd_event.VK_C,
	"V":                   keybd_event.VK_V,
	"B":                   keybd_event.VK_B,
	"N":                   keybd_event.VK_N,
	"M":                   keybd_event.VK_M,
	"F1":                  keybd_event.VK_F1,
	"F2":                  keybd_event.VK_F2,
	"F3":                  keybd_event.VK_F3,
	"F4":                  keybd_event.VK_F4,
	"F5":                  keybd_event.VK_F5,
	"F6":                  keybd_event.VK_F6,
	"F7":                  keybd_event.VK_F7,
	"F8":                  keybd_event.VK_F8,
	"F9":                  keybd_event.VK_F9,
	"F10":                 keybd_event.VK_F10,
	"F11":                 keybd_event.VK_F11,
	"F12":                 keybd_event.VK_F12,
	"F13":                 keybd_event.VK_F13,
	"F14":                 keybd_event.VK_F14,
	"F15":                 keybd_event.VK_F15,
	"F16":                 keybd_event.VK_F16,
	"F17":                 keybd_event.VK_F17,
	"F18":                 keybd_event.VK_F18,
	"F19":                 keybd_event.VK_F19,
	"F20":                 keybd_event.VK_F20,
	"F21":                 keybd_event.VK_F21,
	"F22":                 keybd_event.VK_F22,
	"F23":                 keybd_event.VK_F23,
	"F24":                 keybd_event.VK_F24,
	"NUMLOCK":             keybd_event.VK_NUMLOCK,
	"SCROLLLOCK":          keybd_event.VK_SCROLLLOCK,
	"RESERVED":            keybd_event.VK_RESERVED,
	"MINUS":               keybd_event.VK_MINUS,
	"EQUAL":               keybd_event.VK_EQUAL,
	"BACKSPACE":           keybd_event.VK_BACKSPACE,
	"TAB":                 keybd_event.VK_TAB,
	"LEFTBRACE":           keybd_event.VK_LEFTBRACE,
	"RIGHTBRACE":          keybd_event.VK_RIGHTBRACE,
	"ENTER":               keybd_event.VK_ENTER,
	"SEMICOLON":           keybd_event.VK_SEMICOLON,
	"APOSTROPHE":          keybd_event.VK_APOSTROPHE,
	"GRAVE":               keybd_event.VK_GRAVE,
	"BACKSLASH":           keybd_event.VK_BACKSLASH,
	"COMMA":               keybd_event.VK_COMMA,
	"DOT":                 keybd_event.VK_DOT,
	"SLASH":               keybd_event.VK_SLASH,
	"KPASTERISK":          keybd_event.VK_KPASTERISK,
	"SPACE":               keybd_event.VK_SPACE,
	"CAPSLOCK":            keybd_event.VK_CAPSLOCK,
	"KP0":                 keybd_event.VK_KP0,
	"KP1":                 keybd_event.VK_KP1,
	"KP2":                 keybd_event.VK_KP2,
	"KP3":                 keybd_event.VK_KP3,
	"KP4":                 keybd_event.VK_KP4,
	"KP5":                 keybd_event.VK_KP5,
	"KP6":                 keybd_event.VK_KP6,
	"KP7":                 keybd_event.VK_KP7,
	"KP8":                 keybd_event.VK_KP8,
	"KP9":                 keybd_event.VK_KP9,
	"KPMINUS":             keybd_event.VK_KPMINUS,
	"KPPLUS":              keybd_event.VK_KPPLUS,
	"KPDOT":               keybd_event.VK_KPDOT,
	"LBUTTON":             keybd_event.VK_LBUTTON,
	"RBUTTON":             keybd_event.VK_RBUTTON,
	"CANCEL":              keybd_event.VK_CANCEL,
	"MBUTTON":             keybd_event.VK_MBUTTON,
	"XBUTTON1":            keybd_event.VK_XBUTTON1,
	"XBUTTON2":            keybd_event.VK_XBUTTON2,
	"BACK":                keybd_event.VK_BACK,
	"CLEAR":               keybd_event.VK_CLEAR,
	"PAUSE":               keybd_event.VK_PAUSE,
	"CAPITAL":             keybd_event.VK_CAPITAL,
	"KANA":                keybd_event.VK_KANA,
	"HANGUEL":             keybd_event.VK_HANGUEL,
	"HANGUL":              keybd_event.VK_HANGUL,
	"JUNJA":               keybd_event.VK_JUNJA,
	"FINAL":               keybd_event.VK_FINAL,
	"HANJA":               keybd_event.VK_HANJA,
	"KANJI":               keybd_event.VK_KANJI,
	"CONVERT":             keybd_event.VK_CONVERT,
	"NONCONVERT":          keybd_event.VK_NONCONVERT,
	"ACCEPT":              keybd_event.VK_ACCEPT,
	"MODECHANGE":          keybd_event.VK_MODECHANGE,
	"PAGEUP":              keybd_event.VK_PAGEUP,
	"PAGEDOWN":            keybd_event.VK_PAGEDOWN,
	"END":                 keybd_event.VK_END,
	"HOME":                keybd_event.VK_HOME,
	"LEFT":                keybd_event.VK_LEFT,
	"UP":                  keybd_event.VK_UP,
	"RIGHT":               keybd_event.VK_RIGHT,
	"DOWN":                keybd_event.VK_DOWN,
	"SELECT":              keybd_event.VK_SELECT,
	"PRINT":               keybd_event.VK_PRINT,
	"EXECUTE":             keybd_event.VK_EXECUTE,
	"SNAPSHOT":            keybd_event.VK_SNAPSHOT,
	"INSERT":              keybd_event.VK_INSERT,
	"DELETE":              keybd_event.VK_DELETE,
	"HELP":                keybd_event.VK_HELP,
	"SCROLL":              keybd_event.VK_SCROLL,
	"LMENU":               keybd_event.VK_LMENU,
	"RMENU":               keybd_event.VK_RMENU,
	"BROWSER_BACK":        keybd_event.VK_BROWSER_BACK,
	"BROWSER_FORWARD":     keybd_event.VK_BROWSER_FORWARD,
	"BROWSER_REFRESH":     keybd_event.VK_BROWSER_REFRESH,
	"BROWSER_STOP":        keybd_event.VK_BROWSER_STOP,
	"BROWSER_SEARCH":      keybd_event.VK_BROWSER_SEARCH,
	"BROWSER_FAVORITES":   keybd_event.VK_BROWSER_FAVORITES,
	"BROWSER_HOME":        keybd_event.VK_BROWSER_HOME,
	"VOLUME_MUTE":         keybd_event.VK_VOLUME_MUTE,
	"VOLUME_DOWN":         keybd_event.VK_VOLUME_DOWN,
	"VOLUME_UP":           keybd_event.VK_VOLUME_UP,
	"MEDIA_NEXT_TRACK":    keybd_event.VK_MEDIA_NEXT_TRACK,
	"MEDIA_PREV_TRACK":    keybd_event.VK_MEDIA_PREV_TRACK,
	"MEDIA_STOP":          keybd_event.VK_MEDIA_STOP,
	"MEDIA_PLAY_PAUSE":    keybd_event.VK_MEDIA_PLAY_PAUSE,
	"LAUNCH_MAIL":         keybd_event.VK_LAUNCH_MAIL,
	"LAUNCH_MEDIA_SELECT": keybd_event.VK_LAUNCH_MEDIA_SELECT,
	"LAUNCH_APP1":         keybd_event.VK_LAUNCH_APP1,
	"LAUNCH_APP2":         keybd_event.VK_LAUNCH_APP2,
	"OEM_1":               keybd_event.VK_OEM_1,
	"OEM_PLUS":            keybd_event.VK_OEM_PLUS,
	"OEM_COMMA":           keybd_event.VK_OEM_COMMA,
	"OEM_MINUS":           keybd_event.VK_OEM_MINUS,
	"OEM_PERIOD":          keybd_event.VK_OEM_PERIOD,
	"OEM_2":               keybd_event.VK_OEM_2,
	"OEM_3":               keybd_event.VK_OEM_3,
	"OEM_4":               keybd_event.VK_OEM_4,
	"OEM_5":               keybd_event.VK_OEM_5,
	"OEM_6":               keybd_event.VK_OEM_6,
	"OEM_7":               keybd_event.VK_OEM_7,
	"OEM_8":               keybd_event.VK_OEM_8,
	"OEM_102":             keybd_event.VK_OEM_102,
	"PROCESSKEY":          keybd_event.VK_PROCESSKEY,
	"PACKET":              keybd_event.VK_PACKET,
	"ATTN":                keybd_event.VK_ATTN,
	"CRSEL":               keybd_event.VK_CRSEL,
	"EXSEL":               keybd_event.VK_EXSEL,
	"EREOF":               keybd_event.VK_EREOF,
	"PLAY":                keybd_event.VK_PLAY,
	"ZOOM":                keybd_event.VK_ZOOM,
	"NONAME":              keybd_event.VK_NONAME,
	"PA1":                 keybd_event.VK_PA1,
	"OEM_CLEAR":           keybd_event.VK_OEM_CLEAR,
}
