package configs

import (
	"os"
)

func InitShopifyConfigs() error {
	// if os.Getenv("MODE") != "production" {
	// 	err := godotenv.Load()
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	opt := Config{
		HTTP: HTTP{
			Host:     os.Getenv("HOST"),
			Port:     os.Getenv("PORT"),
			Protocol: os.Getenv("PROTOCOL"),
		},
		Database: Database{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Name:     os.Getenv("DB_NAME"),
			User:     os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
		},
		DatabaseOnlineStore: Database{
			Host:     os.Getenv("DB_ONLINE_STORE_HOST"),
			Port:     os.Getenv("DB_ONLINE_STORE_PORT"),
			Name:     os.Getenv("DB_ONLINE_STORE_NAME"),
			User:     os.Getenv("DB_ONLINE_STORE_USERNAME"),
			Password: os.Getenv("DB_ONLINE_STORE_PASSWORD"),
		},
		Environment: os.Getenv("MODE"),
	}
	AppConfig = &opt

	return nil
}
