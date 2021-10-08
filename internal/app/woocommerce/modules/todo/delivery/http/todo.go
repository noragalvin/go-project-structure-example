package http

import (
	"ecommerce-integrations/driver"
	"ecommerce-integrations/internal/app/woocommerce/modules/todo/repository"
	"ecommerce-integrations/internal/app/woocommerce/modules/todo/usecase"
	"ecommerce-integrations/internal/utils/app"
	"ecommerce-integrations/internal/utils/constants"
	"ecommerce-integrations/pkg/helpers"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	todoUseCase usecase.TodoUseCase
}

func NewTodoHandler(db *driver.Database) TodoHandler {
	tRepo := repository.NewTodoRepository(db)
	pUseCase := usecase.NewTodoUseCase(tRepo)
	return TodoHandler{
		todoUseCase: pUseCase,
	}
}
func (instance *TodoHandler) Index(c *gin.Context) {
	app := app.Gin{C: c}

	number := c.DefaultQuery("n", "10")

	n, err := helpers.StringToInt(number)

	if err != nil {
		app.BadRequest(gin.H{
			"status_code": constants.INVALID_PARAMS,
		})
		return
	}

	todos, err := instance.todoUseCase.GenerateTodos(n)
	if err != nil {
		app.Error(err.Error())
		return
	}

	app.Ok(gin.H{
		"todos": todos,
	})
}
