package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// Function to create the users table
func CreateAdsTable(DB *sql.DB) error {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS ads (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT,
		description TEXT,
		image_url TEXT,
		price REAL
	);
`
	_, err := DB.Exec(createTableSQL)
	return err
}
