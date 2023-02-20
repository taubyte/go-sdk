package event

import (
	"fmt"

	p2phttpEventSym "github.com/taubyte/go-sdk-symbols/p2p/event"
)

// Write writes given data and returns an error.
func (e Event) Write(data []byte) error {
	var dataPtr *byte
	if len(data) != 0 {
		dataPtr = &data[0]
	}

	err := p2phttpEventSym.WriteP2PResponse(uint32(e), dataPtr, uint32(len(data)))
	if err != 0 {
		return fmt.Errorf("Writing p2p response failed with %s", err)
	}

	return nil
}
