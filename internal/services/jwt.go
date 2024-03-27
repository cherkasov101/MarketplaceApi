package services

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/mattn/go-sqlite3"
)

// Data for JWT token
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type TokenResponse struct {
	Token string `json:"token"`
}

// Function to generate token
func GenerateToken(username string, db *sql.DB) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtKey, err := GetSecretKey(db)
	if err != nil {
		return "", err
	}

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// Function to check token and return username
func CheckToken(r *http.Request, db *sql.DB) (string, bool, error) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return "", false, nil
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		jwtKey, err := GetSecretKey(db)
		if err != nil {
			return nil, err
		}
		return jwtKey, nil
	})
	if err != nil {
		return "", false, err
	}

	if !token.Valid {
		return "", false, nil
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", false, nil
	}

	username, ok := claims["username"].(string)
	if !ok {
		return "", false, nil
	}

	return username, true, nil
}
