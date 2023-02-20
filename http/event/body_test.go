package event

import (
	"io"
	"testing"

	httpEventSym "github.com/taubyte/go-sdk-symbols/http/event"
)

func TestReadBasic(t *testing.T) {
	testData := "Hello, world!"
	httpEventSym.MockData{
		Body: []byte(testData),
	}.Mock()

	var e Event
	body, err := io.ReadAll(e.Body())
	if err != nil {
		t.Error(err)
		return
	}
	if string(body) != testData {
		t.Errorf("Got body `%s` expected `%s`", string(body), testData)
		return
	}
}

func TestReadError(t *testing.T) {
	httpEventSym.MockData{
		Body:    []byte("Hello, world"),
		EventId: 1,
	}.Mock()

	var e Event
	_, err := io.ReadAll(e.Body())
	if err == nil {
		t.Error("Expected error")
		return
	}
}

func TestCloseError(t *testing.T) {
	httpEventSym.MockData{
		Body:    []byte("Hello, world"),
		EventId: 1,
	}.Mock()

	var e Event
	err := e.Body().Close()
	if err == nil {
		t.Error("Expected error")
		return
	}
}
