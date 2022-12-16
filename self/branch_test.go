package self

import (
	"testing"

	selfSym "github.com/taubyte/go-sdk-symbols/self"
	"github.com/taubyte/go-sdk/errno"
)

func TestBranchError(t *testing.T) {
	selfSym.BranchSize = func(sizePtr *uint32) errno.Error {
		return 1
	}

	_, err := Branch()
	if err == nil {
		t.Error("Expected error")
		return
	}

	selfSym.BranchSize = func(sizePtr *uint32) errno.Error {
		*sizePtr = 5
		return 0
	}

	selfSym.Branch = func(ptr *byte) errno.Error {
		return 1
	}

	_, err = Branch()
	if err == nil {
		t.Error("Expected error")
		return
	}
}
