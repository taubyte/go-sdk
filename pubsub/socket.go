package pubsub

import (
	"errors"
	"fmt"
	"net/url"

	pubsubSym "github.com/taubyte/go-sdk-symbols/pubsub"
)

func (c *ChannelObject) WebSocket() WebSocket {
	return WebSocket{c.name}
}

// WebSocket Url gets a generated url based on the configured channel's ID and the project ID
// returns a net/url.URL and an error
func (ws WebSocket) Url() (wsUrl url.URL, err error) {
	if len(ws.name) == 0 {
		err = errors.New("Must have a valid channel name to get a url; got ``")
		return
	}

	var size uint32
	err0 := pubsubSym.GetWebSocketURLSize(ws.name, &size)
	if err0 != 0 {
		err = fmt.Errorf("Getting web socket URL for channel `%s` failed with: %s", ws.name, err0)
		return
	}
	if size == 0 {
		return
	}

	socketURL := make([]byte, size)
	err0 = pubsubSym.GetWebSocketURL(ws.name, &socketURL[0])
	if err0 != 0 {
		err = fmt.Errorf("Getting web socket URL for channel `%s` failed with: %s", ws.name, err0)
		return
	}

	wsUrl = url.URL{
		Path: string(socketURL),
	}

	return
}
