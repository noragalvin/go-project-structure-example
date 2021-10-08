package driver

import (
	"ecommerce-integrations/internal/app/models"
	"ecommerce-integrations/internal/configs"
	"ecommerce-integrations/internal/utils/logger"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDatabaseForWoo(c *ConnectionInfo) *gorm.DB {
	var db *gorm.DB

	// dsn := os.Getenv("DATABASE_URL")
	config := configs.AppConfig

	logger.Logger.Infoln(config)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.Database.Host, config.Database.User, config.Database.Password, config.Database.Name, config.Database.Port)

	conn, err := gorm.Open(postgres.New(
		postgres.Config{
			DSN: dsn,
		},
	), &gorm.Config{})
	if err != nil {
		logger.Logger.Error(err)
		panic(err)
	}
	db = conn
	if err != nil {
		panic(err)
	}
	db = conn
	db.Debug().AutoMigrate(&models.Shop{}, &models.ShopMeta{}, &models.Product{}, &models.ProductVariant{})
	logger.Logger.Info("CONNECT TO WOO DATABASE SUCCESSFULLY!")

	return db
}
