package keyboard

import "testing"

func TestConvertKeyName(t *testing.T) {
	t.Run("test matching name from KeyMap", func(t *testing.T) {
		tests := []string{"ENTER", "enter", "shift", "a", "/", "ctrl", "ALT", " ", "DELETE", "backspace"}
		for _, tt := range tests {
			got, err := ConvertKeyName(tt)
			if err != nil {
				t.Fatalf("failed to convert %q to from the KeyMap. err: %s", tt, err)
			}
			if got == -1 {
				t.Fatalf("failed to convert %q to from the KeyMap. converted to %d. err: %s", tt, got, err)
			}
		}
	})

	t.Run("test invalid value from keyMap", func(t *testing.T) {
		tests := []string{"asdfghjkl", "321", "nil", "null"}
		for _, tt := range tests {
			got, err := ConvertKeyName(tt)
			if err == nil {
				t.Fatalf("should have recv'd an err. somehow converted %q to %d", tt, got)
			}
			if got != -1 {
				t.Fatalf("should have recv'd an err. somehow converted %q to %d", tt, got)
			}
		}
	})
}
