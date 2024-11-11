package config

import (
	flag "github.com/spf13/pflag"
	"github.com/ssebs/go-mmp/models"
)

// Will be used as /home/user/mmpConfig.yml, or C:\Users\user\mmpConfig.yml
const ConfigPathShortName string = "mmpConfig.yml"

// CLI flag values will be stored in this
type CLIFlags struct {
	GUIMode     models.GUIMode
	ConfigPath  string
	ResetConfig bool
	Testing     bool
}

// parseFlags will parse the CLI flags that may have been used.
// Useful for disabling the serial listening functionality
func ParseFlags() *CLIFlags {
	cliFlags := &CLIFlags{}

	flag.VarP(&cliFlags.GUIMode, "mode", "m",
		"GUI Mode, defaults to 'NORMAL', use 'GUIOnly' to run without a serial device.")
	flag.BoolVarP(&cliFlags.ResetConfig, "reset-config", "r", false,
		"Reset your ~/mmpConfig.yml file to default. If using config-path, reset that file.")
	flag.StringVarP(&cliFlags.ConfigPath, "path", "p", ConfigPathShortName,
		"Path to your mmpConfig.yml. If used with reset-config, the specifified file will be reset.")
	flag.BoolVarP(&cliFlags.Testing, "dry-run", "n", false,
		"Enable dry run / test mode")
	// TODO: implement verbose flag

	flag.Parse()
	return cliFlags
}
