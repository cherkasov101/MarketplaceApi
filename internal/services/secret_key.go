package services

import (
	"database/sql"
)

// Function to get secret key from the database
func GetSecretKey(db *sql.DB) ([]byte, error) {
	var jwtSecret string
	query := "SELECT secret_key FROM jwt_secrets WHERE id = 1"
	err := db.QueryRow(query).Scan(&jwtSecret)
	if err != nil {
		return nil, err
	}
	result := []byte(jwtSecret)
	return result, nil
}
