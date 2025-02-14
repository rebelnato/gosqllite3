package main

import (
	"fmt"
	"os"

	"github.com/rebelnato/gosqllite3/db/connection"
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

func main() {
	data, err := os.ReadFile("./db/config/config.yml")
	if err != nil {
		fmt.Println("Error reading YAML file:", err)
		return
	}

	db, err := connection.ConnectToSQLiteDB()
	if err != nil {
		fmt.Println("Error connecting to SQLite:", err)
		return
	}
	defer db.Close()

	err = yaml.Unmarshal(data, &dbConfig)
	if err != nil {
		fmt.Println("Error parsing YAML:", err)
		return
	}

	// Print parsed values
	// fmt.Println("Database path:", dbConfig.DatabaseConfig.DbPath)
	// fmt.Println("Database name:", dbConfig.DatabaseConfig.DbName)
}

func PassConfig() (dbName string, dbPath string) {
	return dbConfig.DatabaseConfig.DbName, dbConfig.DatabaseConfig.DbPath
}
