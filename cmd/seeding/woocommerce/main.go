package main

import (
	"ecommerce-integrations/driver"
	"ecommerce-integrations/internal/configs"
	"ecommerce-integrations/internal/utils/logger"
	"ecommerce-integrations/pkg/seeding/woocommerce"
	"log"
)

func main() {
	err := configs.InitWooConfigs()
	if err != nil {
		log.Panicln(err)
	}

	config := configs.AppConfig

	db := driver.OpenDatabaseForWoo(&driver.ConnectionInfo{
		User:     config.Database.User,
		Password: config.Database.Password,
		Host:     config.Database.Host,
		Port:     config.Database.Port,
		Name:     config.Database.Name,
	})

	for _, seed := range woocommerce.All() {
		if err := seed.Run(db); err != nil {
			log.Fatalf("Running seed '%s', failed with error: %s", seed.Name, err)
		} else {
			logger.Logger.Info("Running seed ", seed.Name, " successfully")
		}
	}
}
