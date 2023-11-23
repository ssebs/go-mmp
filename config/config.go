package config

import (
	"io"
	"log"
	"os"
	"time"

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
	MacroLayout  MacroLayout   `yaml:"MacroLayout"`
	SerialDevice SerialDevice  `yaml:"SerialDevice"`
	Macros       map[int]Macro `yaml:"Macros"`
	Delay        time.Duration `yaml:"Delay"`
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
