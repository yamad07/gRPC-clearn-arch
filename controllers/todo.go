package controllers

import (
	"context"

	pb "github.com/yamad07/gRPC-sample/pb"
	"github.com/yamad07/gRPC-sample/usecases"
	"github.com/yamad07/gRPC-sample/usecases/inputs"
	"google.golang.org/grpc/status"
)

type todoController struct {
	uc usecases.TodoUsecase
}

func NewTodoController(
	uc usecases.TodoUsecase,
) todoController {
	return todoController{
		uc,
	}
}

func (con todoController) CreateTodo(c context.Context, req *pb.CreateTodoRequest) (*pb.CreateTodoResponse, error) {
	ipt := inputs.NewCreateTodo(
		req.Title,
		req.Detail,
	)
	opt, err := con.uc.CreateTodo(ipt)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	res := &pb.CreateTodoResponse{
		Todo: &pb.Todo{
			Id:     opt.Todo.ID,
			Title:  opt.Todo.Title,
			Detail: opt.Todo.Detail,
		},
	}
	return res, nil
}

func (con todoController) GetTodo(c context.Context, req *pb.GetTodoRequest) (*pb.GetTodoResponse, error) {
	ipt := inputs.NewGetTodo(
		req.Id,
	)
	opt, err := con.uc.GetTodo(ipt)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	res := &pb.GetTodoResponse{
		Todo: &pb.Todo{
			Id:     opt.Todo.ID,
			Title:  opt.Todo.Title,
			Detail: opt.Todo.Detail,
		},
	}
	return res, nil
}
