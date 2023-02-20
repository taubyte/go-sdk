package event

import (
	"fmt"

	p2phttpEventSym "github.com/taubyte/go-sdk-symbols/p2p/event"
)

// Command returns the command used to call the given p2p function and an error.
func (e Event) Command() (string, error) {
	var size uint32
	err := p2phttpEventSym.GetP2PEventCommandSize(uint32(e), &size)
	if err != 0 {
		return "", fmt.Errorf("Getting command data size failed with: %s", err)
	}
	if size == 0 {
		return "", nil
	}

	data := make([]byte, size)
	err = p2phttpEventSym.GetP2PEventCommand(uint32(e), &data[0])
	if err != 0 {
		return "", fmt.Errorf("Getting command data failed with: %s", err)
	}

	return string(data), nil
}
