package node

import (
	"testing"
	"unsafe"

	nodeSym "github.com/taubyte/go-sdk-symbols/p2p/node"
	"github.com/taubyte/go-sdk/errno"
	"github.com/taubyte/go-sdk/utils/codec"
	"gotest.tools/v3/assert"
)

func TestDiscover(t *testing.T) {
	// No max
	{
		peers, err := Discover(0, 0)
		assert.NilError(t, err)
		assert.Equal(t, len(peers), 0)
	}

	// 0 size
	{
		nodeSym.DiscoverPeersSize = func(max, nsTimeout uint32, id, peersSize *uint32) (error errno.Error) {
			return 0
		}

		_, err := Discover(1, 0)
		assert.NilError(t, err)
	}

	// DiscoverPeersSize error
	{
		nodeSym.DiscoverPeersSize = func(max, nsTimeout uint32, id, peersSize *uint32) (error errno.Error) {
			return 1
		}

		_, err := Discover(1, 0)
		assert.Assert(t, err != nil)
	}

	// DiscoverPeers error
	{
		nodeSym.DiscoverPeersSize = func(max, nsTimeout uint32, id, peersSize *uint32) (error errno.Error) {
			*peersSize = 1
			return 0
		}

		nodeSym.DiscoverPeers = func(id uint32, peersBuf *byte) (error errno.Error) {
			return 1
		}

		_, err := Discover(1, 0)
		assert.Assert(t, err != nil)
	}

	// Conversion error
	{
		nodeSym.DiscoverPeersSize = func(max, nsTimeout uint32, id, peersSize *uint32) (error errno.Error) {
			*peersSize = 12
			return 0
		}

		nodeSym.DiscoverPeers = func(id uint32, peersBuf *byte) (error errno.Error) {
			data := unsafe.Slice(peersBuf, 12)
			copy(data, "Hello, world")
			return 0
		}

		_, err := Discover(1, 0)
		assert.Assert(t, err != nil)
	}

	// cid.Cast error
	{
		var invalid []byte
		err := codec.Convert([][]byte{[]byte("invalid-cid")}).To(&invalid)
		assert.NilError(t, err)

		nodeSym.DiscoverPeersSize = func(max, nsTimeout uint32, id, peersSize *uint32) (error errno.Error) {
			*peersSize = uint32(len(invalid))
			return 0
		}

		nodeSym.DiscoverPeers = func(id uint32, peersBuf *byte) (error errno.Error) {
			data := unsafe.Slice(peersBuf, uint32(len(invalid)))
			copy(data, invalid)
			return 0
		}

		_, err = Discover(1, 0)
		assert.Assert(t, err != nil)
	}
}
