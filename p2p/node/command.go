package node

import (
	"fmt"

	nodeSym "github.com/taubyte/go-sdk-symbols/p2p/node"
)

// Send returns the command response and an error,
// used to send data over p2p on given protocol and command.
func (c Command) Send(body []byte) ([]byte, error) {
	var dataPtr *byte
	if len(body) != 0 {
		dataPtr = &body[0]
	}

	var responseSize uint32
	err := nodeSym.SendCommand(uint32(c), dataPtr, uint32(len(body)), &responseSize)
	if err != 0 {
		return nil, fmt.Errorf("Send command failed with %s", err)
	}
	if responseSize == 0 {
		return nil, nil
	}

	response := make([]byte, responseSize)
	err = nodeSym.ReadCommandResponse(uint32(c), &response[0], responseSize)
	if err != 0 {
		return nil, fmt.Errorf("Reading command response failed with %s", err)
	}

	return response, nil
}
