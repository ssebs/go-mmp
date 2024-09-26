package config

import (
	"fmt"
	"strings"

	flag "github.com/spf13/pflag"
)

// CLI flag values will be stored in this
type CLIFlags struct {
	GUIMode     GUIMode
	ConfigPath  string // fullpath
	ResetConfig bool
}

// parseFlags will parse the CLI flags that may have been used.
// Useful for disabling the serial listening functionality
func ParseFlags() *CLIFlags {
	cliFlags := &CLIFlags{}

	flag.VarP(&cliFlags.GUIMode, "mode", "m", "GUI Mode, defaults to 'NORMAL', use 'GUIOnly' to run without a serial device.")
	flag.BoolVarP(&cliFlags.ResetConfig, "reset-config", "r", false, "Reset your ~/mmpConfig.yml file to default. If using config-path, reset that file.")
	// TODO: implement configPath flag
	// TODO: implement verbose flag

	flag.Parse()
	return cliFlags
}

type GUIMode int

const (
	NOTSET GUIMode = iota // For use in comparing cli flags
	NORMAL                // Serial listener + GUI
	GUIOnly
	// CLIOnly
)

/* GUIMode pflag.Value implementation */
func (g *GUIMode) String() string {
	return fmt.Sprint(*g)
}
func (g *GUIMode) Type() string {
	switch *g {
	case NORMAL:
		return "NORMAL"
	case GUIOnly:
		return "GUIOnly"
	}
	return ""
}
func (g *GUIMode) Set(m string) error {
	switch strings.ToUpper(m) {
	case "NORMAL":
		*g = NORMAL
		return nil
	case "GUIONLY":
		*g = GUIOnly
		return nil
	}
	return fmt.Errorf("could not find mode %s", m)
}
func (g *GUIMode) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var modeString string
	if err := unmarshal(&modeString); err != nil {
		return err
	}

	switch strings.ToUpper(modeString) {
	case "NORMAL":
		*g = NORMAL
	case "GUIONLY":
		*g = GUIOnly
	default:
		return fmt.Errorf("invalid GUIMode: %s", modeString)
	}
	return nil
}
