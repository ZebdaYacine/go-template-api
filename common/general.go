package common

import (
	"go-template-api/config"
	"go-template-api/model"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Id uint32 `json:"id"`
	jwt.RegisteredClaims
}

func GenetatJwt(user model.User_Table) (string, error) {
	secretKey := []byte(config.GetSecrtKey())
	expirationTime := time.Now().Add(30 * time.Minute)
	claims := &Claims{
		Id: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)

	return tokenString, err
}
