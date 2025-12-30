package repositories

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func NewDB(dataSource string) *sql.DB {
	db, err := sql.Open("sqlite3", dataSource)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
