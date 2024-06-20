package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Database struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

type Api struct {
	Port string
}

type Config struct {
	DB  Database
	Api Api
}

var Env = loadConfig()

func loadConfig() *Config {
	godotenv.Load()
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		port = 5432
	}
	return &Config{
		DB: Database{
			Host:     os.Getenv("DB_HOST"),
			Port:     port,
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		},
		Api: Api{},
	}
}
