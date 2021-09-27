package driver

import (
	"fmt"
	"go-project-structure-example/internal/app/models"
	"go-project-structure-example/internal/configs"
	"go-project-structure-example/internal/utils/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ConnectionInfo struct {
	User     string
	Name     string
	Port     string
	Host     string
	Password string
}

type Database struct {
	Conn *gorm.DB
}

// DatabaseWrapper create a *Database
func DatabaseWrapper(db *gorm.DB) *Database {
	return &Database{Conn: db}
}

func OpenDatabase(c *ConnectionInfo) *gorm.DB {
	var db *gorm.DB

	// dsn := os.Getenv("DATABASE_URL")
	config := configs.AppConfig
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
	db.Debug().AutoMigrate(&models.Product{})
	return db
}
