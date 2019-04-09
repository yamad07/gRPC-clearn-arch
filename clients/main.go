package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/yamad07/gRPC-sample/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())
	if err != nil {
		log.Fatal("cleint connection error:", err)
	}

	defer conn.Close()
	client := pb.NewCatClient(conn)
	msg := &pb.GetMyCatMessage{TargetCat: "hana"}
	res, err := client.GetMyCat(context.TODO(), msg)

	fmt.Printf("result:%#v \n", res)
	fmt.Printf("error::%#v \n", err)
}
