package depin

import (
	"github.com/yamad07/gRPC-sample/controllers"
	"github.com/yamad07/gRPC-sample/entities/repositories"
	pb "github.com/yamad07/gRPC-sample/pb"
	"github.com/yamad07/gRPC-sample/usecases"
)

var ProvideTodoController pb.TodoServiceServer = controllers.NewTodoController(
	usecases.NewTodoUsecaseInteractor(
		repositories.NewTodoRepositoryImpl(),
	),
)
