package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv          string
	AppPort         string
	DBHost          string
	DBPort          int
	DBUser          string
	DBPassword      string
	DBName          string
	JWTSecret       string
	JWTExpiresHours int
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment")
	}

	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))

	return &Config{
		AppEnv:          os.Getenv("APP_ENV"),
		AppPort:         os.Getenv("APP_PORT"),
		DBHost:          os.Getenv("DB_HOST"),
		DBPort:          dbPort,
		DBUser:          os.Getenv("DB_USER"),
		DBPassword:      os.Getenv("DB_PASSWORD"),
		DBName:          os.Getenv("DB_NAME"),
		JWTSecret:       os.Getenv("JWT_SECRET"),
		JWTExpiresHours: 72, // atau ambil dari env
	}
}
