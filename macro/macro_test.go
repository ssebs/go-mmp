package macro

import (
	"fmt"
	"testing"
)

func TestConvertActionIDToInt(t *testing.T) {
	const testNum int = 2
	t.Run("test int in string", func(t *testing.T) {
		got, err := convertActionIDToInt(fmt.Sprintf("%d", testNum))
		if err != nil {
			t.Fatalf(err.Error())
		}
		if got != testNum {
			t.Fatalf("got %v want %v", got, testNum)
		}

	})
	t.Run("test float in string", func(t *testing.T) {
		got, err := convertActionIDToInt(fmt.Sprintf("%f", float32(testNum)))
		if err == nil {
			t.Fatalf("%d should have errored", got)
		}
		if got != -1 {
			t.Fatalf("%d should have been the same as %d", got, testNum)
		}
	})
	t.Run("test non-number in string", func(t *testing.T) {
		got, err := convertActionIDToInt("fail here")
		if err == nil {
			t.Fatalf("%d should have errored", got)
		}
		// TODO: test is
		if got != -1 {
			t.Fatalf("%d should have been the same as %d", got, testNum)
		}
	})
}

func TestRunActionFromID(t *testing.T) {
	// mgr, err := NewMacroManager("")
	// if err != nil {
	// 	t.Fatalf("Could not create NewMacroManager. err: %s", err)
	// }

	t.Run("test matching actionID", func(t *testing.T) {
		// TODO: test without running function
		// aka enable this again
		// err := mgr.RunActionFromID("1")
		// if err != nil {
		// 	t.Fatalf("did not expect err, got %s", err)
		// }
	})
	t.Run("test not found actionID", func(t *testing.T) {
		// err := mgr.RunActionFromID("999")
		// if err == nil {
		// 	t.Fail()
		// }
	})
	t.Run("test actionID is not int str type", func(t *testing.T) {})
}
