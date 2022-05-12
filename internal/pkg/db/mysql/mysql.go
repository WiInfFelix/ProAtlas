package database

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
)

var Db *sql.DB

func InitDB() {

	log.Println("Starting DB...")

	db, err := sql.Open("mysql", "root:dbpass@tcp(0.0.0.0)/proatlas")
	if err != nil {
		log.Panicf("Error getting DB connection: %s", err)
	}

	if err = db.Ping(); err != nil {
		log.Panicf("Error pinging DB connection: %s", err)
	}

	Db = db

}

func Migrate() {
	log.Println("Migrating DB...")

	if err := Db.Ping(); err != nil {
		log.Fatal(err)
	}
	driver, _ := mysql.WithInstance(Db, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://internal/pkg/db/migrations/mysql",
		"mysql",
		driver,
	)
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
