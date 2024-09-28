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
	configFullPath string
}

// NewConfig takes in CLIFlags to figure out the correct path and whether or not to reset the file.
func NewConfig(flags *CLIFlags) (*Config, error) {
	c := &Config{}

	c.getAndSetConfigPathFromCLIFlagsTODORename()

	return c, nil

	// old
	// func NewConfigFromFile(filename string) (*Config, error) {
	// 	f, err := os.Open(filename)
	// 	if err != nil {
	// 		return &Config{}, err
	// 	}
	// 	defer f.Close()
	// 	return loadConfig(f)
	// }
}

func (c *Config) getAndSetConfigPathFromCLIFlagsTODORename() {

}

// GetConfigFilePath will...
// 1) Check if there's a neighboring mmpConfig.yml file
// 2) Check if there's a ${HOME}/mmpConfig.yml
// 3) Create default config at ${HOME}/mmpConfig.yml
// 4) Return the full filepath as a string
// If there's an error, return empty string and error.
func GetConfigFilePath() (string, error) {
	// Check for local ./mmpConfig.yml
	if utils.CheckFileExists("./mmpConfig.yml") {
		p, err := os.Getwd()
		if err != nil {
			return "", fmt.Errorf("could not get cwd: %s", err)
		}
		return filepath.FromSlash(p + "/mmpConfig.yml"), nil
	}

	// TODO: REVIEW THIS LOGIC

	// Get homeConfigPath string
	homeDir, _ := os.UserHomeDir()
	homeConfigPath := filepath.FromSlash(homeDir + "/mmpConfig.yml")

	if utils.CheckFileExists(homeConfigPath) {
		return homeConfigPath, nil
	} else {
		ResetDefaultConfig()
		return homeConfigPath, nil
	}
}

// ResetDefaultConfig will save the default config to ${HOME}/mmpConfig.yml
//
// TODO: Rename to resetConfig and use the path
// TODO: make sure copyfile works with fyne exporting, how do resources work??
func ResetDefaultConfig() error {
	homeDir, _ := os.UserHomeDir()
	homeConfigPath := filepath.FromSlash(homeDir + "/mmpConfig.yml")

	// Copy file, if we get an error then return it
	if err := utils.CopyFile("res/defaultConfig.yml", homeConfigPath); err != nil {
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

func (c Config) saveConfig(destFilename string) error {
	f, err := os.Create(destFilename)
	if err != nil {
		return err
	}
	_, err = f.WriteString(c.String())
	return err
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
