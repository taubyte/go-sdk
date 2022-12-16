package self

import (
	"fmt"

	selfSym "github.com/taubyte/go-sdk-symbols/self"
)

func Project() (string, error) {
	var size uint32
	err0 := selfSym.ProjectSize(&size)
	if err0 != 0 {
		return "", fmt.Errorf("Getting project size failed with: %s", err0)
	}
	if size == 0 {
		return "", nil
	}

	project := make([]byte, size)
	err0 = selfSym.Project(&project[0])
	if err0 != 0 {
		return "", fmt.Errorf("Getting project failed with: %s", err0)
	}

	return string(project), nil
}

func Application() (string, error) {
	var size uint32
	err0 := selfSym.ApplicationSize(&size)
	if err0 != 0 {
		return "", fmt.Errorf("Getting application size failed with: %s", err0)
	}
	if size == 0 {
		return "", nil
	}

	application := make([]byte, size)
	err0 = selfSym.Application(&application[0])
	if err0 != 0 {
		return "", fmt.Errorf("Getting application failed with: %s", err0)
	}

	return string(application), nil
}

func Function() (string, error) {
	var size uint32
	err0 := selfSym.IdSize(&size)
	if err0 != 0 {
		return "", fmt.Errorf("Getting function size failed with: %s", err0)
	}
	if size == 0 {
		return "", nil
	}

	function := make([]byte, size)
	err0 = selfSym.Id(&function[0])
	if err0 != 0 {
		return "", fmt.Errorf("Getting function failed with: %s", err0)
	}

	return string(function), nil
}

func Branch() (string, error) {
	var size uint32
	err0 := selfSym.BranchSize(&size)
	if err0 != 0 {
		return "", fmt.Errorf("Getting branch size failed with: %s", err0)
	}
	if size == 0 {
		return "", nil
	}

	branch := make([]byte, size)
	err0 = selfSym.Branch(&branch[0])
	if err0 != 0 {
		return "", fmt.Errorf("Getting branch failed with: %s", err0)
	}

	return string(branch), nil
}

func Commit() (string, error) {
	var size uint32
	err0 := selfSym.CommitSize(&size)
	if err0 != 0 {
		return "", fmt.Errorf("Getting commit size failed with: %s", err0)
	}
	if size == 0 {
		return "", nil
	}

	commit := make([]byte, size)
	err0 = selfSym.Commit(&commit[0])
	if err0 != 0 {
		return "", fmt.Errorf("Getting commit failed with: %s", err0)
	}

	return string(commit), nil
}
