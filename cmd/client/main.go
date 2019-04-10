package main

import (
	"context"
	"log"

	pb "github.com/yamad07/gRPC-sample/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("application:19003", grpc.WithInsecure())
	if err != nil {
		log.Fatal("cleint connection error:", err)
	}

	defer conn.Close()
	client := pb.NewTodoServiceClient(conn)
	// msg := &pb.GetTodoRequest{
	// 	Id: 1,
	// }

	msg := &pb.CreateTodoRequest{
		Title:  "Write Code",
		Detail: "Write Code",
	}
	res, err := client.CreateTodo(context.TODO(), msg)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(res.Todo)
	}
}
