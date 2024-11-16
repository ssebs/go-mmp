package models

import (
	"log"
	"time"

	"gopkg.in/yaml.v3"
)

type Metadata struct {
	Columns        int           `yaml:"Columns"`
	SerialPortName string        `yaml:"SerialPortName"`
	SerialBaudRate int           `yaml:"SerialBaudRate"`
	Delay          time.Duration `yaml:"Delay"`
	GUIMode        GUIMode       `yaml:"GUIMode"`
}

func NewMetadata(
	portName string, baudRate int, guiMode GUIMode, cols int, delay time.Duration,
) *Metadata {
	return &Metadata{
		Columns:        cols,
		SerialPortName: portName,
		SerialBaudRate: baudRate,
		Delay:          delay,
		GUIMode:        guiMode,
	}
}

func NewDefaultMetadata() *Metadata {
	m := &Metadata{
		Columns:        2,
		SerialPortName: "",
		SerialBaudRate: 9600,
		Delay:          125 * time.Millisecond,
		GUIMode:        GUIOnly,
	}

	// if ports, err := serial.GetPortsList(); err != nil {
	// 	// default to first port?
	// }

	return m
}

func (m Metadata) String() string {
	data, err := yaml.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}
