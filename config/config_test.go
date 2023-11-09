package config

import (
	"bytes"
	"reflect"
	"testing"
)

var sampleYamlConfig string = `---
---
MacroLayout:
  SizeX: 3
  SizeY: 3
  Width: 300
  Height: 200
SerialDevice:
  PortName: COM7
  BaudRate: 9600
Macros:
  1:
    Name: Open Task Mgr
    ActionID: 1
    Actions:
      - TaskMgr: ""
  10:
    Name: Skip song
    Actions:
      - PressKey: VK_MEDIA_NEXT_TRACK
Delay: 20ms
`

func TestLoadConfig(t *testing.T) {
	t.Run("make sure loadconfig works", func(t *testing.T) {
		buff := bytes.Buffer{}
		buff.WriteString(sampleYamlConfig)

		got, err := LoadConfig(&buff)

		if err != nil {
			t.Fatal(err)
		}
		want, err := NewConfigFromFile("../res/defaultConfig.yml")
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v want %v", got, want)
		}
	})
}
