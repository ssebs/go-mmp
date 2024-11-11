package main

import (
	"fmt"
	"os"

	"github.com/ssebs/go-mmp/config"
	"github.com/ssebs/go-mmp/gui"
	"github.com/ssebs/go-mmp/macro"
)

func main() {
	cliFlags := config.ParseFlags()
	conf, err := config.NewConfig(cliFlags)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}

	macroMgr, err := macro.NewMacroManager(conf)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}

	if cliFlags.GUIMode != config.NOTSET {
		macroMgr.Config.GUIMode = cliFlags.GUIMode
	}

	g := gui.NewGUI(macroMgr)

	g.EditConfig()
	g.App.Run()
}
