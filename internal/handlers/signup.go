package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/cherkasov101/MarketplaceApi/internal/models"
)

// Function to signup
func Signup(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !checkData(user, db) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = createUser(db, user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	type SignupResponse struct {
		Status   string `json:"status"`
		Username string `json:"username"`
	}

	response := SignupResponse{
		Status:   "ok",
		Username: user.Name,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Function to check name and password
func checkData(user models.User, db *sql.DB) bool {
	if user.Name == "" || user.Password == "" ||
		len(user.Name) < 3 || len(user.Password) < 3 ||
		len(user.Name) > 20 || len(user.Password) > 20 ||
		strings.ContainsAny(user.Name, " ") || strings.ContainsAny(user.Password, " ") {
		return false
	}

	var count int
	row := db.QueryRow("SELECT COUNT(*) FROM users WHERE name = ?", user.Name)
	err := row.Scan(&count)
	if err != nil {
		return false
	}

	if count > 0 {
		return false
	}

	return true
}

// Function to add user to the database
func createUser(db *sql.DB, user models.User) error {
	stmt, err := db.Prepare("INSERT INTO users (name, password) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Password)
	if err != nil {
		return err
	}

	return nil
}
