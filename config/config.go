package config

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/ssebs/go-mmp/utils"
	"gopkg.in/yaml.v3"
)

// SerialDevice object
type SerialDevice struct {
	PortName string `yaml:"PortName"`
	BaudRate int    `yaml:"BaudRate"`
}

// MacroLayout object
type MacroLayout struct {
	SizeX  int `yaml:"SizeX"`
	SizeY  int `yaml:"SizeY"`
	Width  int `yaml:"Width"`
	Height int `yaml:"Height"`
}

// Macro object
type Macro struct {
	Name    string              `yaml:"Name"`
	Actions []map[string]string `yaml:"Actions"`
}

// Config object
// Stores related configuration details. No side effects here.
type Config struct {
	MacroLayout    MacroLayout   `yaml:"MacroLayout"`
	SerialDevice   SerialDevice  `yaml:"SerialDevice"`
	Macros         map[int]Macro `yaml:"Macros"`
	Delay          time.Duration `yaml:"Delay"`
	GUIMode        GUIMode       `yaml:"GUIMode"`
	ConfigFullPath string
}

// NewConfig takes in CLIFlags to figure out the correct path and whether or not to reset the file.
func NewConfig(flags *CLIFlags) (*Config, error) {
	c := &Config{}

	// Get the correct fullpath
	if err := c.getAndSetConfigPathFromCLIFlagsTODORename(flags); err != nil {
		return c, err
	}

	f, err := os.Open(c.ConfigFullPath)
	if err != nil {
		return c, err
	}
	defer f.Close()

	return loadConfig(f)
}

func (c *Config) getAndSetConfigPathFromCLIFlagsTODORename(flags *CLIFlags) error {

	// Get the fullpath of the default config
	hd, _ := os.UserHomeDir()
	defaultFullPath, err := filepath.Abs(filepath.Join(hd, ConfigPathShortName))
	if err != nil {
		return fmt.Errorf("failed to expand ${HOME}/mmpConfig.yml to a full path, %e", err)
	}

	// if the user doesn't set a --path arg
	if flags.ConfigPath == defaultFullPath {

		if !utils.CheckFileExists(defaultFullPath) {
			fmt.Printf("writing default config to %s\n", defaultFullPath)

			// TODO: Move / fix this!
			if err = utils.CopyFile("res/defaultConfig.yml", defaultFullPath); err != nil {
				return fmt.Errorf("failed to save defaultconfig. %e", err)
			}
		}

		c.ConfigFullPath = defaultFullPath
		return nil
	}

	// Get the full --path
	fullConfigPath, err := filepath.Abs(flags.ConfigPath)
	if err != nil {
		return fmt.Errorf("failed to expand %s to a full path, %e", flags.ConfigPath, err)
	}

	if !utils.CheckFileExists(fullConfigPath) {
		fmt.Printf("--path %s does not exist, writing default config to %s\n", fullConfigPath, fullConfigPath)

		if err = utils.CopyFile("res/defaultConfig.yml", fullConfigPath); err != nil {
			return fmt.Errorf("failed to save defaultconfig. %e", err)
		}
	}

	c.ConfigFullPath = fullConfigPath
	return nil
}

// ResetDefaultConfig will save the default config to ${HOME}/mmpConfig.yml
//
// TODO: Rename to resetConfig and use the path
// TODO: make sure copyfile works with fyne exporting, how do resources work??
func ResetDefaultConfig() error {
	hd, _ := os.UserHomeDir()
	defaultFullPath, err := filepath.Abs(filepath.Join(hd, ConfigPathShortName))
	if err != nil {
		return fmt.Errorf("failed to expand ${HOME}/mmpConfig.yml to a full path, %e", err)
	}

	// Copy file, if we get an error then return it
	if err := utils.CopyFile("res/defaultConfig.yml", defaultFullPath); err != nil {
		return fmt.Errorf("failed to save defaultconfig. %e", err)
	}
	return nil
}

// Create *Config from a io.Reader by marshalling the yaml to a Config
func loadConfig(f io.Reader) (*Config, error) {
	c := &Config{}
	err := yaml.NewDecoder(f).Decode(c)
	if err != nil {
		return c, err
	}

	// Set default delay if it's 0
	if c.Delay == 0 {
		c.Delay = 100 * time.Millisecond
	}

	return c, nil
}

// func (c Config) saveConfig(destFilename string) error {
// 	f, err := os.Create(destFilename)
// 	if err != nil {
// 		return err
// 	}
// 	_, err = f.WriteString(c.String())
// 	return err
// }

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
