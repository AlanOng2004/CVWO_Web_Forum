package repositories

import (
	"database/sql"
	"io/ioutil"
	"log"
	"path/filepath"
)

func RunMigrations(db *sql.DB, dir string) {
	files, err := filepath.Glob(filepath.Join(dir, "*.sql"))
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}

		if _, err := db.Exec(string(content)); err != nil {
			log.Fatalf("migration failed (%s): %v", file, err)
		}
	}
}
