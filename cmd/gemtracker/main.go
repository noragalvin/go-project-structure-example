package main

import (
	"fmt"
	"gemtracker/driver"
	"gemtracker/internal/configs"
	"gemtracker/internal/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/rs/cors"
)

func startHTTP(db *gorm.DB, port string) error {
	database := driver.DatabaseWrapper(db)

	routes := mux.NewRouter()
	handler.RegisterHTTP(database, routes)

	corsOptions := cors.Options{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowedMethods:   []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Accept-Language", "Content-Type", "Authorization"},
		AllowCredentials: true,
		Debug:            false,
	}

	// Cors domain
	handle := cors.New(corsOptions).Handler(routes)

	log.Println("Server is running on port " + port)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), handle)
}

func main() {

	if err := configs.Init(); err != nil {
		log.Panicln(err)
	}

	config := configs.AppConfig

	db := driver.OpenDatabase(&driver.ConnectionInfo{
		User:     config.Database.User,
		Password: config.Database.Password,
		Host:     config.Database.Host,
		Port:     config.Database.Port,
		Name:     config.Database.Name,
	})

	defer db.Close()

	if err := startHTTP(db, config.HTTP.Port); err != nil {
		log.Panicln(err)
	}
}
