package services

import (
	"gempages-admin-server/src/driver"
	"gempages-admin-server/src/models"
	"gempages-admin-server/src/repositories"
)

// NewSQLPostRepo retunrs implement of post repository interface
func NewTodoService(db *driver.Database) repositories.TodoRepository {
	return &todoImpl{
		database: db,
	}
}

type todoImpl struct {
	database *driver.Database
}

func (instance *todoImpl) List(num int64) ([]*models.Todo, error) {
	var todos []*models.Todo
	instance.database.Conn.Limit(num).Find(&todos)
	return todos, nil
}
func (instance *todoImpl) GetByID(id int64) (*models.Todo, error) {
	var todo *models.Todo
	if err := instance.database.Conn.First(&todo, id).Error; err != nil {
		return nil, err
	}
	return todo, nil
}
func (instance *todoImpl) Create(t *models.Todo) (int64, error) {
	instance.database.Conn.Create(t)
	return int64(t.ID), nil
}
func (instance *todoImpl) Update(t *models.Todo) (*models.Todo, error) {
	instance.database.Conn.Update(t)
	return t, nil
}
func (instance *todoImpl) Delete(id int64) (bool, error) {
	var todo *models.Todo
	if err := instance.database.Conn.First(&todo, id).Error; err != nil {
		return false, err
	}
	instance.database.Conn.Delete(&todo)
	return true, nil
}
