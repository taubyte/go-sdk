package storage

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	storageSym "github.com/taubyte/go-sdk-symbols/storage"
	"github.com/taubyte/go-sdk/utils/codec"
)

// Get returns the storage with given name.
// Returns the storage if it exists and an error.
func Get(storageName string) (Storage, error) {
	var id uint32
	err := storageSym.StorageGet(storageName, &id)
	if err != 0 {
		return 0, fmt.Errorf("Failed to getting storage: `%s` with: %s, did you create with New()?", storageName, err)
	}

	return Storage(id), nil
}

// ListFiles looks up all saved files in the given storage.
// Returns an array of all Files found in the storage and an error.
func (s Storage) ListFiles() ([]File, error) {
	var size uint32
	var files []File
	err := storageSym.StorageListFilesSize(uint32(s), &size)
	if err != 0 {
		return nil, fmt.Errorf("Failed storage list file size with %v", err)
	}

	var buf = make([]byte, size)
	err = storageSym.StorageListFiles(uint32(s), &buf[0])
	if err != 0 {
		return nil, fmt.Errorf("Failed storage list files with %v", err)
	}

	values := strings.Split(string(buf), "/")
	for idx, value := range values {
		// Incase some person decides to name their file same as our
		if value == "file" && values[idx+1] != "file" {
			versionStr := trimInvisibleChar(values[idx+2])
			version, err := strconv.Atoi(versionStr)
			if err != nil {
				return nil, fmt.Errorf("Failed converting version to int with %v", err)
			}
			files = append(files, File{
				storage: s,
				name:    values[idx+1],
				version: uint32(version),
			})
		}
		continue
	}
	return files, nil
}

// Cid looks up the given filename in the given storage.
// Returns the cid corresponding to the file if found and an error.
func (s Storage) Cid(fileName string) (string, error) {
	var idPtr uint32

	err := storageSym.StorageCidSize(uint32(s), fileName, &idPtr)
	if err != 0 {
		return "", fmt.Errorf("Failed getting cid size for %s with %v", fileName, err)
	}

	__cid := make([]byte, 59)
	err = storageSym.StorageCid(&__cid[0], &idPtr)
	if err != 0 {
		return "", fmt.Errorf("Failed getting cid for %s with %v", fileName, err)
	}

	return string(__cid), nil
}

// File uses the name passed in and creates a new instance of File that holds the storage and filename.
// Returns the File structure.
func (s Storage) File(fileName string) *File {
	return &File{storage: s, name: fileName}
}

// Remaining capacity loops through to given storage and calculates how much space left is available.
// Returns the remaining space available and an error.
func (s Storage) RemainingCapacity() (int, error) {
	var size uint32
	if err := storageSym.StorageUsedSize(uint32(s), &size); err != 0 {
		return 0, fmt.Errorf("Getting used storage failed with: %s", err)
	}

	var _used []byte
	if size != 0 {
		_used = make([]byte, size)
		if err := storageSym.StorageUsed(uint32(s), &_used[0]); err != 0 {
			return 0, fmt.Errorf("Getting used storage failed with: %s", err)
		}
	}

	if err := storageSym.StorageCapacitySize(uint32(s), &size); err != 0 {
		return 0, fmt.Errorf("Getting total capacity failed with: %s", err)
	}

	var _capacity []byte
	if size != 0 {
		_capacity = make([]byte, size)
		if err := storageSym.StorageCapacity(uint32(s), &_capacity[0]); err != 0 {
			return 0, fmt.Errorf("Getting storage capacity failed with: %s", err)
		}
	}

	capacityString, usedString := string(_capacity), string(_used)
	capacity, err0 := strconv.Atoi(capacityString)
	if err0 != nil {
		return 0, err0
	}

	used, err0 := strconv.Atoi(usedString)
	if err0 != nil {
		return 0, err0
	}

	return capacity - used, nil
}

// Version uses current file and given versions to create a new instance of VersionedFile
// Returns the created VersionedFile structure.
func (f *File) Version(version uint32) *VersionedFile {
	file := *f
	file.version = version
	return &VersionedFile{&file}
}

// Current version looks up in the storage the latest version that is stored for that specific file.
// Returns the latest version for the file, if found, and an error.
func (f *File) CurrentVersion() (int, error) {
	var size uint32
	err := storageSym.StorageCurrentVersionSize(uint32(uint32(f.storage)), f.name, &size)
	if err != 0 {
		return 0, fmt.Errorf("Getting current version size failed with %s", err)
	}

	if size == 0 {
		return 0, fmt.Errorf("Size of current version for file %s is 0", f.name)
	}

	var _version = make([]byte, size)
	err = storageSym.StorageCurrentVersion(f.name, &_version[0])
	if err != 0 {
		return 0, fmt.Errorf("Getting current version failed with %s", err)
	}

	versionString := string(_version)

	version, err0 := strconv.Atoi(versionString)
	if err0 != nil {
		return 0, fmt.Errorf("Getting current version failed with error code: %s", err0)
	}

	return version, nil
}

// Add uses the data and overwrite given and adds the file to the storage.
// If overwrite is set to true then the current version number is not updated, and the data for the current version is updated.
// If overwrite is set to false, then for versioning enabled storages a new version is created
// For versioning disabled storages, file is only added if there is no file with the same name in the current storage.
// Returns the current version of the file.
func (f *File) Add(data []byte, overWrite bool) (int, error) {
	if len(data) == 0 {
		return 0, errors.New("Data cannot be empty")
	}

	var _overWrite uint32
	if overWrite {
		_overWrite = 1
	}

	var version uint32

	err := storageSym.StorageAddFile(uint32(f.storage), f.name, &version, &data[0], uint32(len(data)), _overWrite)
	if err != 0 {
		return 0, fmt.Errorf("Adding file failed with error code: %s", err)
	}

	return int(version), nil
}

// Delete uses the current file structure and the given version to delete the specific version of a file from the storage.
// Returns an error
func (f *File) Delete() error {
	if err := storageSym.StorageDeleteFile(uint32(f.storage), f.name, f.version, 0); err != 0 {
		return fmt.Errorf("Deleting file: `%s` failed with: %s", f.name, err)
	}

	return nil
}

// DeleteAllVersions uses the current file and deletes all version of it in the storage.
// Returns an error
func (f *File) DeleteAllVersions() error {
	err0 := storageSym.StorageDeleteFile(uint32(f.storage), f.name, 0, 1)
	if err0 != 0 {
		return fmt.Errorf("Deleting file: `%s` failed with: %s", f.name, err0)
	}

	return nil
}

// ListVersions uses the current file and looks up all version of the file.
// Returns []string of all the versions and an error
func (f *File) Versions() ([]string, error) {
	var size uint32
	err := storageSym.StorageListVersionsSize(uint32(f.storage), f.name, &size)
	if err != 0 || size == 0 {
		return nil, fmt.Errorf("List versions for file: `%s` failed with: %s", f.name, err)
	}
	versions := make([]byte, size)
	err = storageSym.StorageListVersions(uint32(f.storage), f.name, &versions[0])
	if err != 0 || size == 0 {
		return nil, fmt.Errorf("List versions for file: `%s` failed with: %s", f.name, err)
	}

	var conversion []string
	err0 := codec.Convert(versions).To(&conversion)
	if err0 != nil || len(conversion) == 0 {
		return nil, fmt.Errorf("Converting versions for file: `%s` failed with: %s", f.name, err0)
	}

	return conversion, nil
}
