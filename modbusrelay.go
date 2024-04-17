// modbusrelay package implements control for Modbus relay modules.
package modbusrelay

import (
	"github.com/goburrow/modbus"
)

// RelayController represents a Modbus relay controller.
type RelayController struct {
	handler *modbus.RTUClientHandler
	client  modbus.Client
}

// NewRelayController creates a new RelayController instance.
func NewRelayController(device string, baudRate int, slaveID byte) (*RelayController, error) {
	handler := modbus.NewRTUClientHandler(device)
	handler.BaudRate = baudRate
	handler.DataBits = 8
	handler.Parity = "N"
	handler.StopBits = 1
	handler.SlaveId = slaveID

	err := handler.Connect()
	if err != nil {
		return nil, err
	}

	client := modbus.NewClient(handler)

	return &RelayController{
		handler: handler,
		client:  client,
	}, nil
}

// SetRelayOn turns on the specified relay.
func (rc *RelayController) SetRelayOn(relayNum uint16) error {
	_, err := rc.client.WriteSingleCoil(relayNum, 0xFF00)
	return err
}

// SetRelayOff turns off the specified relay.
func (rc *RelayController) SetRelayOff(relayNum uint16) error {
	_, err := rc.client.WriteSingleCoil(relayNum, 0x0000)
	return err
}

// Close closes the Modbus connection.
func (rc *RelayController) Close() error {
	return rc.handler.Close()
}
