package event

import (
	"testing"

	"github.com/ipfs/go-cid"
	symbols "github.com/taubyte/go-sdk-symbols/p2p/event"
	"github.com/taubyte/go-sdk/errno"
)

func TestFrom(t *testing.T) {
	var e Event

	testFrom, err := cid.Decode("QmZjBpQzRw8UovKaMoWJ3qvQrHZvErqQuXNyXNusm4XYK3")
	if err != nil {
		return
	}

	symbols.MockData{ClientId: 1, From: testFrom}.Mock()

	_, err = e.From()
	if err == nil {
		t.Error("Expected error")
		return
	}

	symbols.MockData{From: testFrom}.Mock()

	from, err := e.From()
	if err != nil {
		t.Error(err)
		return
	}
	if string(from.Bytes()) != string(testFrom.Bytes()) {
		t.Errorf("Expected %s, got %s", string(testFrom.Bytes()), string(from.Bytes()))
		return
	}

	// Basic get error
	symbols.MockData{From: testFrom}.Mock()

	symbols.GetP2PEventFrom = func(eventId uint32, bufPtr *byte) (error errno.Error) {
		return 1
	}

	_, err = e.From()
	if err == nil {
		t.Error("Expected error")
		return
	}
}
