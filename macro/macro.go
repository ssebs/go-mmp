package macro

import (
	"log"
	"time"

	"github.com/micmonay/keybd_event"
	"github.com/ssebs/go-mmp/config"
	"github.com/ssebs/go-mmp/keyboard"
)

// MacroManager
type MacroManager struct {
	Keeb   *keyboard.Keyboard
	Config *config.Config
}

// Creates a new MacroManager struct
// Will load a Config from the configFilePath. If this is empty, load the default.
// Also creates a Keyboard
func NewMacroManager(configFilePath string) (*MacroManager, error) {
	// Create Keyboard
	kb, err := keyboard.NewKeyboard()
	if err != nil {
		return &MacroManager{}, err
	}
	// Create/Load Config
	if configFilePath == "" {
		configFilePath = "res/defaultConfig.yml"
	}
	config, err := config.NewConfigFromFile(configFilePath)
	if err != nil {
		log.Fatal(err)
	}

	mgr := &MacroManager{Config: config, Keeb: kb}
	return mgr, nil
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
