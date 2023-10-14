package mmp

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/ssebs/go-mmp/utils"
	"go.bug.st/serial"
)

// MMPSerialDevice
type MMPSerialDevice struct {
	portName string
	mode     *serial.Mode
	conn     serial.Port
}

func (s *MMPSerialDevice) loadConnection(baud int) (err error) {
	s.mode = &serial.Mode{BaudRate: int(baud)}
	s.conn, err = serial.Open(s.portName, s.mode)
	if err != nil {
		return err
	}
	return nil
}

func NewMMPSerialDevice(portName string, baudRate int, timeout time.Duration) (MMPSerialDevice, error) {
	arduino := MMPSerialDevice{portName: portName}

	ports, err := serial.GetPortsList()
	if err != nil {
		return arduino, err
	}

	fmt.Println("ports", ports)
	switch len(ports) {
	case 0:
		return arduino, ErrSerialDeviceNotFound{}
	case 1:
		if _, isFound := utils.SliceContains[string](&ports, portName); isFound {
			return arduino, arduino.loadConnection(baudRate)
		}
		log.Printf("%s not found in serial ports found: %v", portName, ports)
		arduino.portName = ports[1]
	default:
		if _, isFound := utils.SliceContains[string](&ports, portName); isFound {
			return arduino, arduino.loadConnection(baudRate)
		}
		// TODO: ask user which to choose if it doesn't match
		return arduino, ErrSerialPortNameMismatch{want: portName, got: strings.Join(ports, ", ")}
	}

	return arduino, serial.PortError{}
}

// Errors
type ErrSerialDeviceNotFound struct{}
type ErrSerialPortNameMismatch struct {
	got  string
	want string
}

func (e ErrSerialDeviceNotFound) Error() string {
	return "serial device not found"
}
func (e ErrSerialPortNameMismatch) Error() string {
	return fmt.Sprintf("portName not found in serial.GetPortsList. want %v, got %v", e.got, e.want)
}
