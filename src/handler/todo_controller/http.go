package todo_controller

import (
	"gempages-admin-server/src/driver"

	"gempages-admin-server/src/repositories"
	"gempages-admin-server/src/services"

	"gempages-admin-server/src/utils/app"

	"github.com/gin-gonic/gin"
)

// Post ...
type Todo struct {
	repo repositories.TodoRepository
}

// NewTodoHandler ...
func NewTodoHandler(db *driver.Database) *Todo {
	return &Todo{
		repo: services.NewTodoService(db),
	}
}

func (t *Todo) List(c *gin.Context) {
	app := app.Gin{C: c}

	todos, err := t.repo.List(10)
	if err != nil {
		app.Error(err.Error())
		return
	}

	app.Ok(gin.H{
		"todos": todos,
	})
}

// func (t *Todo) Create(w http.ResponseWriter, r *http.Request) {
// 	RespondSuccess(w, Message(true, "success"))
// }

// func (t *Todo) GetByID(w http.ResponseWriter, r *http.Request) {
// 	RespondSuccess(w, Message(true, "success"))
// }

// func (t *Todo) Update(w http.ResponseWriter, r *http.Request) {
// 	RespondSuccess(w, Message(true, "success"))
// }

// func (t *Todo) Delete(w http.ResponseWriter, r *http.Request) {
// 	RespondSuccess(w, Message(true, "success"))
// }
