package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	DB_USER     string
	DB_HOST     string
	DB_PORT     int
	DB_DBNAME   string
	DB_PASSWORD string
	JWT_SECRET  string
}

func GetConfig() *AppConfig {
	godotenv.Load()

	DB_PORT, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal(err.Error())
	}

	return &AppConfig{
		DB_DBNAME:   os.Getenv("DB_DBNAME"),
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     DB_PORT,
		DB_USER:     os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		JWT_SECRET:  os.Getenv("JWT_SECRET"),
	}
}
