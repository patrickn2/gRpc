package main

import (
	"log"

	pb "github.com/patrickn2/grpc/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v\n", err)
	}

	c := pb.NewGreetServiceClient(conn)

	doGreet(c)


	defer conn.Close()
	
}