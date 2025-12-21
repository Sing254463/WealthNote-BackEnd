package jwt

import (
	"errors"
	"time"

	"WealthNoteBackend/internal/config"

	"github.com/golang-jwt/jwt/v4"
)

// GenerateToken generates a new JWT token with a specified expiration time.
func GenerateToken(userID string, isLongLived bool) (string, error) {
	var expirationTime time.Duration
	if isLongLived {
		duration, err := time.ParseDuration(config.AppConfig.JWTExpirationLong)
		if err != nil {
			expirationTime = 15 * 24 * time.Hour // fallback to 15 days
		} else {
			expirationTime = duration
		}
	} else {
		duration, err := time.ParseDuration(config.AppConfig.JWTExpirationShort)
		if err != nil {
			expirationTime = 15 * time.Minute // fallback to 15 minutes
		} else {
			expirationTime = duration
		}
	}

	claims := &jwt.RegisteredClaims{
		Subject:   userID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expirationTime)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.AppConfig.JWTSecret))
}

// ValidateToken validates the JWT token and returns the claims if valid.
func ValidateToken(tokenString string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(config.AppConfig.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid claims")
}
