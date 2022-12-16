package event

import (
	"io"
	"testing"

	symbols "github.com/taubyte/go-sdk-symbols/event"
)

func TestReadBasic(t *testing.T) {
	testData := "Hello, world!"
	symbols.MockData{
		Body: []byte(testData),
	}.Mock()

	var e HttpEvent
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
	symbols.MockData{
		Body:    []byte("Hello, world"),
		EventId: 1,
	}.Mock()

	var e HttpEvent
	_, err := io.ReadAll(e.Body())
	if err == nil {
		t.Error("Expected error")
		return
	}
}

func TestCloseError(t *testing.T) {
	symbols.MockData{
		Body:    []byte("Hello, world"),
		EventId: 1,
	}.Mock()

	var e HttpEvent
	err := e.Body().Close()
	if err == nil {
		t.Error("Expected error")
		return
	}
}
