package usecase

import (
	"ecommerce-integrations/internal/app/models"
	"ecommerce-integrations/internal/app/shopify/modules/todo/repository"

	"github.com/bxcodec/faker/v3"
)

type TodoUseCase struct {
	todoRepo repository.TodoRepository
}

func NewTodoUseCase(pr repository.TodoRepository) TodoUseCase {
	return TodoUseCase{
		todoRepo: pr,
	}
}

func (instance *TodoUseCase) GenerateTodos(n int) ([]models.Todo, error) {
	todoFakeData := make([]interface{}, 0)
	for i := 0; i < n; i++ {
		p := struct {
			Title string `json:"title" faker:"name"`
		}{}
		_ = faker.FakeData(&p)

		todoFakeData = append(todoFakeData, p)
	}

	todos, err := instance.todoRepo.BulkInsertTodos(todoFakeData)

	if err != nil {
		return todos, err
	}
	return todos, nil
}
