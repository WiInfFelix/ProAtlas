package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

var DB *sql.DB

func SetUpDB() (err error) {

	log.Println("Setting up DB...")

	DB, err = sql.Open("sqlite3", "./database.db")

	if err != nil {
		log.Fatal(err)
	}

	err = initModels()
	if err != nil {
		log.Fatal("bla")
	}

	return
}

func initModels() (err error) {

	modelsString, err := os.ReadFile("database/init_db.sql")

	fmt.Printf("Script: %s \n", string(modelsString))

	_, err = GetDB().Exec(string(modelsString))

	if err != nil {
		log.Fatalf("Could not initialize models: \n %s", err)
	}

	return
}

func GetDB() *sql.DB {
	if DB == nil {
		SetUpDB()
	}

	return DB
}
