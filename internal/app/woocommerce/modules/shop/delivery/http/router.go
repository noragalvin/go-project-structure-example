package http

import (
	"ecommerce-integrations/driver"

	"github.com/gin-gonic/gin"
)

func InitAPIWooCommerceShop(r *gin.RouterGroup, db *driver.Database) {
	shopController := NewShopHandler(db)

	shopRoutes := r.Group("/shops")
	{
		shopRoutes.GET("", shopController.Index)
	}
}
