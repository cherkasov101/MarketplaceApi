package services

import (
	"database/sql"
	"fmt"
)

type AdForList struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	ImageURL    string  `json:"image_url"`
	Price       float64 `json:"price"`
	Author      string  `json:"author"`
}

// Function to get advertisements
func GetAdvertisements(db *sql.DB, page int, sortBy string, sortDir string) ([]AdForList, error) {
	var ads []AdForList

	// Calculate the number of ads to retrieve based on page number
	adsPerPage := 10
	offset := (page - 1) * adsPerPage

	// Construct SQL query based on sorting criteria
	var query string
	switch sortBy {
	case "price":
		query = "SELECT ads.id, ads.title, ads.description, ads.image_url, ads.price, users.name AS author " +
			"FROM ads " +
			"JOIN user_ads ON ads.id = user_ads.ad_id " +
			"JOIN users ON user_ads.user_id = users.id " +
			"ORDER BY ads.price " + sortDir + " LIMIT ? OFFSET ?"
	default:
		query = "SELECT ads.id, ads.title, ads.description, ads.image_url, ads.price, users.name AS author " +
			"FROM ads " +
			"JOIN user_ads ON ads.id = user_ads.ad_id " +
			"JOIN users ON user_ads.user_id = users.id " +
			"ORDER BY user_ads.created_at DESC LIMIT ? OFFSET ?"
	}

	// Execute the query
	rows, err := db.Query(query, adsPerPage, offset)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	// Iterate through the rows and populate the ads slice
	for rows.Next() {
		var ad AdForList
		if err := rows.Scan(&ad.ID, &ad.Title, &ad.Description, &ad.ImageURL, &ad.Price, &ad.Author); err != nil {
			return nil, err
		}
		ads = append(ads, ad)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ads, nil
}
