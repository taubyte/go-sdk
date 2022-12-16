package internal_test

import (
	"errors"
	"testing"

	globalSym "github.com/taubyte/go-sdk-symbols/globals"
	"github.com/taubyte/go-sdk/errno"
	"github.com/taubyte/go-sdk/globals/internal"
	"github.com/taubyte/go-sdk/globals/scope"
	"github.com/taubyte/go-sdk/globals/str"
)

func TestBaseBasic(t *testing.T) {
	// Mocking the calls to the vm for usage in tests and playground
	globalSym.MockData{
		Data: map[string][]uint8{
			"/string/name":             []byte("Hello, world!"),
			"/application/string/name": []byte("Hello, world! (scoped)"),
			"/string/zero":             nil,
		},
	}.Mock()

	name, err := str.GetOrCreate("name")
	if err != nil {
		return
	}
	if name.Value() != "Hello, world!" {
		t.Errorf("Expected 'Hello, world!', got '%s'", name.Value())
		return
	}

	err = name.Set("Hello, Mars!")
	if err != nil {
		return
	}

	name, err = str.Get("name")
	if err != nil {
		return
	}
	if name.Value() != "Hello, Mars!" {
		t.Errorf("Expected 'Hello, Mars!', got '%s'", name.Value())
		return
	}

	name, err = str.Get("zero")
	if err != nil {
		return
	}
	if name.Value() != "" {
		t.Errorf("Expected '', got '%s'", name.Value())
		return
	}

	name, err = str.Get("name", scope.Application)
	if err != nil {
		return
	}
	if name.Value() != "Hello, world! (scoped)" {
		t.Errorf("Expected 'Hello, world! (scoped)', got '%s'", name.Value())
		return
	}

	return
}

func TestBaseErrors(t *testing.T) {
	globalSym.MockData{
		Data: map[string][]uint8{
			"/string/name":             []byte("Hello, world!"),
			"/application/string/name": []byte("Hello, world! (scoped)"),
			"/string/zero":             nil,
		},
	}.Mock()

	base := &internal.Base[string]{
		Name: "name",
		Key:  "string",

		ToBase: func(b []byte) (string, error) {
			return "", errors.New("Failure")
		},

		ToBytes: func(s string) ([]byte, error) {
			return nil, errors.New("Failure")
		},

		Scope: nil,
	}

	err := base.Get()
	if err == nil {
		t.Errorf("Expected error, got nil")
		return
	}

	err = base.GetOrCreate()
	if err == nil {
		t.Errorf("Expected error, got nil")
		return
	}

	err = base.Put()
	if err == nil {
		t.Errorf("Expected error, got nil")
		return
	}

	base = &internal.Base[string]{
		Name: "name",
		Key:  "string",

		ToBase: func(b []byte) (string, error) {
			return "", nil
		},

		ToBytes: func(s string) ([]byte, error) {
			return nil, nil
		},

		Scope: nil,
	}

	globalSym.PutGlobalValue = func(name string, application, function uint32, value []byte) errno.Error {
		return 1
	}

	err = base.Put()
	if err == nil {
		t.Errorf("Expected error, got nil")
		return
	}
}
