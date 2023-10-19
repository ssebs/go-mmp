package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// Config
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
	Macros []map[string]struct {
		ActionID int                 `yaml:"ActionID"`
		Actions  []map[string]string `yaml:"Actions"`
	} `yaml:"Macros"`
}

func (c *Config) String() string {
	data, err := yaml.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

// Create a new config
func NewConfigFromFile(filename string) (*Config, error) {
	config := &Config{}

	f, err := os.Open(filename)
	if err != nil {
		return config, err
	}
	defer f.Close()

	err = yaml.NewDecoder(f).Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}
