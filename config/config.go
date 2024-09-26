package config

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ssebs/go-mmp/utils"
	"gopkg.in/yaml.v3"
)

type GUIMode int

const (
	NOTSET GUIMode = iota // For use in comparing cli flags
	NORMAL                // Serial listener + GUI
	GUIOnly
	// CLIOnly
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
	MacroLayout  MacroLayout   `yaml:"MacroLayout"`
	SerialDevice SerialDevice  `yaml:"SerialDevice"`
	Macros       map[int]Macro `yaml:"Macros"`
	Delay        time.Duration `yaml:"Delay"`
	GUIMode      GUIMode       `yaml:"GUIMode"`
}

// TODO: Rewrite Config save/loading using io interfaces

// Save Config object to a destFilename
// Returns an error if one occurred
func (c Config) SaveConfig(destFilename string) error {
	f, err := os.Create(destFilename)
	if err != nil {
		return err
	}
	_, err = f.WriteString(c.String())
	return err
}

// Create *Config from a io.Reader
// Returns a pointer to the created Config, and an error if there is one
func LoadConfig(f io.Reader) (*Config, error) {
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

// Create a new config from a filename/path.
// Returns a pointer to the created Config, and an error if there is one
func NewConfigFromFile(filename string) (*Config, error) {
	f, err := os.Open(filename)
	if err != nil {
		return &Config{}, err
	}
	defer f.Close()
	return LoadConfig(f)
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

	// Get homePath string
	homePath, err := getHomeConfigPath()
	if err != nil {
		return "", err
	}

	// Check if homePath exists,
	if utils.CheckFileExists(homePath) {
		return homePath, nil
	} else {
		// if not, copy defaultconfig to homePath
		err = utils.CopyFile("res/defaultConfig.yml", homePath)
		if err != nil {
			return homePath, fmt.Errorf("could not save defaultConfig: %s", err)
		}
		return homePath, nil
	}
}

// ResetDefaultConfig will save the default config to ${HOME}/mmpConfig.yml
func ResetDefaultConfig() error {
	// Get homePath string
	homePath, err := getHomeConfigPath()
	if err != nil {
		return err
	}
	// Copy file, if we get an error then return it
	if err := utils.CopyFile("res/defaultConfig.yml", homePath); err != nil {
		return err
	}
	return nil
}

// getHomeConfigPath will generate the ${HOME}/mmpConfig.yml full path as a string
// Returns an error if we couldn't find the home dir.
func getHomeConfigPath() (string, error) {
	// TODO: don't use errors.New
	homeDir, err := os.UserHomeDir()
	if err != nil {
		// TODO: change this error so we can check for it.
		return "", errors.New("could not get user home dir: " + err.Error())
	}

	homePath := filepath.FromSlash(homeDir + "/mmpConfig.yml")
	return homePath, nil
}

/* Stringers */
// Return the Config as a yaml string
func (c *Config) String() string {
	data, err := yaml.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

// Return the Macro as a yaml string
func (c *Macro) String() string {
	data, err := yaml.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

// Return the SerialDevice as a yaml string
func (c *SerialDevice) String() string {
	data, err := yaml.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

// Return the MacroLayout as a yaml string
func (c *MacroLayout) String() string {
	data, err := yaml.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

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
