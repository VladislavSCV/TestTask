package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/lib/pq"
)

type ServerBD struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"username"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
}

var dbConfig ServerBD

func ConnToDb() (*sql.DB, error) {
	// Read JSON file
	jsonFile, err := os.Open(`internal\secrets\db_config.json`)
	if err != nil {
		return nil, fmt.Errorf("failed to open JSON file: %v", err)
	}
	defer jsonFile.Close()

	// Read file
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read JSON file: %v", err)
	}
	var dbConfig ServerBD
	err = json.Unmarshal(byteValue, &dbConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	fmt.Printf("Config read from JSON: %+v\n", dbConfig)

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open DB connection: %v", err)
	}
	return db, nil
}
