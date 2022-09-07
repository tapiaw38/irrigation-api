package test

import (
	"testing"

	"github.com/tapiaw38/irrigation-api/storage"
)

func TestDatabase(t *testing.T) {
	db, err := storage.ConnectPostgres()

	if db == nil || err != nil {
		t.Error("Error connecting to database")
	}
}
