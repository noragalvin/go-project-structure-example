package driver

import (
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
	Conn   *gorm.DB
	OSConn *gorm.DB
}

// DatabaseWrapper create a *Database
func DatabaseWrapper(db *gorm.DB, dbOS *gorm.DB) *Database {
	return &Database{Conn: db, OSConn: dbOS}
}
