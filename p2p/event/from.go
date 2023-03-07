package event

import (
	"fmt"

	"github.com/ipfs/go-cid"
	p2pEventSym "github.com/taubyte/go-sdk-symbols/p2p/event"
	"github.com/taubyte/go-sdk/utils/codec"
)

// From returns a cid.Cid of the sending node and an error.
func (e Event) From() (cid.Cid, error) {
	reader := codec.CidReader()
	err := p2pEventSym.GetP2PEventFrom(uint32(e), reader.Ptr())
	if err != 0 {
		return cid.Cid{}, fmt.Errorf("Getting from address data failed with: %s", err)
	}

	return reader.Parse()
}
