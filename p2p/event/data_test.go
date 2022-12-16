package event

import (
	"testing"

	symbols "github.com/taubyte/go-sdk-symbols/p2p/event"
	"github.com/taubyte/go-sdk/errno"
)

func TestData(t *testing.T) {
	var e Event
	testData := []byte("Hello, world!")

	symbols.MockData{ClientId: 1, Data: testData}.Mock()

	_, err := e.Data()
	if err == nil {
		t.Error("Expected error")
		return
	}

	symbols.MockData{Data: testData}.Mock()

	data, err := e.Data()
	if err != nil {
		t.Error(err)
		return
	}
	if string(data) != string(testData) {
		t.Errorf("Expected %s, got %s", string(testData), string(data))
		return
	}

	// Size 0
	symbols.MockData{Data: nil}.Mock()

	_, err = e.Data()
	if err != nil {
		t.Error(err)
		return
	}

	// Basic get error
	symbols.MockData{Data: testData}.Mock()

	symbols.GetP2PEventData = func(eventId uint32, bufPtr *byte) (error errno.Error) {
		return 1
	}

	_, err = e.Data()
	if err == nil {
		t.Error("Expected error")
		return
	}
}
