package node

import (
	"fmt"
	"time"

	"github.com/ipfs/go-cid"
	nodeSym "github.com/taubyte/go-sdk-symbols/p2p/node"
	"github.com/taubyte/go-sdk/utils/codec"
)

func Discover(max uint32, timeout time.Duration) ([]cid.Cid, error) {
	if max == 0 {
		return []cid.Cid{}, nil
	}

	var (
		id        uint32
		peersSize uint32
	)

	err0 := nodeSym.DiscoverPeersSize(max, uint32(timeout), &id, &peersSize)
	if err0 != 0 {
		return nil, fmt.Errorf("discover failed with: %s", err0)
	}
	if peersSize == 0 {
		return []cid.Cid{}, nil
	}

	peerBytes := make([]byte, peersSize)
	err0 = nodeSym.DiscoverPeers(id, &peerBytes[0])
	if err0 != 0 {
		return nil, fmt.Errorf("reading peers after discover failed with: %s", err0)
	}

	var peerBytesSlice [][]byte
	err := codec.Convert(peerBytes).To(&peerBytesSlice)
	if err != nil || len(peerBytesSlice) == 0 {
		return nil, fmt.Errorf("converting bytes failed with: %s", err)
	}

	peers := make([]cid.Cid, len(peerBytesSlice))
	for idx, peer := range peerBytesSlice {
		peers[idx], err = cid.Cast(peer)
		if err != nil {
			return nil, err
		}
	}

	return peers, nil
}
