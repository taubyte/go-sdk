package node

// Event sent to a p2p function call
type P2PEvent uint32

// A p2p service with on a given protocol
type Service struct {
	protocol string
}

// A p2p command
type Command uint32
