package database_test

import (
	"fmt"

	symbols "github.com/taubyte/go-sdk-symbols/database"
	"github.com/taubyte/go-sdk/database"
	"github.com/taubyte/go-sdk/event"
)

var (
	testId   = uint32(5)
	testName = "someDatabase"
	testData = map[string][]byte{}
)

func databaseFunction(e event.Event) uint32 {
	db, err := database.New(testName)
	if err != nil {
		return 1
	}

	err = db.Put("value/hello", []byte("Hello, world"))
	if err != nil {
		return 1
	}

	err = db.Put("value/hello2", []byte("Hello, world"))
	if err != nil {
		return 1
	}

	keys, err := db.List("value")
	if len(keys) != 2 || err != nil {
		return 1
	}

	data, err := db.Get("value/hello")
	if err != nil {
		return 1
	}

	if string(data) != "Hello, world" {
		return 1
	}

	err = db.Delete("value/hello")
	if err != nil {
		return 1
	}

	data, err = db.Get("value/hello")
	if err == nil {
		return 1
	}

	err = db.Close()
	if err != nil {
		return 1
	}

	return 0
}

func ExampleDatabase() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.Mock(testId, testName, testData)

	e := databaseFunction(event.Event(1))
	if e != 0 {
		fmt.Println(e)
		return
	}

	fmt.Println("Success")
	// Output: Success
}
