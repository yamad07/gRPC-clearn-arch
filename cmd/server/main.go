package main

import (
	"log"
	"net"

	"github.com/yamad07/gRPC-sample/depin"
	"github.com/yamad07/gRPC-sample/infra"
	pb "github.com/yamad07/gRPC-sample/pb"
	"google.golang.org/grpc"
)

func main() {
	infra.InitDB()
	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	pb.RegisterTodoServiceServer(server, depin.ProvideTodoController)
	server.Serve(listenPort)
}
