package storage_test

import (
	"fmt"

	symbols "github.com/taubyte/go-sdk-symbols/storage"
	"github.com/taubyte/go-sdk/storage"
)

func ExampleStorageFile_Read() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.StorageFileMockData{
		StorageId:   5,
		StorageName: "testStorage",
		FileId:      10,
		Data:        []byte("Hello World"),
	}.Mock()

	db, err := storage.New("testStorage")
	if err != nil {
		return
	}

	file := db.File("testFile")
	dbFile, err := file.GetFile()
	if err != nil {
		return
	}

	p := make([]byte, len("Hello World"))
	read, err := dbFile.Read(p)
	if err != nil {
		return
	}

	fmt.Println(read)
	// Output: 11

}

func ExampleStorageFile_Close() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.StorageFileMockData{StorageId: 5,
		StorageName: "testStorage",
	}.Mock()

	db, err := storage.New("testStorage")
	if err != nil {
		return
	}

	file := db.File("testFile")
	dbFile, err := file.GetFile()
	if err != nil {
		return
	}

	err = dbFile.Close()
	if err != nil {
		return
	}

	fmt.Println("Success")
	// Output: Success

}
