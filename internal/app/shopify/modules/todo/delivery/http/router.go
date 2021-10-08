package http

import (
	"ecommerce-integrations/driver"

	"github.com/gin-gonic/gin"
)

func InitAPIShopifyTodo(r *gin.RouterGroup, db *driver.Database) {
	todoController := NewTodoHandler(db)

	todoRoutes := r.Group("/todos")
	{
		todoRoutes.GET("", todoController.Index)
	}

}
