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
