package utils

import (
	"fmt"

	flag "github.com/spf13/pflag"
)

type GUIMode int

const (
	CLIGUI GUIMode = iota
	CLIOnly
	GUIOnly
)

// CLI flag values will be stored in this
type CLIFlags struct {
	Mode        GUIMode
	ResetConfig bool
}

// parseFlags will parse the CLI flags that may have been used.
// Useful for disabling the serial listening functionality
func ParseFlags() CLIFlags {
	// TODO: Allow for GUI to pop up with a failed arduino connection... might not be possible with fyne
	// for now, let's use CLI flags
	cliFlags := CLIFlags{}
	flag.VarP(&cliFlags.Mode, "gui-mode", "g", "GUI Mode")
	// flag.IntVar(&int(cliFlags.GUIMode), "gui-only", fmt.Sprintf("Open %s in GUI Only Mode. Useful if you don't have a working arduino.", utils.ProjectName))
	flag.BoolVarP(&cliFlags.ResetConfig, "reset-config", "r", false, "Reset your ~/mmpConfig.yml file.")

	flag.Parse()
	return cliFlags
}

// String implements pflag.Value.
func (g *GUIMode) String() string {
	return fmt.Sprint(*g)
}

// Set implements pflag.Value.
func (g *GUIMode) Set(m string) error {
	switch m {
	case "CLIGUI":
		*g = CLIGUI
		return nil
	case "CLIOnly":
		*g = CLIOnly
		return nil
	case "GUIOnly":
		*g = GUIOnly
		return nil
	}
	return fmt.Errorf("could not find mode %s", m)
}

// Type implements pflag.Value.
func (g *GUIMode) Type() string {
	switch *g {
	case CLIGUI:
		return "CLIGUI"
	case CLIOnly:
		return "CLIOnly"
	case GUIOnly:
		return "GUIOnly"
	}
	return ""
}
