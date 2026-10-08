package token_manager

import (
	"strconv"
	"time"

	// "github.com/dgrijalva/jwt-go"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("Tm9HbNPPzW946bzJqGSv+opLIKtVqUCFSKq/NbmvgZw=")

// var tokenExpiryDuration = 1 * time.Hour

// get role access.
func getRole(username string) string {
	if username == "senior" {
		return "senior"
	}
	return "employee"
}

func CreateToken(username string) (string, error) {
	// Create a new JWT token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // Token expiration time
		"iat":      time.Now().Unix(),                    // Issued at time
		// "nbf":      time.Now().Unix(), // Not before time
		"iss": "your-issuer",                            // Issuer
		"aud": getRole(username),                        // Audience
		"jti": strconv.FormatInt(time.Now().Unix(), 10), // Unique identifier for the token,
	})

	// Sign the token with the secret key
	tokenString, err := claims.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}



