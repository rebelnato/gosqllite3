package gosqlite3

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/yaml.v3"
)

// Global vars will be used for storing config details
type Database struct {
	DbPath string `yaml:"path"`
	DbName string `yaml:"name"`
}

// Struct to hold the full YAML config
type Config struct {
	DatabaseConfig Database `yaml:"dbConfig"`
}

// Create a var of type Config
var dbConfig Config

/*
 Below function is to be used for reading the config.yml file
 Default path for config file : "<parent directory>/db/config/config.yml"
 config.yml file contains the path and name of db that is to be created if doesn't exists already
 If the config.yml file is not available then it will fail to connect to db
 Post reading the db config it also call ConnectToSQLiteDB in order to stablish a live connection with db
 The connection is then returned to parent function which calls this function

 This function will be called insed each of the CRUD function available under ./crud/staticCRUD.go file

 How to call this function ?
 No need to import for existing CRUD operations as it alreadys gets called from within them to get live db connection

 Follow below steps in case user still wishes to call this function for testing or utilizing the db connection for custom functions.
 Import the package : "github.com/rebelnato/gosqlite3"
 Call this function as db,err := gosqlite3.ReadDbConfig()

 While calling the function note that we assigning return value to 2 vars db and err
 db contains data base connection that can be passed as an argment to custom functions
 err contains error incase any is encountered while performing db connection

 When using this function for custom functions user must close the db connection post completion of teh function
 db connection can be closed using one of below commands
 db.Close() - Immediately closed db connection
 defer db.Close() - Waits for the current function to end in order to close db connection
*/

func ReadDbConfig() (*sql.DB, error) {
	data, err := os.ReadFile("./db/config/config.yml")
	if err != nil {
		fmt.Println("Error reading YAML file:", err)
		return nil, err
	}

	err = yaml.Unmarshal(data, &dbConfig)
	if err != nil {
		fmt.Println("Error parsing YAML:", err)
		return nil, err
	}

	db, err := ConnectToSQLiteDB(dbConfig.DatabaseConfig.DbName, dbConfig.DatabaseConfig.DbPath)
	if err != nil {
		fmt.Println("Connection to database failed")
		return nil, err
	}
	return db, err
}

/*
 Below functions helps with db connection
 Accepts db name and path as argument
 Returns live connection to db
 Error in case any is encountered while performing db connection
*/

func ConnectToSQLiteDB(dbName, dbPath string) (*sql.DB, error) {

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

	_, createTableError := db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, username TEXT NOT NULL UNIQUE, password TEXT);")
	if createTableError != nil {
		fmt.Printf("Failed to create users table eventhough it doesn't exists with error message %q", createTableError)
		return db, createTableError
	}

	return db, err
}
