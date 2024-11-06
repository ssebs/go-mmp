package config

import (
	_ "embed"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/ssebs/go-mmp/utils"
	"gopkg.in/yaml.v3"
)

// TODO: SET GUIONLY IN CONFIG IF CLI FLAG IS SET ON DEFAULT
// TODO: Make UI elements listen for changes with config. See: observer?

//go:embed defaultConfig.yml
var defaultConfigFile []byte

//go:embed testConfig.yml
var testConfigFile []byte

// NewConfig takes in CLIFlags to figure out the correct path and whether or not to reset the file.
func NewConfig(flags *CLIFlags) (*Config, error) {
	c := &Config{}

	if err := c.figureOutConfigPath(flags.ConfigPath); err != nil {
		return c, err
	}

	if flags.ResetConfig {
		if err := c.saveDefaultConfig(flags.Testing); err != nil {
			return c, fmt.Errorf("could not reset config, %e", err)
		}
	}

	if flags.GUIMode != c.GUIMode {
		c.GUIMode = flags.GUIMode
	}

	err := c.loadConfig()
	return c, err
}

func NewMacro(name string, actions []map[string]string) Macro {
	m := Macro{
		Name:    name,
		Actions: actions,
	}
	if actions == nil {
		m.Actions = make([]map[string]string, 0)
	}

	return m
}

/* Macro / Config */
// Stores related configuration details.
type Config struct {
	MacroLayout    MacroLayout     `yaml:"MacroLayout"`
	SerialDevice   SerialDevice    `yaml:"SerialDevice"`
	Macros         map[BtnId]Macro `yaml:"Macros"`
	Delay          time.Duration   `yaml:"Delay"`
	GUIMode        GUIMode         `yaml:"GUIMode"`
	ConfigFullPath string          `yaml:"-"`
}
type BtnId int
type SerialDevice struct {
	PortName string `yaml:"PortName"`
	BaudRate int    `yaml:"BaudRate"`
}

type MacroLayout struct {
	SizeX  int `yaml:"SizeX"`
	SizeY  int `yaml:"SizeY"`
	Width  int `yaml:"Width"`
	Height int `yaml:"Height"`
}

type Macro struct {
	Name    string              `yaml:"Name"`
	Actions []map[string]string `yaml:"Actions"`
}

// load c.ConfigFullPath and parse it into c
// TODO: Support updating the main UI when this is called? Or use observer?
func (c *Config) loadConfig() error {
	f, err := os.Open(c.ConfigFullPath)
	if err != nil {
		return fmt.Errorf("could not open %s, %e", c.ConfigFullPath, err)
	}
	defer f.Close()

	return parseConfig(f, c)
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
			c.saveDefaultConfig(false)
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
		c.saveDefaultConfig(false)
	}
	return nil
}

// Save the defaultconfig at c.ConfigFullPath
func (c *Config) saveDefaultConfig(testEnabled bool) error {
	f, err := os.Create(c.ConfigFullPath)
	if err != nil {
		return fmt.Errorf("could not open file %s, %e", c.ConfigFullPath, err)
	}
	if testEnabled {
		_, err = f.Write(testConfigFile)
	} else {
		_, err = f.Write(defaultConfigFile)
	}

	if err != nil {
		return fmt.Errorf("could not write file, %e", err)
	}
	return err
}

// parse contents of f into c as Config
func parseConfig(f io.Reader, c *Config) error {
	err := yaml.NewDecoder(f).Decode(c)
	if err != nil {
		return fmt.Errorf("could not decode %s into a Config", f)
	}

	// Set default delay if it's 0
	if c.Delay == 0 {
		c.Delay = 100 * time.Millisecond
	}

	return nil
}

// get macro position from macro name, if not found return -1
func (c *Config) GetIdxFromMacro(m Macro) BtnId {
	for idx, macro := range c.Macros {
		if m.Name == macro.Name {
			return idx
		}
	}
	return -1
}

/* Stringers */
func (c *Config) String() string {
	data, err := yaml.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func (c *Macro) String() string {
	data, err := yaml.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func (c *SerialDevice) String() string {
	data, err := yaml.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func (c *MacroLayout) String() string {
	data, err := yaml.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}
