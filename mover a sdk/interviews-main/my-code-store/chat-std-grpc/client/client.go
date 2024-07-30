package main

import (
	"context"
	"fmt"
	"log"

	"github.com/devpablocristo/golang-examples/grpc/chat/chatpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	c := chatpb.NewChatServiceClient(conn)

	biDi(c)
}

func biDi(c chatpb.ChatServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")

	req := &chatpb.AddRequest{
		Num1: 5,
		Num2: 10,
	}

	res, err := c.Add(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Add RPC: %v", err)
	}

	log.Printf("Response from Add: %v", res.Result)
}
