package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/tapiaw38/irrigation-api/handlers"
	"github.com/tapiaw38/irrigation-api/storage"
)

func main() {
	// load env file
	err := godotenv.Load()
	if err != nil {
		log.Println(err, "Error loading .env file")
	}

	// connect to database
	db := storage.NewConnection()

	if err := db.DB.Ping(); err != nil {
		log.Fatal(err)
	}

	// start server
	handlers.HandlerServer()
}
