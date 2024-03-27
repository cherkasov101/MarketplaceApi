package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Function to connect to the database
func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "db/marketplace.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}
