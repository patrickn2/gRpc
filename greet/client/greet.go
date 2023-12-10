package main

import (
	"context"
	"log"

	pb "github.com/patrickn2/grpc/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{FirstName: "Patrick"})

	if err != nil {
		log.Fatalf("failed to call Greet: %v\n", err)
	}

	log.Printf("Greet result: %s\n", res.Result)
}