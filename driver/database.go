package driver

import (
	"fmt"
	"gemtracker/internal/models"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // ...
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
	if os.Getenv("MODE") != "production" {
		dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", c.Host, c.Port, c.User, c.Name, c.Password)
		fmt.Println("=== DB URL ===")
		fmt.Println(dbURI)

		conn, err := gorm.Open("postgres", dbURI)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		db = conn
		db.LogMode(false)
		if err != nil {
			panic(err)
		}
		db = conn
	} else {
		dbURI := os.Getenv("DATABASE_URL")
		fmt.Println("=== DB URL ===")
		fmt.Println(dbURI)

		conn, err := gorm.Open("postgres", dbURI)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		db = conn
		db.LogMode(false)
		if err != nil {
			panic(err)
		}
		db = conn
	}
	db.Debug().AutoMigrate(&models.Error{}, &models.Todo{})
	return db
}
