package http

import (
	"gemtracker/driver"
	"net/http"

	"gemtracker/internal/repositories"
	"gemtracker/internal/services"

	"github.com/gorilla/mux"
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

func (t *Todo) List(w http.ResponseWriter, r *http.Request) {
	todos, err := t.repo.List(10)
	if err != nil {
		data := Message(false, err.Error())
		RespondBadRequest(w, data)
		return
	}
	data := Message(true, "success")
	data["todos"] = todos

	RespondSuccess(w, data)
}

func (t *Todo) Create(w http.ResponseWriter, r *http.Request) {
	RespondSuccess(w, Message(true, "success"))
}

func (t *Todo) GetByID(w http.ResponseWriter, r *http.Request) {
	RespondSuccess(w, Message(true, "success"))
}

func (t *Todo) Update(w http.ResponseWriter, r *http.Request) {
	RespondSuccess(w, Message(true, "success"))
}

func (t *Todo) Delete(w http.ResponseWriter, r *http.Request) {
	RespondSuccess(w, Message(true, "success"))
}

func RegisterTodo(todoHandler *Todo, routes *mux.Router) {
	routes.HandleFunc("/todos", todoHandler.List).Methods("GET")
	routes.HandleFunc("/todos", todoHandler.Create).Methods("POST")
	routes.HandleFunc("/todos/{id}", todoHandler.GetByID).Methods("GET")
	routes.HandleFunc("/todos/{id}", todoHandler.Update).Methods("PUT")
	routes.HandleFunc("/todos/{id}", todoHandler.Delete).Methods("DELETE")
}
