package serialdevice

import (
	"bufio"
	"fmt"
	"log/slog"
	"time"

	"github.com/ssebs/go-mmp/models"
	"go.bug.st/serial"
)

// SerialDevice is used to manage an arduino's serial connections
type SerialDevice struct {
	portName string
	mode     *serial.Mode
	Conn     serial.Port
}

// This will Open a connection to the portName with baudRate
func NewSerialDevice(portName string, baudRate int, timeout time.Duration) (*SerialDevice, error) {
	serialDevice := &SerialDevice{
		portName: portName,
		mode:     &serial.Mode{BaudRate: baudRate},
		Conn:     nil,
	}

	conn, err := serial.Open(serialDevice.portName, serialDevice.mode)
	if err != nil {
		return serialDevice, fmt.Errorf("failed to open serial %s, err: %s", portName, err)
	}
	serialDevice.Conn = conn

	err = serialDevice.Conn.SetReadTimeout(timeout)
	if err != nil {
		return serialDevice, fmt.Errorf("failed to set read timeout, err: %s", err)
	}

	return serialDevice, nil
}

func NewSerialDeviceFromConfig(c *models.Config, timeout time.Duration) (*SerialDevice, error) {
	arduino, err := NewSerialDevice(c.Metadata.SerialPortName, c.Metadata.SerialBaudRate, timeout)
	return arduino, err
}

func (s *SerialDevice) CloseConnection() error {
	return s.Conn.Close()
}

// Listen for data from a *SerialDevice, to be used in a goroutine
// Takes in a btnch to send data to when the serial connection gets something,
// and a quitch if we need to stop the goroutine
func (s *SerialDevice) Listen(btnch chan string, quitch chan struct{}) {
free:
	// Keep looping since sd.Listen() will return if no data is sent
	for {
		select {
		case <-quitch:
			break free
		default:
			// If we get data, send to chan
			actionID, err := s.ScanThing()
			if err != nil {
				slog.Debug(fmt.Sprint("Listen err: ", err))
			}
			btnch <- actionID
		}
	}
}

// Listen & return data
// Runs in a bufio.Scanner.Scan() loop
func (s *SerialDevice) ScanThing() (actionID string, err error) {
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
