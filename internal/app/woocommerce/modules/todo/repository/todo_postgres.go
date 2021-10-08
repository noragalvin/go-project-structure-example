package repository

import (
	"ecommerce-integrations/driver"
	"ecommerce-integrations/internal/app/models"

	"github.com/mitchellh/mapstructure"
)

type todoImpl struct {
	database *driver.Database
}

func NewTodoRepository(db *driver.Database) TodoRepository {
	return &todoImpl{
		database: db,
	}
}
func (instance *todoImpl) BulkInsertTodos(someArr []interface{}) ([]models.Todo, error) {
	todos := []models.Todo{}
	mapstructure.Decode(someArr, &todos)
	// err := instance.database.Conn.Create(&todos).Error
	// if err != nil {
	// 	return nil, err
	// }
	return todos, nil
}
