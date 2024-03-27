package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// Function to create the users table
func CreateUsersTable(DB *sql.DB) error {
	_, err := DB.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT, password TEXT)")
	return err
}
