package main

import (
	"gempages-admin-server/src/configs"
	"gempages-admin-server/src/driver"
	"gempages-admin-server/src/utils/e"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"

	routers "gempages-admin-server/src/routers"
)

func startHTTP(db *gorm.DB, port string) {
	database := driver.DatabaseWrapper(db)

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
		c.AbortWithStatus(e.ERROR)
	}))

	r.GET("/panic", func(c *gin.Context) {
		// panic with a string -- the custom middleware could save this to a database or report it to the user
		panic("foo")
	})

	routers.InitAPI(r.Group("/api"), database)
	routers.InitHome(r.Group("/"))

	// routes := mux.NewRouter()
	// todoHandler.RegisterHTTP(database, routes)

	// log.Println("Server is running on port " + port)
	// return http.ListenAndServe(fmt.Sprintf(":%s", port), handle)
	r.Run()
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

	startHTTP(db, config.HTTP.Port)
}
