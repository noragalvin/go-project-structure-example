package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	HTTP        HTTP
	Database    Database
	Environment string
}

type HTTP struct {
	Host string
	Port string
}

type Database struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

var AppConfig *Config

func Init() error {
	if os.Getenv("MODE") != "production" {
		err := godotenv.Load()
		if err != nil {
			return err
		}
	}
	opt := Config{
		HTTP: HTTP{
			Port: os.Getenv("PORT"),
		},
		Database: Database{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Name:     os.Getenv("DB_NAME"),
			User:     os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
		},
		Environment: os.Getenv("MODE"),
	}

	AppConfig = &opt
	return nil
}
