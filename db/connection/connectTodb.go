package connection

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// setting db path
var dbName, dbPath string

func PassConfig(databaseName string, databasePath string) {
	dbPath = databasePath
	dbName = databaseName
}

func ConnectToSQLiteDB() (*sql.DB, error) {
	// setting db path

	db, err := sql.Open("sqlite3", dbPath+"/"+dbName)
	if err != nil {
		return nil, err
	}

	// Ping the database to ensure connectivity
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to SQLite Database!")
	return db, nil

}
