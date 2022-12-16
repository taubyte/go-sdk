package self

import (
	"testing"

	selfSym "github.com/taubyte/go-sdk-symbols/self"
	"github.com/taubyte/go-sdk/errno"
)

func TestApplicationError(t *testing.T) {
	selfSym.ApplicationSize = func(SizePtr *uint32) errno.Error {
		return 1
	}

	_, err := Application()
	if err == nil {
		t.Error("Expected error")
		return
	}

	selfSym.ApplicationSize = func(SizePtr *uint32) errno.Error {
		*SizePtr = 5
		return 0
	}

	selfSym.Application = func(Ptr *byte) errno.Error {
		return 1
	}

	_, err = Application()
	if err == nil {
		t.Error("Expected error")
		return
	}
}
