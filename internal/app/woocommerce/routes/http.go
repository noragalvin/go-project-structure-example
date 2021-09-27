package routes

import (
	"go-project-structure-example/driver"
	"go-project-structure-example/internal/app/woocommerce/modules/product/delivery/http"

	"github.com/gin-gonic/gin"
)

// var Router *gin.Engine

// func NewHTTPRouter() *gin.Engine {
// 	router := gin.Default()
// 	// This handler will match /user/john but will not match /user/ or /user
// 	router.GET("/test", func(c *gin.Context) {
// 		name := c.Param("name")
// 		c.String(http.StatusOK, "Hello %s", name)
// 	})

// 	return router
// }

func InitAPIWooCommerce(r *gin.RouterGroup, db *driver.Database) {
	productController := http.NewProductHandler(db)

	productRoutes := r.Group("/products")
	{
		productRoutes.GET("/generate", productController.GenerateProductsFaker)
		productRoutes.GET("/:id", productController.GetProductById)
		productRoutes.GET("/", productController.GetProductList)
	}

}
