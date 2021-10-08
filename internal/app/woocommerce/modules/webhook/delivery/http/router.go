package http

import (
	"ecommerce-integrations/driver"

	"github.com/gin-gonic/gin"
)

func InitAPIWooCommerceWebhook(r *gin.RouterGroup, db *driver.Database) {
	webhookController := NewWebhookHandler(db)

	productRoutes := r.Group("/webhooks")
	{
		productRoutes.GET("", webhookController.Index)
	}

}
