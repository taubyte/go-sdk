package event

import (
	"testing"

	symbols "github.com/taubyte/go-sdk-symbols/p2p/event"
)

func TestWrite(t *testing.T) {
	var e Event
	testData := []byte("Hello, world")

	symbols.MockData{
		ClientId: 1,
	}.Mock()

	err := e.Write(nil)
	if err == nil {
		t.Error("Expected error")
		return
	}

	m := symbols.MockData{}.Mock()

	err = e.Write(testData)
	if err != nil {
		t.Error(err)
		return
	}
	if string(m.DataToSend) != string(testData) {
		t.Errorf("Expected %s, got %s", string(testData), string(m.DataToSend))
		return
	}
}
