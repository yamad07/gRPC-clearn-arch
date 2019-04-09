package usecases

import (
	"github.com/yamad07/gRPC-sample/entities/repositories"
	"github.com/yamad07/gRPC-sample/usecases/inputs"
	"github.com/yamad07/gRPC-sample/usecases/outputs"
)

type TodoUsecase interface {
	CreateTodo(ipt inputs.CreateTodo) (opt outputs.Todo, err error)
	GetTodo(ipt inputs.CreateTodo) (opt outputs.Todo, err error)
}
type todoUsecaseInteractor struct {
	todoRepository repositories.TodoReposotory
}

func (i todoUsecaseInteractor) CreateTodo(ipt inputs.CreateTodo) (*outputs.Todo, error) {
	t, err := i.todoRepository.Save(
		ipt.Title,
		ipt.Detail,
	)
	if err != nil {
		return nil, err
	}
	opt := outputs.NewTodo(
		t,
	)
	return &opt, nil
}
