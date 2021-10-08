package http

import (
	"ecommerce-integrations/driver"

	"github.com/gin-gonic/gin"
)

func InitAPIWooCommerceTodo(r *gin.RouterGroup, db *driver.Database) {
	todoController := NewTodoHandler(db)

	productRoutes := r.Group("/todos")
	{
		productRoutes.GET("", todoController.Index)
	}

}
