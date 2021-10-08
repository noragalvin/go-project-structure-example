package driver

import (
	"ecommerce-integrations/internal/configs"
	"ecommerce-integrations/internal/utils/logger"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDatabaseForOnlineStore(c *ConnectionInfo) *gorm.DB {
	var db *gorm.DB

	// dsn := os.Getenv("DATABASE_URL")
	config := configs.AppConfig
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.DatabaseOnlineStore.Host, config.DatabaseOnlineStore.User, config.DatabaseOnlineStore.Password, config.DatabaseOnlineStore.Name, config.DatabaseOnlineStore.Port)

	conn, err := gorm.Open(postgres.New(
		postgres.Config{
			DSN: dsn,
		},
	), &gorm.Config{})
	if err != nil {
		logger.Logger.Error(err)
		// TODO: push to sentry if error
		// panic(err)
	}
	db = conn
	// db.Debug()
	logger.Logger.Info("CONNECT TO ONLINE STORE DATABASE SUCCESSFULLY!")
	return db
}
