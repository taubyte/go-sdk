package pubsub

import "errors"

// Channel will create a new ChannelObject based on the provided name
// returns a ChannelObject and an error
func Channel(name string) (*ChannelObject, error) {
	if name == "" {
		return nil, errors.New("Channel name cannot be empty")
	}

	// TODO, We should call a lookup here, and return an error if
	// no valid channels are found

	return &ChannelObject{name: name}, nil
}

// returns the name of the channel
func (c *ChannelObject) Name() string {
	return c.name
}
