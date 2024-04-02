package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func LoadEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	godotenv.Load(".env")
}

type SettingUrl struct {
	Ip_server   string
	Port_server string
}

type Conn struct {
	db    *gorm.DB
	Error error
}

func GetURL() string {
	LoadEnv()
	url := SettingUrl{Ip_server: os.Getenv("SERVER_ADDRESS"), Port_server: os.Getenv("SERVER_PORT")}
	return url.Ip_server + ":" + url.Port_server
}

func GetSecrtKey() string {
	LoadEnv()
	return os.Getenv("SECRET_KEY")
}

func GetConn() Conn {
	return Conn{db: nil, Error: nil}
}
