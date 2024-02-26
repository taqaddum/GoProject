package util

import (
	"GoProject/main/enum"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type ClaimsToken struct {
	ID        int            `json:"id"`
	Username  string         `json:"username"`
	Authority enum.Authority `json:"authority"`
	jwt.RegisteredClaims
}

func SignToken(id int, name string, auth enum.Authority) string {
	var claims = ClaimsToken{
		ID:        id,
		Username:  name,
		Authority: auth,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "admin",
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	return tokenString
}

func ParseToken(tokenString string) (*ClaimsToken, error) {
	token, err := jwt.ParseWithClaims(tokenString, &ClaimsToken{}, func(token *jwt.Token) (any, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if claims, ok := token.Claims.(*ClaimsToken); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
