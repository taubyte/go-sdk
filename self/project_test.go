package self

import (
	"testing"

	selfSym "github.com/taubyte/go-sdk-symbols/self"
	"github.com/taubyte/go-sdk/errno"
)

func TestProjectError(t *testing.T) {
	selfSym.ProjectSize = func(sizePtr *uint32) errno.Error {
		return 1
	}

	_, err := Project()
	if err == nil {
		t.Error("Expected error")
		return
	}

	selfSym.ProjectSize = func(sizePtr *uint32) errno.Error {
		*sizePtr = 5
		return 0
	}

	selfSym.Project = func(ptr *byte) errno.Error {
		return 1
	}

	_, err = Project()
	if err == nil {
		t.Error("Expected error")
		return
	}
}
