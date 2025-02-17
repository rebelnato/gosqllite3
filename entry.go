package gosqlite3

import (
	"fmt"
	"os"

	// will be used for calling static query functions

	"github.com/rebelnato/gosqlite3/db/connection"
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

var dbConfig Config

func InitiateFlows() {
	data, err := os.ReadFile("./db/config/config.yml")
	if err != nil {
		fmt.Println("Error reading YAML file:", err)
		return
	}

	err = yaml.Unmarshal(data, &dbConfig)
	if err != nil {
		fmt.Println("Error parsing YAML:", err)
		return
	}

	connection.PassConfig(dbConfig.DatabaseConfig.DbName, dbConfig.DatabaseConfig.DbPath)

	db, err := connection.ConnectToSQLiteDB()
	if err != nil {
		fmt.Println("Error connecting to SQLite:", err)
		return
	}
	defer db.Close()

	// Will call the examples function which will initiate terminal level interactive CRUD examples testing
	examples()

	// crud.QueryUserList(db)
	// Print parsed values
	// fmt.Println("Database path:", dbConfig.DatabaseConfig.DbPath)
	// fmt.Println("Database name:", dbConfig.DatabaseConfig.DbName)
}
