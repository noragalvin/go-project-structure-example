package http

import (
	"ecommerce-integrations/driver"

	"github.com/gin-gonic/gin"
)

func InitAPIWooCommerceProduct(r *gin.RouterGroup, db *driver.Database) {
	productController := NewProductHandler(db)

	productRoutes := r.Group("/products")
	{
		productRoutes.GET("/generate", productController.GenerateProductsFaker)
		productRoutes.GET("/:id", productController.GetProductById)
		productRoutes.GET("", productController.GetProductList)
	}

}
