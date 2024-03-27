package services

import (
	"database/sql"
	"errors"
	"time"

	"github.com/cherkasov101/MarketplaceApi/internal/models"
)

// Function to create ad
func CreateAd(username string, db *sql.DB, ad models.Ad) error {
	if ad.Title == "" || ad.Description == "" || ad.ImageURL == "" || ad.Price == 0 {
		return errors.New("incorrect data")
	}

	if len(ad.Title) > 100 || len(ad.Description) > 1000 {
		return errors.New("incorrect data")
	}

	var userID int
	err := db.QueryRow("SELECT id FROM users WHERE name = ?", username).Scan(&userID)
	if err != nil {
		return err
	}

	result, err := db.Exec("INSERT INTO ads (title, description, image_url, price) VALUES (?, ?, ?, ?)",
		ad.Title, ad.Description, ad.ImageURL, ad.Price)
	if err != nil {
		return err
	}

	adID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO user_ads (user_id, ad_id, created_at) VALUES (?, ?, ?)",
		userID, adID, time.Now())

	return err
}
