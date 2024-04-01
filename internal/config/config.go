package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type PostgresConfig struct {
	DBName   string
	SSLMode  string
	Host     string
	Password string
	Port     string
	Username string
}

type Config struct {
	Pg        PostgresConfig
	JwtSecret string
	ApiPort   int
}

func MustLoad() *Config {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	pgConfig := PostgresConfig{
		DBName:   os.Getenv("DB_NAME"),
		Host:     os.Getenv("DB_HOST"),
		Password: os.Getenv("DB_PASSWORD"),
		Port:     os.Getenv("DB_PORT"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		Username: os.Getenv("DB_USERNAME"),
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	apiPort, _ := strconv.Atoi(os.Getenv("API_PORT"))

	return &Config{ApiPort: apiPort, JwtSecret: jwtSecret, Pg: pgConfig}
}
