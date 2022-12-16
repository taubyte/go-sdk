package event

import (
	"testing"

	symbols "github.com/taubyte/go-sdk-symbols/p2p/event"
	"github.com/taubyte/go-sdk/errno"
)

func TestCommand(t *testing.T) {
	var e Event
	testCommand := "someCommand"

	symbols.MockData{ClientId: 1, Command: testCommand}.Mock()

	_, err := e.Command()
	if err == nil {
		t.Error("Expected error")
		return
	}

	symbols.MockData{Command: testCommand}.Mock()

	data, err := e.Command()
	if err != nil {
		t.Error(err)
		return
	}
	if data != testCommand {
		t.Errorf("Expected %s, got %s", testCommand, data)
		return
	}

	// Size 0
	symbols.MockData{Command: ""}.Mock()

	_, err = e.Command()
	if err != nil {
		t.Error(err)
		return
	}

	// Basic get error
	symbols.MockData{Command: testCommand}.Mock()
	symbols.GetP2PEventCommand = func(eventId uint32, bufPtr *byte) (error errno.Error) {
		return 1
	}

	_, err = e.Command()
	if err == nil {
		t.Error("Expected error")
		return
	}
}
