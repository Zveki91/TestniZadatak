package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "test/proto"
	"time"
)

func main() {
	// Connect to gRPC server
	conn, err := grpc.Dial("localhost:50051")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	// Set a timeout for the request
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Make the RPC call
	response, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Alice"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Response from server: %s", response.GetMessage())
}
