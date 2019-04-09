package main

import (
	"log"
	"net"

	controller "github.com/yamad07/gRPC-sample/controllers"
	pb "github.com/yamad07/gRPC-sample/pb"
	"google.golang.org/grpc"
)

func main() {
	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	catCon := &controller.MyCatController{}

	pb.RegisterCatServer(server, catCon)
	server.Serve(listenPort)
}
