package repository

import (
	"ecommerce-integrations/internal/app/models"
)

type TodoRepository interface {
	BulkInsertTodos([]interface{}) ([]models.Todo, error)
}
