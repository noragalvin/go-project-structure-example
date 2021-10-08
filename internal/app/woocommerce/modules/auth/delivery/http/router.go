package http

import (
	"ecommerce-integrations/driver"

	"github.com/gin-gonic/gin"
)

func InitAPIWooCommerceAuth(r *gin.RouterGroup, db *driver.Database) {
	authController := NewAuthHandler(db)

	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/callback", authController.Callback)
		authRoutes.GET("", authController.RequestAuthorizeApp)
	}

}
