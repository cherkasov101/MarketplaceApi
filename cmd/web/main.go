package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"

	"github.com/cherkasov101/MarketplaceApi/db"
	"github.com/cherkasov101/MarketplaceApi/internal/handlers"
)

var DB *sql.DB

func main() {
	var err error
	DB, err = db.ConnectDB()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer DB.Close()

	err = db.LaunchDB(DB)
	if err != nil {
		log.Fatal(err)
		return
	}

	r := mux.NewRouter()

	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		handlers.Login(w, r, DB)
	}).Methods(http.MethodPost)
	r.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		handlers.Signup(w, r, DB)
	}).Methods(http.MethodPost)
	r.HandleFunc("/create_ad", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateAd(w, r, DB)
	}).Methods(http.MethodPost)
	r.HandleFunc("/get_ads", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAds(w, r, DB)
	}).Methods(http.MethodGet)

	http.ListenAndServe(":8080", r)
}
