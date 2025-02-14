package connection

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rebelnato/gosqllite3/main"
)

// setting db path
var dbName, dbPath string = main.PassConfig()

func ConnectToSQLiteDB() (*sql.DB, error) {

	db, err := sql.Open("sqlite3", "./db/mydb.db")
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
