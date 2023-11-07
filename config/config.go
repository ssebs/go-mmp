package config

import (
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// Config object
// Stores related configuration details. No side effects here.
type Config struct {
	MacroLayout struct {
		SizeX  int `yaml:"SizeX"`
		SizeY  int `yaml:"SizeY"`
		Width  int `yaml:"Width"`
		Height int `yaml:"Height"`
	} `yaml:"MacroLayout"`
	SerialDevice struct {
		PortName string `yaml:"PortName"`
		BaudRate int    `yaml:"BaudRate"`
	} `yaml:"SerialDevice"`
	Macros map[string]struct {
		ActionID int                 `yaml:"ActionID"`
		Actions  []map[string]string `yaml:"Actions"`
	} `yaml:"Macros"`
}

// Return the Config as a yaml string
func (c *Config) String() string {
	data, err := yaml.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

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
