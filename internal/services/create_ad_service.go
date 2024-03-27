package services

import (
	"database/sql"
	"errors"
	"time"

	"github.com/cherkasov101/MarketplaceApi/internal/models"
)

func CreateAd(username string, db *sql.DB, ad models.Ad) error {
	if ad.Title == "" || ad.Description == "" || ad.ImageURL == "" || ad.Price == 0 {
		return errors.New("incorrect data")
	}

	if len(ad.Title) > 100 || len(ad.Description) > 1000 {
		return errors.New("incorrect data")
	}
	// Получаем ID пользователя по его имени
	var userID int
	err := db.QueryRow("SELECT id FROM users WHERE name = ?", username).Scan(&userID)
	if err != nil {
		return err
	}

	// Создаем запись объявления в таблице ads
	result, err := db.Exec("INSERT INTO ads (title, description, image_url, price) VALUES (?, ?, ?, ?)",
		ad.Title, ad.Description, ad.ImageURL, ad.Price)
	if err != nil {
		return err
	}

	// Получаем ID нового объявления
	adID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	// Создаем запись в таблице user_ads, связывающую пользователя и объявление
	_, err = db.Exec("INSERT INTO user_ads (user_id, ad_id, created_at) VALUES (?, ?, ?)",
		userID, adID, time.Now())

	return err
}
