package storage_test

import (
	"fmt"

	symbols "github.com/taubyte/go-sdk-symbols/storage"
	"github.com/taubyte/go-sdk/storage"
)

func ExampleFile_GetFile() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.FileMockData{
		StorageId:   5,
		StorageName: "testStorage",
		Id:          10,
		FileName:    "testFile",
		FileId:      25,
	}.Mock()

	db, err := storage.New("testStorage")
	if err != nil {
		return
	}

	file := db.File("testFile")

	_, err = file.GetFile()
	if err != nil {
		return
	}

	fmt.Println("Success")
	// Output: Success
}

func ExampleFile_Delete() {
	// Mocking the calls to the vm for usage in tests and playground
	db, err := storage.New("testStorage")
	if err != nil {
		return
	}

	file := db.File("testFile")
	err = file.Delete()
	if err != nil {
		return
	}

	fmt.Println("Success")
	// Output: Success
}

func ExampleFile_DeleteAllVersions() {
	db, err := storage.New("testStorage")
	if err != nil {
		return
	}

	file := db.File("testFile")
	err = file.DeleteAllVersions()
	if err != nil {
		return
	}

	fmt.Println("Success")
	// Output: Success
}

func ExampleFile_Version() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.FileMockData{
		StorageId:   5,
		StorageName: "testStorage",
		Id:          10,
		FileName:    "testFile",
		FileId:      25,
	}.Mock()

	db, err := storage.New("testStorage")
	if err != nil {
		return
	}

	file := db.File("testFile")
	versionedFile := file.Version(12)

	if versionedFile == nil {
		return
	}

	fmt.Println("Success")
	// Output: Success
}

func ExampleFile_CurrentVersion() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.FileMockData{
		StorageId:   5,
		StorageName: "testStorage",
		FileName:    "testFile",
		Version:     2,
	}.Mock()

	db, err := storage.New("testStorage")
	if err != nil {
		return
	}

	file := db.File("testFile")
	version, err := file.CurrentVersion()
	if err != nil {
		return
	}

	fmt.Println(version)
	// Output: 2
}

func ExampleFile_Add() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.FileMockData{
		StorageId:   5,
		StorageName: "testStorage",
		FileName:    "testFile",
		Version:     2,
		Data:        []byte("Hello World"),
	}.Mock()

	db, err := storage.New("testStorage")
	if err != nil {
		return
	}

	file := db.File("testFile")
	version, err := file.Add([]byte("Hello World"), true)
	if err != nil {
		return
	}

	fmt.Println(version)
	// Output: 3
}

func ExampleFile_Versions() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.FileMockData{
		StorageId:   5,
		StorageName: "testStorage",
		FileName:    "testFile",
		Versions:    map[string][]string{"testFile": {"1", "2", "3"}},
	}.Mock()

	db, err := storage.New("testStorage")
	if err != nil {
		return
	}

	file := db.File("testFile")
	versions, err := file.Versions()
	if err != nil {
		return
	}

	fmt.Println(versions)
	// Output: [1 2 3]
}
