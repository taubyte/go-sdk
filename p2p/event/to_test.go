package event

import (
	"testing"

	"github.com/ipfs/go-cid"
	symbols "github.com/taubyte/go-sdk-symbols/p2p/event"
	"github.com/taubyte/go-sdk/errno"
)

func TestTo(t *testing.T) {
	var e Event

	testTo, err := cid.Parse("QmZjBpQzRw8UovKaMoWJ3qvQrHZvErqQuXNyXNusm4XYK3")
	if err != nil {
		return
	}

	symbols.MockData{ClientId: 1, To: testTo}.Mock()

	_, err = e.To()
	if err == nil {
		t.Error("Expected error")
		return
	}

	symbols.MockData{To: testTo}.Mock()

	to, err := e.To()
	if err != nil {
		t.Error(err)
		return
	}
	if string(to.Bytes()) != string(testTo.Bytes()) {
		t.Errorf("Expected %s, got %s", string(testTo.Bytes()), string(to.Bytes()))
		return
	}

	// Size 0
	symbols.MockData{To: cid.Cid{}}.Mock()

	_, err = e.To()
	if err != nil {
		t.Error(err)
		return
	}

	// Basic get error
	symbols.MockData{To: testTo}.Mock()

	symbols.GetP2PEventTo = func(eventId uint32, bufPtr *byte) (error errno.Error) {
		return 1
	}

	_, err = e.To()
	if err == nil {
		t.Error("Expected error")
		return
	}
}
