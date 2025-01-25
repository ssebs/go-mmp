package serialdevice

import (
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/ssebs/go-mmp/models"
	"go.bug.st/serial"
)

// SerialDevice is used to manage an arduino's serial connections
type SerialDevice struct {
	Conn     serial.Port
	PortName string
	Mode     *serial.Mode
	Timeout  time.Duration // defaults to 20ms
}

// This will Open a connection to the portName with baudRate
func NewSerialDevice(portName string, baudRate int) (*SerialDevice, error) {
	serialDevice := &SerialDevice{
		PortName: portName,
		Timeout:  time.Millisecond * 20,
		Mode:     &serial.Mode{BaudRate: baudRate},
		Conn:     nil,
	}

	if err := serialDevice.openConnection(); err != nil {
		return nil, err
	}

	return serialDevice, nil
}
func NewSerialDeviceFromConfig(c *models.Config) (*SerialDevice, error) {
	arduino, err := NewSerialDevice(c.Metadata.SerialPortName, c.Metadata.SerialBaudRate)
	return arduino, err
}

func (s *SerialDevice) ChangePortAndReconnect(portName string, baudRate int) error {
	if err := s.Conn.Close(); err != nil {
		return fmt.Errorf("failed to close connection, err:%s", err)
	}

	s.PortName = portName
	s.Mode.BaudRate = baudRate
	return s.openConnection()
}

func (s *SerialDevice) CloseConnection() error {
	return s.Conn.Close()
}

// Listen for data from a *SerialDevice, to be used in a goroutine
// Takes in a btnch to send data to when the serial connection gets something,
// and a quitch if we need to stop the goroutine
func (s *SerialDevice) Listen(btnch chan string, quitch chan struct{}) {
	buf := make([]byte, 8)
free:
	// Keep looping since sd.Listen() will return if no data is sent
	for {
		select {
		case <-quitch:
			break free
		default:
			// If we get data, send to chan
			actionID := &strings.Builder{}
			for {
				n, err := s.Conn.Read(buf)
				if err != nil {
					slog.Debug(fmt.Sprint("Listen err: ", err))
				}
				if n == 0 {
					break
				}

				data := buf[0:n]
				actionID.WriteString(strings.TrimSpace(string(data)))

			}
			btnch <- actionID.String()
		}
	}
}

func (s *SerialDevice) openConnection() error {
	conn, err := serial.Open(s.PortName, s.Mode)
	if err != nil {
		return fmt.Errorf("failed to open serial %s, err: %s", s.PortName, err)
	}
	s.Conn = conn

	err = s.Conn.SetReadTimeout(s.Timeout)
	if err != nil {
		return fmt.Errorf("failed to set read timeout, err: %s", err)
	}

	return nil
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
