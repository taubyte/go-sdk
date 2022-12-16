package self

import (
	"testing"

	selfSym "github.com/taubyte/go-sdk-symbols/self"
	"github.com/taubyte/go-sdk/errno"
)

func TestFunctionError(t *testing.T) {
	selfSym.IdSize = func(sizePtr *uint32) errno.Error {
		return 1
	}

	_, err := Function()
	if err == nil {
		t.Error("Expected error")
		return
	}

	selfSym.IdSize = func(sizePtr *uint32) errno.Error {
		*sizePtr = 5
		return 0
	}

	selfSym.Id = func(ptr *byte) errno.Error {
		return 1
	}

	_, err = Function()
	if err == nil {
		t.Error("Expected error")
		return
	}
}
