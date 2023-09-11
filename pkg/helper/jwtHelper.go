package helper

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateJWTToken(userID int, username string) (string, error) {
	// Define the token claims
	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // Token expiration time (1 hour)
		"iat":      time.Now().Unix(),                    // Token issue time
	}
	secretKey := os.Getenv("SECRET_KEY")
	var jwtSecret = []byte(secretKey)

	// Create the token with the claims and signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Parse and validate the token
func ParseToken(tokenString string) (*jwt.MapClaims, error) {
	secret := os.Getenv("SECRET_KEY")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Specify the key used for signing the token
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("Invalid token")
	}

	return &claims, nil
}
