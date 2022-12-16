package node

import (
	"errors"
	"fmt"

	nodeSym "github.com/taubyte/go-sdk-symbols/p2p/node"
)

// New returns a p2p Service bound to a given protocol
func New(protocol string) *Service {
	return &Service{protocol: protocol}
}

// Returns the protocol of a given service
func (s *Service) Protocol() string {
	return s.protocol
}

// Command returns a p2p Command and an error
// Used for sending a given command on a protocol
func (s *Service) Command(command string) (Command, error) {
	var id uint32
	err := nodeSym.NewCommand(s.protocol, command, &id)
	if err != 0 {
		return 0, fmt.Errorf("New P2P command failed with: %s", err)
	}

	return Command(id), nil
}

// Listen returns the projectProtocol and an error.
// Asks a node to listen on a protocol, and returns the protocol
// the node is set to listen on.
func (s *Service) Listen() (protocol string, err error) {
	var protocolSize uint32
	err0 := nodeSym.ListenToProtocolSize(s.protocol, &protocolSize)
	if protocolSize == 0 {
		return "", errors.New("Project protocol not found")
	}

	projectProtocol := make([]byte, protocolSize)
	err1 := nodeSym.ListenToProtocol(s.protocol, &projectProtocol[0], protocolSize)
	protocol = string(projectProtocol)

	// Errors are checked after so that the protocol is sent to the user whether
	// or not the actual listen request fails.
	if err0 != 0 {
		return protocol, fmt.Errorf("Listen to protocol `%s` getting size failed with: %s", s.protocol, err0)
	}
	if err1 != 0 {
		return protocol, fmt.Errorf("Listen to protocol `%s` failed with: %s", s.protocol, err1)
	}

	return protocol, nil
}
