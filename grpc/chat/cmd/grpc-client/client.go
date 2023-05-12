package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "chat/internal/pb"
)

const (
	port = ":50051"
)

func main() {
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Connection failed: %v", err)
	}
	defer conn.Close()

	c := pb.NewChatServiceClient(conn)

	biDi(c)
}

func biDi(c pb.ChatServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")

	req := &pb.AddRequest{
		Num1: 5,
		Num2: 10,
	}

	res, err := c.Add(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Add RPC: %v", err)
	}

	log.Printf("Response from Add: %v", res.Result)
}
