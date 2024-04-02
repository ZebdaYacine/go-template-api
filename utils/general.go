package utils

import (
	"fmt"
	"go-template-api/config"
	"go-template-api/model"
	"log"
	"math/rand/v2"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	jwt.RegisteredClaims
	IdAccount string    `json:"IdAccount"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

func GenetatJwt(user model.User_Table) (string, error) {
	secretKey := []byte(config.GetSecrtKey())
	expirationTime := time.Now().Add(30 * time.Minute)
	var account string
	if user.Phone != "" {
		account = user.Phone
	} else {
		account = user.Email
	}
	claims := &Claims{
		IdAccount: account,
		Password:  user.Password,
		CreatedAt: time.Now(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	return tokenString, err
}

func VerifyToken(tokenString string) error {
	secretKey := []byte(config.GetSecrtKey())
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		log.Fatal("Invalid token")
	}
	return nil
}

func GenerateNewCode(id uint32) string {
	return strconv.Itoa(rand.IntN(100)) + "UseR" + strconv.FormatUint(uint64(id), 10)
}

func PrintError(message string) {
	red := "\033[31m"
	reset := "\033[0m"
	fmt.Println(red + message + reset)
}

func PrintInfo(message string) {
	red := "\033[34m"
	reset := "\033[0m"
	fmt.Println(red + message + reset)
}
