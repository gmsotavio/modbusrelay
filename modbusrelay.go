// modbusrelay package implements control for Modbus relay modules.
package modbusrelay

import (
	"fmt"
	"github.com/simonvetter/modbus"
	"time"
)

// RelayControllerInterface defines methods for controlling relays.
type RelayControllerInterface interface {
	SetRelayOn(relayNum uint16) error
	SetRelayOff(relayNum uint16) error
	Close() error
}

// RelayController represents a Modbus relay controller.
type RelayController struct {
	client *modbus.ModbusClient
}

// NewRelayController creates a new RelayController instance.
func NewRelayController(device string, baudRate uint, slaveID byte) (RelayControllerInterface, error) {

	client, err := modbus.NewClient(&modbus.ClientConfiguration{
		URL:      "rtu://" + device,
		Speed:    baudRate,
		DataBits: 8,
		Parity:   modbus.PARITY_NONE,
		StopBits: 2,
		Timeout:  300 * time.Millisecond,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create Modbus client: %w", err)
	}

	err = client.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open Modbus connection: %w", err)
	}

	err = client.SetUnitId(slaveID)
	if err != nil {
		// Close the client if setting unit ID fails
		client.Close()
		return nil, fmt.Errorf("failed to set unit ID: %w", err)
	}

	return &RelayController{
		client: client,
	}, nil

}

// SetRelayOn turns on the specified relay.
func (rc *RelayController) SetRelayOn(relayNum uint16) error {
	err := rc.client.WriteCoil(relayNum, true)
	return err
}

// SetRelayOff turns off the specified relay.
func (rc *RelayController) SetRelayOff(relayNum uint16) error {
	err := rc.client.WriteCoil(relayNum, false)
	return err
}

// Close closes the Modbus connection.
func (rc *RelayController) Close() error {
	return rc.client.Close()
}
