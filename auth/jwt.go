package auth

import (
	"errors"
	"go-crud/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(config.JWTSecret)

type Claims struct {
    Username string `json:"username"`
    Role     string `json:"role"`
    jwt.StandardClaims
}

// GenerateJWT generates a new JWT for a given username
func GenerateJWT(username string, role string) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        Username: username,
        Role:     role,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

// ValidateJWT validates the given JWT and returns the claims if valid
func ValidateJWT(tokenString string) (*Claims, error) {
    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })

    if err != nil {
        return nil, err
    }

    if !token.Valid {
        return nil, errors.New("invalid token")
    }

    return claims, nil
}
