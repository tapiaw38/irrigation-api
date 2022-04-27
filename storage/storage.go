package storage

import (
	"context"
	"database/sql"
	"log"
	"sync"
)

var (
	data *Data
	once sync.Once
)

// Data is the data structure that holds the database connection
type Data struct {
	DB *sql.DB
}

// initDB initializes the database connection
func initDB() {
	db, err := ConnectPostgres()

	if err != nil {
		panic(err)
	}

	err = MakeMigration(db)

	if err != nil {
		panic(err)
	}

	data = &Data{
		DB: db,
	}

	log.Println("Database connection established")
}

// NewConnection returns a new only database connection
func NewConnection() *Data {
	once.Do(initDB)

	return data
}

// CheckConnection checks if the database connection is established
func CheckConnection() bool {
	db, err := data.DB.Conn(context.Background())

	if err != nil {
		panic(err)
	}

	err = db.PingContext(context.Background())

	if err != nil {
		panic(err)
	}

	return err == nil
}
