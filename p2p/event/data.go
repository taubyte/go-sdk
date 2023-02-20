package event

import (
	"fmt"

	p2phttpEventSym "github.com/taubyte/go-sdk-symbols/p2p/event"
)

// Data returns the data sent to the function as bytes and an error.
func (e Event) Data() ([]byte, error) {
	var size uint32
	err := p2phttpEventSym.GetP2PEventDataSize(uint32(e), &size)
	if err != 0 {
		return nil, fmt.Errorf("Getting command data size failed with: %s", err)
	}
	if size == 0 {
		return nil, nil
	}

	data := make([]byte, size)
	err = p2phttpEventSym.GetP2PEventData(uint32(e), &data[0])
	if err != 0 {
		return nil, fmt.Errorf("Getting command data failed with: %s", err)
	}

	return data, nil
}
