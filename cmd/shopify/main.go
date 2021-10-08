package main

import (
	"ecommerce-integrations/driver"
	todoHttp "ecommerce-integrations/internal/app/shopify/modules/todo/delivery/http"

	"ecommerce-integrations/internal/configs"
	"ecommerce-integrations/internal/utils/logger"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"ecommerce-integrations/internal/utils/constants"
)

func startHTTP(db *gorm.DB, dbOS *gorm.DB, port string) {
	database := driver.DatabaseWrapper(db, dbOS)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		AllowMethods:     []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"Accept", "Accept-Language", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	// r.Use(gin.Recovery())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			// c.String(e.ERROR, fmt.Sprintf("error: %s", err))
			c.JSON(500, gin.H{
				"message": err,
			})
		}
		c.AbortWithStatus(constants.ERROR)
	}))

	r.GET("/panic", func(c *gin.Context) {
		// panic with a string -- the custom middleware could save this to a database or report it to the user
		panic("foo")
	})

	// baseRouter := r.Group("")
	apiRouter := r.Group("api")

	// productHttp.InitAPIWooCommerceProduct(apiRouter, database)
	todoHttp.InitAPIShopifyTodo(apiRouter, database)
	// authHttp.InitAPIWooCommerceAuth(baseRouter, database)

	// TODO: do same with api v2
	// routers.InitHome(r.Group("/api/v2"))

	r.Run()
}

func main() {
	err := configs.InitShopifyConfigs()
	if err != nil {
		log.Panicln(err)
	}

	config := configs.AppConfig

	logger.Logger.Info(config)

	db := driver.OpenDatabaseForShopify(&driver.ConnectionInfo{
		User:     config.Database.User,
		Password: config.Database.Password,
		Host:     config.Database.Host,
		Port:     config.Database.Port,
		Name:     config.Database.Name,
	})

	dbOS := driver.OpenDatabaseForOnlineStore((&driver.ConnectionInfo{
		User:     config.DatabaseOnlineStore.User,
		Password: config.DatabaseOnlineStore.Password,
		Host:     config.DatabaseOnlineStore.Host,
		Port:     config.DatabaseOnlineStore.Port,
		Name:     config.DatabaseOnlineStore.Name,
	}))

	// defer db.Close()

	startHTTP(db, dbOS, config.HTTP.Port)
}
