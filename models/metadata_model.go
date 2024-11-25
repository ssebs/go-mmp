package models

import (
	"fmt"
	"log"
	"time"

	"go.bug.st/serial"
	"gopkg.in/yaml.v3"
)

type Metadata struct {
	Columns        int           `yaml:"Columns"`
	SerialPortName string        `yaml:"SerialPortName"`
	SerialBaudRate int           `yaml:"SerialBaudRate"`
	Delay          time.Duration `yaml:"Delay"`
	GUIMode        GUIMode       `yaml:"GUIMode"`
	Indexing       int           `yaml:"Indexing"`
}

func NewMetadata(
	portName string, baudRate int, guiMode GUIMode, cols int, delay time.Duration,
	indexing int,
) *Metadata {
	return &Metadata{
		Columns:        cols,
		SerialPortName: portName,
		SerialBaudRate: baudRate,
		Delay:          delay,
		GUIMode:        guiMode,
		Indexing:       indexing,
	}
}

func NewDefaultMetadata() *Metadata {
	m := &Metadata{
		Columns:        2,
		SerialPortName: "",
		SerialBaudRate: 9600,
		Delay:          125 * time.Millisecond,
		GUIMode:        GUIOnly,
		Indexing:       0,
	}

	// if ports, err := serial.GetPortsList(); err != nil {
	// 	// default to first port?
	// }

	return m
}

func (m *Metadata) UpdateAllFields(updated Metadata) {
	m.Columns = updated.Columns
	m.SerialPortName = updated.SerialPortName
	m.SerialBaudRate = updated.SerialBaudRate
	m.Delay = updated.Delay
	m.GUIMode = updated.GUIMode
	m.Indexing = updated.Indexing
}

func GetSerialPortsList() []string {
	ports, err := serial.GetPortsList()
	if err != nil {
		fmt.Println("error getting serial ports.", err)
		return nil
	}
	return ports
}

func (m Metadata) String() string {
	data, err := yaml.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}
