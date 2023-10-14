package mmp

import (
	"log"
	"time"

	"github.com/micmonay/keybd_event"
	"github.com/ssebs/go-mmp/keyboard"
)

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
