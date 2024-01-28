package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var signingKey = "super secret key"

func NewToken(userID string, ttl time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"id":  userID,
		"exp": time.Now().Add(ttl).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(signingKey))
}
