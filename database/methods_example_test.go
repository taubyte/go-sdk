package database_test

import (
	"fmt"
	"reflect"
	"sort"

	symbols "github.com/taubyte/go-sdk-symbols/database"
	"github.com/taubyte/go-sdk/database"
)

func ExampleNew() {
	// Mocking the calls to the vm for usage in tests and playground
	testName := "someDatabase"
	symbols.MockNew(16, testName)

	db, err := database.New(testName)
	if err != nil {
		return
	}

	fmt.Println(db)
	// Output: 16
}

func ExampleDatabase_Close() {
	// Mocking the calls to the vm for usage in tests and playground
	testId := uint32(3)
	symbols.MockClose(testId)

	err := database.Database(testId).Close()

	fmt.Println(err)
	// Output: <nil>
}

func ExampleDatabase_Delete() {
	// Mocking the calls to the vm for usage in tests and playground
	testId := uint32(5)
	testKey := "someKey"
	symbols.MockDelete(testId, testKey, map[string][]byte{})

	err := database.Database(testId).Delete(testKey)

	fmt.Println(err)
	// Output: <nil>
}

func ExampleDatabase_Get() {
	// Mocking the calls to the vm for usage in tests and playground
	testId := uint32(5)
	testKey := "someKey"
	testData := map[string][]byte{
		testKey: []byte("Hello, world!"),
	}
	symbols.MockGet(testId, testData)

	data, err := database.Database(testId).Get(testKey)
	if err != nil {
		return
	}

	fmt.Println(string(data))
	// Output: Hello, world!
}

func ExampleDatabase_List() {
	// Mocking the calls to the vm for usage in tests and playground
	testId := uint32(5)
	testKey := "someKey"
	testData := map[string][]byte{
		testKey + "/a":    {},
		testKey + "/bb":   {},
		testKey + "/cccd": {},
	}
	symbols.MockList(testId, testKey, testData)

	data, err := database.Database(testId).List(testKey)
	if err != nil {
		return
	}

	expected := []string{"someKey/a", "someKey/bb", "someKey/cccd"}
	sort.Strings(expected)
	sort.Strings(data)

	if reflect.DeepEqual(data, expected) == false {
		return
	}

	fmt.Println("Success")
	// Output: Success
}

func ExampleDatabase_Put() {
	// Mocking the calls to the vm for usage in tests and playground
	testId := uint32(18)
	testPut := map[string][]byte{}
	symbols.MockPut(testId, testPut)

	testKey := "someKey"
	err := database.Database(testId).Put(testKey, []byte("Hello, world"))
	if err != nil {
		return
	}

	fmt.Println(string(testPut[testKey]))
	// Output: Hello, world
}
