package keyboard

import (
	"reflect"
	"testing"
)

func TestHotKeyModifiers(t *testing.T) {
	t.Run("test get active modifiers", func(t *testing.T) {
		mods := HotKeyModifiers{Shift: true, Control: true, Super: false}
		want := []HotKey{SHIFT, CTRL}
		got := mods.GetActiveModifiers()
		if !reflect.DeepEqual(want, got) {
			t.Errorf("want %s, got %s", want, got)
		}
		t.Logf("want %s, got %s", want, got)
	})
}
