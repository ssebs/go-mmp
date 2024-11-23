package models

import (
	_ "embed"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"slices"
	"time"

	"github.com/ssebs/go-mmp/utils"
	"gopkg.in/yaml.v3"
)

type Config struct {
	*Metadata      `yaml:"Metadata"`
	Macros         []*Macro `yaml:"Macros"`
	ConfigFullPath string   `yaml:"-"`
}

func NewConfig(meta *Metadata, macros []*Macro) *Config {
	c := &Config{
		Metadata: meta,
		Macros:   macros,
	}

	if c.Macros == nil {
		c.Macros = make([]*Macro, 0)
	}

	return c
}

func NewConfigFromFile(flags *CLIFlags) (*Config, error) {
	c := &Config{}

	if err := c.figureOutConfigPath(flags.ConfigPath); err != nil {
		return c, err
	}

	if flags.ResetConfig {
		if err := c.saveDefaultConfig(); err != nil {
			return c, fmt.Errorf("could not reset config, %e", err)
		}
	}

	if c.Metadata != nil && flags.GUIMode != c.Metadata.GUIMode {
		c.Metadata.GUIMode = flags.GUIMode
	}

	err := c.loadConfig()
	return c, err
}

// depending on CLI args, and what files already exist, save default config if needed, and set c.ConfigFullPath
func (c *Config) figureOutConfigPath(configPath string) error {
	// Get the fullpath of the default config
	hd, _ := os.UserHomeDir()
	defaultFullPath, err := filepath.Abs(filepath.Join(hd, ConfigPathShortName))
	if err != nil {
		return fmt.Errorf("failed to expand ${HOME}/mmpConfig.yml to a full path, %e", err)
	}

	// if the user doesn't set a --path arg, save default.
	// ConfigPathShortName is the default if no arg is specified
	if configPath == ConfigPathShortName {
		c.ConfigFullPath = defaultFullPath

		if !utils.CheckFileExists(defaultFullPath) {
			fmt.Printf("writing default config to %s\n", defaultFullPath)
			c.saveDefaultConfig()
		}
		return nil
	}

	// Get the full --path
	fullConfigPath, err := filepath.Abs(configPath)
	if err != nil {
		return fmt.Errorf("failed to expand %s to a full path, %e", configPath, err)
	}

	c.ConfigFullPath = fullConfigPath
	if !utils.CheckFileExists(fullConfigPath) {
		fmt.Printf("--path %s does not exist, writing default config.\n", fullConfigPath)
		c.saveDefaultConfig()
	}
	return nil
}

// parse contents of f into c as ConfigM
func parseConfig(f io.Reader, c *Config) error {
	err := yaml.NewDecoder(f).Decode(c)
	if err != nil {
		return fmt.Errorf("could not decode %s into a Config. err: %s", f, err)
	}

	// Set default delay if it's 0
	if c.Delay == 0 {
		c.Delay = 100 * time.Millisecond
	}

	return nil
}

// load c.ConfigFullPath and parse it into c
func (c *Config) loadConfig() error {
	f, err := os.Open(c.ConfigFullPath)
	if err != nil {
		return fmt.Errorf("could not open %s, %e", c.ConfigFullPath, err)
	}
	defer f.Close()

	return parseConfig(f, c)
}

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
	c.ConfigFullPath = srcFullPath

	return c.loadConfig()
}

//go:embed defaultConfig.yml
var defaultConfigFile []byte

// Write the defaultconfig to c.ConfigFullPath
func (c *Config) saveDefaultConfig() error {
	f, err := os.Create(c.ConfigFullPath)
	if err != nil {
		return fmt.Errorf("could not open file %s, %e", c.ConfigFullPath, err)
	}

	_, err = f.Write(defaultConfigFile)

	if err != nil {
		return fmt.Errorf("could not write file, %e", err)
	}
	return err
}

// MVC functions
func (c *Config) AddMacro(m *Macro) {
	c.Macros = append(c.Macros, m)
}

func (c *Config) GetMacro(idx int) (*Macro, error) {
	if !c.isValidBoundsInMacros(idx) {
		return nil, fmt.Errorf("idx out of bounds of Macros")
	}

	return c.Macros[idx], nil
}

func (c *Config) UpdateMacro(idx int, updatedMacro *Macro) error {
	if !c.isValidBoundsInMacros(idx) {
		return fmt.Errorf("idx out of bounds of Macros")
	}

	c.Macros[idx] = updatedMacro
	return nil
}

func (c *Config) DeleteMacro(idx int) error {
	if !c.isValidBoundsInMacros(idx) {
		return fmt.Errorf("idx out of bounds of Macros")
	}
	c.Macros = slices.Delete(c.Macros, idx, idx+1)
	return nil
}

func (c *Config) SwapMacroPositions(idx1, idx2 int) error {
	if !c.isValidBoundsInMacros(idx1) || !c.isValidBoundsInMacros(idx2) {
		return fmt.Errorf("idx out of bounds of Macro's actions")
	}

	c.Macros[idx1], c.Macros[idx2] = c.Macros[idx2], c.Macros[idx1]
	return nil
}

func (c *Config) isValidBoundsInMacros(idx int) bool {
	return idx <= len(c.Macros) && idx >= 0
}

func (c Config) String() string {
	data, err := yaml.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}
