package event

import (
	"fmt"

	p2pEventSym "github.com/taubyte/go-sdk-symbols/p2p/event"
)

// Protocol returns the protocol which the event was called and an error.
func (e Event) Protocol() (string, error) {
	var size uint32
	err := p2pEventSym.GetP2PEventProtocolSize(uint32(e), &size)
	if err != 0 {
		return "", fmt.Errorf("Getting protocol data size failed with: %s", err)
	}
	if size == 0 {
		return "", nil
	}

	data := make([]byte, size)
	err = p2pEventSym.GetP2PEventProtocol(uint32(e), &data[0])
	if err != 0 {
		return "", fmt.Errorf("Getting protocol data failed with: %s", err)
	}

	return string(data), nil
}
