package database

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// BuildDatabase takes a path and builds database directories and files
func BuildDatabase(path string) (err error) {
	err = generateDatabaseGoFiles(path)
	if err != nil {
		return err
	}
	return nil
}

// generateDatabaseGoFiles generates Go files to build configurations
func generateDatabaseGoFiles(path string) (err error) {
	// Creates a directory called database
	path = filepath.Clean(fmt.Sprintf("%s/database", path))
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}

	d1 := []byte(`package database

import (
	"database/sql"
	"errors"
)

type DBService interface {
	CreateConnection(driver, username, password, hostname, databaseName string) error
	CreateDatabase(dbName string) error
	GetConnection() *sql.DB
	CloseConnection() error
}

// Creates an Database object given the driver type
func CreateDBObject(dbDriver string) (DBService, error) {
	switch dbDriver {
	case "mysql":
		return nil, nil
	default:
		return nil, errors.New("unsupported driver")

	}
}
`)
	err = ioutil.WriteFile(filepath.Clean(fmt.Sprintf("%s/database_service.go", path)), d1, 0644)
	if err != nil {
		return err
	}

	return nil
}
