package storage_test

//ListFiles -> StorageListFilesSize, StorageListFiles

import (
	"fmt"

	symbols "github.com/taubyte/go-sdk-symbols/storage"
	"github.com/taubyte/go-sdk/storage"
)

func ExampleStorage_Cid() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.StorageMockData{
		StorageId:   2,
		StorageName: "testStorage",
		FileId:      15,
		FileName:    "testFile",
		Cid:         "bafybeiavzbz2ugyergky6plyluts2evta5hu6ehhicmuow4zm6vd42pugq",
	}.Mock()

	db, err := storage.New("testStorage")
	if err != nil {
		return
	}

	cid, err := db.Cid("testFile")
	if err != nil {
		return
	}

	fmt.Println(cid)

	// Output: bafybeiavzbz2ugyergky6plyluts2evta5hu6ehhicmuow4zm6vd42pugq
}

func ExampleStorage_File() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.StorageMockData{
		StorageId:   2,
		StorageName: "testStorage",
	}.Mock()
	db, err := storage.New("testStorage")
	if err != nil {
		return
	}

	file := db.File("testFile")

	fmt.Println(file)
	// Output: &{2 testFile 0}
}

func ExampleStorage_ListFiles() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.StorageMockData{
		StorageId:   2,
		StorageName: "testStorage",
		Files:       "file/testFile1/1/file/testFile2/2",
	}.Mock()

	db, err := storage.New("testStorage")
	if err != nil {
		return
	}

	files, err := db.ListFiles()
	if err != nil {
		return
	}

	fmt.Println(files)
	// Output: [{2 testFile1 1} {2 testFile2 2}]
}

func ExampleStorage_RemainingCapacity() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.StorageMockData{
		StorageId:   2,
		StorageName: "testStorage",
		Used:        25,
		Capacity:    100,
	}.Mock()

	db, err := storage.New("testStorage")
	if err != nil {
		return
	}

	remaining, err := db.RemainingCapacity()
	if err != nil {
		return
	}

	fmt.Println(remaining)
	// Output: 75
}
