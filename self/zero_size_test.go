package self

import (
	"testing"

	selfSym "github.com/taubyte/go-sdk-symbols/self"
)

func TestZeroSize(t *testing.T) {
	selfSym.MockProject("")
	selfSym.MockApplication("")
	selfSym.MockBranch("")
	selfSym.MockCommit("")
	selfSym.MockFunction("")

	project, err := Project()
	if err != nil {
		t.Error(err)
		return
	}
	if project != "" {
		t.Error("Expected no value for project")
		return
	}

	application, err := Application()
	if err != nil {
		t.Error(err)
		return
	}
	if application != "" {
		t.Error("Expected no value for application")
		return
	}

	branch, err := Branch()
	if err != nil {
		t.Error(err)
		return
	}
	if branch != "" {
		t.Error("Expected no value for branch")
		return
	}

	commit, err := Commit()
	if err != nil {
		t.Error(err)
		return
	}
	if commit != "" {
		t.Error("Expected no value for commit")
		return
	}

	function, err := Function()
	if err != nil {
		t.Error(err)
		return
	}
	if function != "" {
		t.Error("Expected no value for function")
		return
	}
}
