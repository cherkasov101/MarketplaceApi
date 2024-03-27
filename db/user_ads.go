package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// Function to create the table with information about user ads
func CreateUserAdsTable(db *sql.DB) error {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS user_ads (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		ad_id INTEGER,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id),
		FOREIGN KEY (ad_id) REFERENCES ads(id)
	);
`
	_, err := db.Exec(createTableSQL)
	return err
}
