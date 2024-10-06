package utils

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"fmt"
)

var JwtKey = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	Username string `json:"username"`
	UserID   uint   `json:"user_id"`
	jwt.StandardClaims
}

func GenerateJWT(username string, userID uint) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) 
	claims := &Claims{
		Username: username,
		UserID:   userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtKey)
}

func ValidateToken(tokenString string) (*Claims, error) {
    fmt.Println("Token: ", tokenString) 
    fmt.Println("JwtKey: ", JwtKey) 

    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return JwtKey, nil
    })
    if err != nil {
        return nil, err
    }

    if !token.Valid {
        fmt.Println("Invalid token")
        return nil, errors.New("invalid token")
    }

    return claims, nil
}
