package auth

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWT secret key
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// GenerateToken creates a new JWT token
func GenerateToken(userID int) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// VerifyToken verifies the JWT token
func VerifyToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
}
