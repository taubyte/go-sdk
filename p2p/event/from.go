package event

import (
	"fmt"

	"github.com/ipfs/go-cid"
	p2phttpEventSym "github.com/taubyte/go-sdk-symbols/p2p/event"
)

// From returns a cid.Cid of the sending node and an error.
func (e Event) From() (cid.Cid, error) {
	var size uint32
	err := p2phttpEventSym.GetP2PEventFromSize(uint32(e), &size)
	if err != 0 {
		return cid.Cid{}, fmt.Errorf("Getting from address data size failed with: %s", err)
	}
	if size == 0 {
		return cid.Cid{}, nil
	}

	cidBytes := make([]byte, size)
	err = p2phttpEventSym.GetP2PEventFrom(uint32(e), &cidBytes[0])
	if err != 0 {
		return cid.Cid{}, fmt.Errorf("Getting from address data failed with: %s", err)
	}

	_, cidFromBytes, err0 := cid.CidFromBytes(cidBytes)
	return cidFromBytes, err0
}
