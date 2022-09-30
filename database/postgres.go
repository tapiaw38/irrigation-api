package database

import (
	"context"
	"database/sql"
	"io/ioutil"

	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

// ConnectPostgres connects to the postgres database
func ConnectPostgres(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)

	if err != nil {
		return nil, err
	}

	return &PostgresRepository{db}, nil
}

// MakeMigration creates the database schema
func (postgres *PostgresRepository) MakeMigration() error {
	b, err := ioutil.ReadFile("database/models.sql")

	if err != nil {
		return err
	}

	rows, err := postgres.db.Query(string(b))

	if err != nil {
		return err
	}

	return rows.Close()
}

// CheckConnection db
func (postgres *PostgresRepository) CheckConnection() bool {
	db, err := postgres.db.Conn(context.Background())

	if err != nil {
		panic(err)
	}

	err = db.PingContext(context.Background())

	if err != nil {
		panic(err)
	}

	return err == nil
}
