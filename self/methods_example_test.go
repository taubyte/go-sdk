package self_test

import (
	"fmt"

	selfSym "github.com/taubyte/go-sdk-symbols/self"
	"github.com/taubyte/go-sdk/self"
)

func ExampleProject() {
	// Mocking the calls to the vm for usage in tests and playground
	selfSym.MockProject("QmZY4u91d1YALDN2LTbpVtgwW8iT5cK9PE1bHZqX9J51Tv")

	// Called from a function in the project "QmZY4u91d1YALDN2LTbpVtgwW8iT5cK9PE1bHZqX9J51Tv"
	project, err := self.Project()
	if err != nil {
		panic(err)
	}

	fmt.Println(project)
	// Output: QmZY4u91d1YALDN2LTbpVtgwW8iT5cK9PE1bHZqX9J51Tv
}

func ExampleApplication() {
	// Mocking the calls to the vm for usage in tests and playground
	selfSym.MockApplication("QmYiQuR2K377j4ZUjxesY6a6oVvnLhzMMJsQSqmxrGZkyG")

	// Called from a function in the application "QmYiQuR2K377j4ZUjxesY6a6oVvnLhzMMJsQSqmxrGZkyG"
	application, err := self.Application()
	if err != nil {
		panic(err)
	}

	fmt.Println(application)
	// Output: QmYiQuR2K377j4ZUjxesY6a6oVvnLhzMMJsQSqmxrGZkyG
}

func ExampleFunction() {
	// Mocking the calls to the vm for usage in tests and playground
	selfSym.MockFunction("QmNtG4SdhqrC3bEtz5euGZeEmaUhQrz6cSJ1LuEyu8Z5NM")

	// Called from a function with id "QmNtG4SdhqrC3bEtz5euGZeEmaUhQrz6cSJ1LuEyu8Z5NM"
	function, err := self.Function()
	if err != nil {
		panic(err)
	}

	fmt.Println(function)
	// Output: QmNtG4SdhqrC3bEtz5euGZeEmaUhQrz6cSJ1LuEyu8Z5NM
}

func ExampleCommit() {
	// Mocking the calls to the vm for usage in tests and playground
	selfSym.MockCommit("189d781643e5efc1d90130fc8e2c526f1040e10d")

	// Called from a function with id "189d781643e5efc1d90130fc8e2c526f1040e10d"
	commit, err := self.Commit()
	if err != nil {
		panic(err)
	}

	fmt.Println(commit)
	// Output: 189d781643e5efc1d90130fc8e2c526f1040e10d
}

func ExampleBranch() {
	// Mocking the calls to the vm for usage in tests and playground
	selfSym.MockBranch("master")

	// Called from a function in the branch "master"
	branch, err := self.Branch()
	if err != nil {
		panic(err)
	}

	fmt.Println(branch)
	// Output: master
}
