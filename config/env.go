package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type IDatabase struct {
	Name string
	User string
	Pass string
	Host string
	Port string
}

type IEnv struct {
	Database   IDatabase
	ServerPort string
	JwtSecret  string
}

var Env = initConfig()

func initConfig() *IEnv {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &IEnv{
		Database: IDatabase{
			Name: os.Getenv("DB_NAME"),
			User: os.Getenv("DB_USER"),
			Pass: os.Getenv("DB_PASSWORD"),
			Host: os.Getenv("DB_HOST"),
			Port: os.Getenv("DB_PORT"),
		},
		ServerPort: os.Getenv("PORT"),
		JwtSecret:  os.Getenv("JWT_SECRET"),
	}
}
