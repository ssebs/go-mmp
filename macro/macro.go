package macro

import (
	"log"
	"time"

	"github.com/micmonay/keybd_event"
	"github.com/ssebs/go-mmp/keyboard"
)

// Macro hold specific macro info
type Macro struct {
	Title    string
	Callback func(v ...any)
}

// MacroManager holds macro data
type MacroManager struct {
	Macros []Macro
	Keeb   keyboard.Keyboard
}

// Open Task Manager by running CTRL + SHIFT + ESC
func OpenTaskManager() {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		log.Println(err)
	}
	keeb := keyboard.Keyboard{KeyBonding: &kb}
	hkm := keyboard.HotKeyModifiers{Shift: true, Control: true}
	keeb.RunHotKey(10*time.Millisecond, hkm, keybd_event.VK_ESC)
}

func (mm *MacroManager) RunShortcut(hkm keyboard.HotKeyModifiers, keys ...int) {
	mm.Keeb.RunHotKey(10*time.Millisecond, hkm, keys...)
}
