package repositories

import (
	"github.com/yamad07/gRPC-sample/entities/models"
	"github.com/yamad07/gRPC-sample/infra"
)

type TodoReposotory interface {
	Save(title string, detail string) (*models.Todo, error)
	Get(id uint64) (*models.Todo, error)
}

type todoRepositoryImpl struct {
}

func NewTodoRepositoryImpl() todoRepositoryImpl {
	return todoRepositoryImpl{}
}

func (repo todoRepositoryImpl) Save(
	title string,
	detail string,
) (*models.Todo, error) {
	t := &models.Todo{
		Title:  title,
		Detail: detail,
	}
	err := infra.DB.Create(t).Error
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (repo todoRepositoryImpl) Get(
	id uint64,
) (*models.Todo, error) {
	t := &models.Todo{}
	t.ID = id
	err := infra.DB.First(t).Error
	if err != nil {
		return nil, err
	}
	return t, nil
}
