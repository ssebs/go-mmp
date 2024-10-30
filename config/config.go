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

//go:embed defaultConfig.yml
var defaultConfigFile []byte

// Config object
// Stores related configuration details.
type Config struct {
	MacroLayout    MacroLayout     `yaml:"MacroLayout"`
	SerialDevice   SerialDevice    `yaml:"SerialDevice"`
	Macros         map[BtnId]Macro `yaml:"Macros"`
	Delay          time.Duration   `yaml:"Delay"`
	GUIMode        GUIMode         `yaml:"GUIMode"`
	ConfigFullPath string
}
type BtnId int

// TODO: Add public methods for CRUD'ing config
// TO IMPLEMENT:
// [x] load
// [ ] save
// [ ] save as
// [ ] edit / update macros
// [ ] delete macros
// [ ] new macros

// NewConfig takes in CLIFlags to figure out the correct path and whether or not to reset the file.
func NewConfig(flags *CLIFlags) (*Config, error) {
	c := &Config{}

	if err := c.figureOutConfigPath(flags.ConfigPath); err != nil {
		return c, err
	}

	if flags.ResetConfig {
		if err := c.saveDefaultConfig(); err != nil {
			return c, fmt.Errorf("could not reset config, %e", err)
		}
	}

	if flags.GUIMode != c.GUIMode {
		c.GUIMode = flags.GUIMode
	}

	err := c.loadConfig()
	return c, err
}

func (c *Config) loadConfig() error {
	f, err := os.Open(c.ConfigFullPath)
	if err != nil {
		return fmt.Errorf("could not open %s, %e", c.ConfigFullPath, err)
	}
	defer f.Close()

	return parseConfig(f, c)
}

// func (c Config) saveConfig(destFilename string) error {
// 	f, err := os.Create(destFilename)
// 	if err != nil {
// 		return err
// 	}
// 	_, err = f.WriteString(c.String())
// 	return err
// }

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

// Save the defaultconfig at c.ConfigFullPath
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

/* Macro structs within config */
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
