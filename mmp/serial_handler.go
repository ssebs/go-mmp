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
	Conn     serial.Port
}

func (s *MMPSerialDevice) loadConnection(baud int) (err error) {
	s.mode = &serial.Mode{BaudRate: int(baud)}
	s.Conn, err = serial.Open(s.portName, s.mode)
	return err
}

func NewMMPSerialDevice(portName string, baudRate int, timeout time.Duration) (arduino MMPSerialDevice, err error) {
	arduino.portName = portName

	// TODO: replace with enumerator.GetDetailedPortsList
	ports, err := serial.GetPortsList()
	if err != nil {
		return arduino, err
	}

	switch len(ports) {
	case 0:
		err = ErrSerialDeviceNotFound{}
	case 1:
		if _, isFound := utils.SliceContains[string](&ports, portName); isFound {
			err = arduino.loadConnection(baudRate)
			return arduino, err
		}
		log.Printf("%s not found in serial ports found: %v", portName, ports)
		arduino.portName = ports[1]
	default:
		if _, isFound := utils.SliceContains[string](&ports, portName); isFound {
			err = arduino.loadConnection(baudRate)
			return arduino, err
		}
		// TODO: ask user which to choose if it doesn't match
		err = ErrSerialPortNameMismatch{want: portName, got: strings.Join(ports, ", ")}
	}
	return arduino, err
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
