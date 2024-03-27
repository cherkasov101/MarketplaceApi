package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/cherkasov101/MarketplaceApi/internal/models"
	"github.com/cherkasov101/MarketplaceApi/internal/services"
)

func CreateAd(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	username, check, err := services.CheckToken(r, db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !check {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var ad models.Ad
	err = json.NewDecoder(r.Body).Decode(&ad)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = services.CreateAd(username, db, ad)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ок"))
}
