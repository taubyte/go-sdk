package event

import (
	"testing"

	symbols "github.com/taubyte/go-sdk-symbols/p2p/event"
	"github.com/taubyte/go-sdk/errno"
)

func TestProtocol(t *testing.T) {
	var e Event
	testProtocol := "/test/v1"

	symbols.MockData{ClientId: 1, Protocol: testProtocol}.Mock()

	_, err := e.Protocol()
	if err == nil {
		t.Error("Expected error")
		return
	}

	symbols.MockData{Protocol: testProtocol}.Mock()

	protocol, err := e.Protocol()
	if err != nil {
		t.Error(err)
		return
	}
	if protocol != testProtocol {
		t.Errorf("Expected %s, got %s", testProtocol, protocol)
		return
	}

	// Size 0
	symbols.MockData{Protocol: ""}.Mock()

	_, err = e.Protocol()
	if err != nil {
		t.Error(err)
		return
	}

	// Basic get error
	symbols.MockData{Protocol: testProtocol}.Mock()

	symbols.GetP2PEventProtocol = func(eventId uint32, bufPtr *byte) (error errno.Error) {
		return 1
	}

	_, err = e.Protocol()
	if err == nil {
		t.Error("Expected error")
		return
	}
}
