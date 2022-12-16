package database

import (
	"fmt"

	databaseSym "github.com/taubyte/go-sdk-symbols/database"
)

// New creates a new database
// Returns a database and error
func New(name string) (Database, error) {
	var databaseId uint32
	err := databaseSym.NewDatabase(name, &databaseId)
	if err != 0 {
		return 0, fmt.Errorf("Failed creating new database with: %s", err)
	}

	return Database(databaseId), nil
}
