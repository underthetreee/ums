package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const (
	signingKey            = "supersecretkey"
	claimsKey  contextKey = "claims"
)

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			log.Println("empty token string")
			return
		}
		token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(signingKey), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "unathorized", http.StatusUnauthorized)
			log.Println(err)
			return
		}

		ctx := context.WithValue(r.Context(), claimsKey, token.Claims.(jwt.MapClaims))
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func GenerateToken(userID string, ttl time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"id":  userID,
		"exp": time.Now().Add(ttl).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(signingKey))
}

func GetUserIDFromClaims(ctx context.Context) (string, error) {
	claims, ok := ctx.Value(claimsKey).(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid claims type")
	}

	userID, ok := claims["id"].(string)
	if !ok {
		return "", fmt.Errorf("unable to extract user id from claims")
	}

	return userID, nil
}
