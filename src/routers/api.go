package routers

import (
	"gempages-admin-server/src/driver"
	"gempages-admin-server/src/handler/todo_controller"

	"github.com/gin-gonic/gin"
)

func InitAPI(r *gin.RouterGroup, db *driver.Database) {
	todoController := todo_controller.NewTodoHandler(db)

	apiV1 := r.Group("/v1")
	{
		apiV1.GET("/todos", todoController.List)
	}
}
