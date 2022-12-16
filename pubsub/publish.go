package pubsub

import (
	"fmt"

	pubsubSym "github.com/taubyte/go-sdk-symbols/pubsub"
)

// Publish will publish provided data onto the pub-sub channel
// returns an error
func (c *ChannelObject) Publish(data []byte) error {
	var dataPtr *byte
	if len(data) != 0 {
		dataPtr = &data[0]
	}

	err := pubsubSym.PublishToChannel(c.name, dataPtr, uint32(len(data)))
	if err != 0 {
		return fmt.Errorf("Publish to channel `%s` failed with: %s", c.name, err)
	}

	return nil
}
