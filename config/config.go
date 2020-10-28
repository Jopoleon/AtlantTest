package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const DefaultPort = "8080"

type Config struct {
	DBConfig        *DB
	HttpPort        string
	ProductionStart string //is proto_server starts in production mode, or just local debug run
}

type DB struct {
	DBUser string
	DBPass string
	DBName string
	DBHost string
	DBPort string
}

func NewConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("[WARNING] no .env file, reading config from OS ENV variables, error: ", err)
	}
	cfg := &Config{
		HttpPort:        os.Getenv("HTTP_PORT"),
		ProductionStart: os.Getenv("PRODUCTION_START"),
		DBConfig: &DB{
			DBUser: os.Getenv("DB_USER"),
			DBPass: os.Getenv("DB_PASSWORD"),
			DBName: os.Getenv("DB_NAME"),
			DBHost: os.Getenv("DB_HOST"),
			DBPort: os.Getenv("DB_PORT"),
		},
	}
	if cfg.HttpPort == "" {
		cfg.HttpPort = DefaultPort
	}
	return cfg
}
