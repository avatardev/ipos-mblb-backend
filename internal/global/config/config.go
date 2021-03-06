package config

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddress string

	DBAddress  string
	DBUsername string
	DBPassword string
	DBName     string
	LocalRepo  string
	BaseURL    string

	JWT_ISSUER         string
	JWT_AT_EXPIRATION  time.Duration
	JWT_RT_EXPIRATION  time.Duration
	JWT_SIGNING_METHOD jwt.SigningMethod
	JWT_SIGNATURE_KEY  []byte
}

var config Config

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("[Init] godotenv loading error")
	}

	config.ServerAddress = os.Getenv("SERVER_ADDRESS")

	config.DBAddress = os.Getenv("DB_ADDRESS")
	config.DBUsername = os.Getenv("DB_USERNAME")
	config.DBPassword = os.Getenv("DB_PASSWORD")
	config.DBName = os.Getenv("DB_NAME")

	config.LocalRepo = os.Getenv("LOCALREPO")
	config.BaseURL = os.Getenv("BASEURL")

	config.JWT_ISSUER = os.Getenv("JWT_ISSUER")
	config.JWT_SIGNING_METHOD = jwt.SigningMethodHS256
	config.JWT_SIGNATURE_KEY = []byte(os.Getenv("JWT_SIGNATURE_KEY"))
	config.JWT_AT_EXPIRATION = time.Duration(5) * time.Minute
	config.JWT_RT_EXPIRATION = time.Duration(168) * time.Hour
}

func GetConfig() *Config {
	return &config
}
