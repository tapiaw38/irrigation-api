package database

import (
	"context"
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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
func (postgres *PostgresRepository) MakeMigration(databaseUrl string) error {
	m, err := migrate.New(
		"file://database/migrations",
		databaseUrl,
	)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("migrations:", err)
			return nil
		}
		return err
	}

	return nil
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
