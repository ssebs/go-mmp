package config

import (
	"github.com/ssebs/go-mmp/gui"
	"github.com/ssebs/go-mmp/serialdevice"
)

// ButtonLayout
// X/y size for GUI elements that should match the macro pad
type ButtonLayout struct {
	X int
	Y int
}

// Config
// Manages configuration details of MMP
type Config struct {
	Layout ButtonLayout
	Gui    gui.GUI
	Sd     *serialdevice.SerialDevice
}

// Create a new config
func NewConfig(layout ButtonLayout, gui gui.GUI, sd *serialdevice.SerialDevice) *Config {
	return &Config{Layout: layout, Gui: gui, Sd: sd}
}
