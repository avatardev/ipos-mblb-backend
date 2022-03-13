package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddress string

	DBAddress  string
	DBUsername string
	DBPassword string
	DBName     string
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
}

func GetConfig() *Config {
	return &config
}
