package serialdevice

import (
	"bufio"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/ssebs/go-mmp/models"
	"github.com/ssebs/go-mmp/utils"
	"go.bug.st/serial"
)

// SerialDevice is used to manage an arduino's serial connections
// From loading the port, to opening connections and listening for serial input
type SerialDevice struct {
	portName string
	mode     *serial.Mode
	Conn     serial.Port
	timeout  time.Duration
}

// Create a new SerialDevice
// Returns a SerialDevice, and an error.
// This will:
//   - set the serial port based on the given portName
//   - open the connection using the supplied baudRate, and set the timeout
func NewSerialDevice(pn string, baudRate int, timeout time.Duration) (*SerialDevice, error) {
	arduino := &SerialDevice{portName: pn, timeout: timeout}

	// err := arduino.SetSerialPort(pn) // not needed?
	// if err != nil {
	// 	return arduino, err
	// }
	err := arduino.OpenConnection(baudRate)
	if err != nil {
		return arduino, err
	}
	err = arduino.Conn.SetReadTimeout(timeout)
	if err != nil {
		return arduino, err
	}
	return arduino, nil
}

// Create a new SerialDevice from a Config struct
// Returns a SerialDevice, and an error.
// See NewSerialDevice.
func NewSerialDeviceFromConfig(c *models.Config, timeout time.Duration) (*SerialDevice, error) {
	arduino, err := NewSerialDevice(c.Metadata.SerialPortName, c.Metadata.SerialBaudRate, timeout)
	return arduino, err
}

// Open a serial connection based on the baudrate,
// and save the opened conn to SerialDevice.Conn
func (s *SerialDevice) OpenConnection(baud int) (err error) {
	s.mode = &serial.Mode{BaudRate: int(baud)}
	s.Conn, err = serial.Open(s.portName, s.mode)
	return err
}

// Close the serial connection
func (s *SerialDevice) CloseConnection() error {
	return s.Conn.Close()
}

// Find & set the SerialDevice portName field.
// Depends on what the requestedPortName is, and what serial devices are found.
func (s *SerialDevice) SetSerialPort(requestedPortName string) (err error) {
	// Get list of serial ports that are found
	ports, err := serial.GetPortsList()
	// TODO: replace with enumerator.GetDetailedPortsList
	if err != nil {
		return err
	}
	// If no serial devices are found
	if len(ports) == 0 {
		return ErrSerialDeviceNotFound{}
	}
	// Check if the requestedPortName matches any of the ports that were found
	if _, isFound := utils.SliceContains[string](&ports, requestedPortName); isFound {
		s.portName = requestedPortName
		return nil
	}
	// TODO: give the user option to use one of the listed ports
	return ErrSerialPortNameMismatch{Got: strings.Join(ports, ", "), Want: requestedPortName}
}

// Listen & run callback when data comes in
// Runs in a bufio.Scanner.Scan() loop
// callback must return true to break this loop
func (s *SerialDevice) ListenCallback(fn func(strData string) bool) (shouldBreak bool) {
	scanner := bufio.NewScanner(s.Conn)
	for scanner.Scan() {
		shouldBreak = fn(scanner.Text())
	}
	return shouldBreak
}

// Listen & send data thru chan
// Runs in a bufio.Scanner.Scan() loop
func (s *SerialDevice) ListenChan(ch chan string) {
	scanner := bufio.NewScanner(s.Conn)
	for scanner.Scan() {
		log.Println("ListenChan txt: ", scanner.Text())
		ch <- scanner.Text()
	}
}

// Listen & return data
// Runs in a bufio.Scanner.Scan() loop
func (s *SerialDevice) Listen() (actionID string, err error) {
	scanner := bufio.NewScanner(s.Conn)
	for scanner.Scan() {
		return scanner.Text(), nil
	}
	return "", scanner.Err()
}

// Errors
type ErrSerialDeviceNotFound struct{}
type ErrSerialPortNameMismatch struct {
	Got  string
	Want string
}

func (e ErrSerialDeviceNotFound) Error() string {
	return "serial device not found"
}

func (e ErrSerialPortNameMismatch) Error() string {
	return fmt.Sprintf(
		"%q not found in your serial devices, are you sure that's the right port?\n\nExpecting: %q, Found: %q\n\nPlease check your config file.", e.Want, e.Want, e.Got)
}
