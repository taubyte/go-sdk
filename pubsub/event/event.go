package event

import (
	"errors"
	"fmt"

	pubsubSym "github.com/taubyte/go-sdk-symbols/pubsub"
	pubsubNode "github.com/taubyte/go-sdk/pubsub/node"
)

// Data will get the data received with the pub-sub event
// returns a byte slice and an error
func (e Event) Data() ([]byte, error) {
	var size uint32
	err := pubsubSym.GetMessageDataSize(uint32(e), &size)
	if err != 0 {
		return nil, fmt.Errorf("Getting data size failed with: %s", err)
	}
	if size == 0 {
		return nil, nil
	}

	data := make([]byte, size)
	err = pubsubSym.GetMessageData(uint32(e), &data[0])
	if err != 0 {
		return nil, fmt.Errorf("Getting message data failed with: %s", err)
	}

	return data, nil
}

// Channel will get the name of the event's channel
// returns a ChannelObject and an error
func (e Event) Channel() (*pubsubNode.ChannelObject, error) {
	var size uint32
	err := pubsubSym.GetMessageChannelSize(uint32(e), &size)
	if err != 0 {
		return nil, fmt.Errorf("Getting message channel size failed with: %s", err)
	}
	if size == 0 {
		return nil, errors.New("Got an empty channel")
	}

	channelName := make([]byte, size)
	err = pubsubSym.GetMessageChannel(uint32(e), &channelName[0])
	if err != 0 {
		return nil, fmt.Errorf("Getting message channel failed with: %s", err)
	}

	return pubsubNode.Channel(string(channelName))
}
