package outputs

import "github.com/yamad07/gRPC-sample/entities/models"

type Todo struct {
	Todo *models.Todo
}

func NewTodo(
	todo *models.Todo,
) Todo {
	return Todo{
		todo,
	}
}
