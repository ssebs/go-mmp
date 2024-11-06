package config

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
)

// TODO: Add public methods for CRUD'ing config
// TO IMPLEMENT:
// [x] load from CLI
// [ ] load from Open file
// [x] save
// [x] save as
// [ ] edit / update macros
// [ ] delete macros
// [ ] new macros

// Save config to file. If destFullPath is empty, use c.ConfigFullPath
func (c *Config) SaveConfig(destFullPath string) error {
	if destFullPath == "" {
		destFullPath = c.ConfigFullPath
	} else {
		filepath.Abs(destFullPath)
	}
	fmt.Println("Saving Config to", destFullPath)

	f, err := os.Create(destFullPath)
	if err != nil {
		return fmt.Errorf("could not write file %s, %e", destFullPath, err)
	}
	defer f.Close()

	_, err = f.Write([]byte(c.String()))
	if err != nil {
		return fmt.Errorf("could not write file, %e", err)
	}
	return err
}

// Open config from file.
func (c *Config) OpenConfig(srcFullPath string) error {
	if srcFullPath == "" {
		return fmt.Errorf("%s should not be empty", srcFullPath)
	}
	fmt.Println("Opening Config from", srcFullPath)

	fmt.Println("TODO: FIX OPEN CONFIG NOT SETTING MACROS PROPERLY")

	return c.loadConfig()
}

// Add macro to the end of the Macros map
func (c *Config) AddMacro(macro Macro) {
	c.Macros[BtnId(len(c.Macros)+1)] = macro
}

func (c *Config) DelMacro(macro Macro) {
	delete(c.Macros, c.GetIdxFromMacro(macro))
	panic("Fix this not deleting")
}

func (c *Config) AddMacroAction(macro Macro) {
	// add a new action to the end of the macro's actions list

	// Find the passed in macro so we can
}
