package storage

import (
	"database/sql"
	"io/ioutil"
	"os"

	_ "github.com/lib/pq"
)

// ConnectPostgres connects to the postgres database
func ConnectPostgres() (*sql.DB, error) {

	client := os.Getenv("DATABASE_URL")

	return sql.Open("postgres", client)
}

// MakeMigration creates the database schema
func MakeMigration(db *sql.DB) error {
	b, err := ioutil.ReadFile("storage/models.sql")

	if err != nil {
		return err
	}

	rows, err := db.Query(string(b))

	if err != nil {
		return err
	}

	return rows.Close()
}
