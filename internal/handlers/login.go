package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/cherkasov101/MarketplaceApi/internal/models"
	"github.com/cherkasov101/MarketplaceApi/internal/services"
)

// Function to login
func Login(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var credentials models.User
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !verifyUser(credentials, db) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	token, err := services.GenerateToken(credentials.Name, db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := services.TokenResponse{Token: token}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Function to check login and password in the database
func verifyUser(credentials models.User, db *sql.DB) bool {
	var count int
	row := db.QueryRow("SELECT COUNT(*) FROM users WHERE name = ? AND password = ?", credentials.Name, credentials.Password)
	err := row.Scan(&count)
	if err != nil {
		return false
	}

	if count == 0 {
		return false
	}

	return true
}
