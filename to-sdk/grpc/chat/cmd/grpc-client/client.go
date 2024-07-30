// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"

// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials/insecure"

// 	pb "chat/internal/pb"
// )

// const (
// 	port = ":50051"
// )

// func main() {
// 	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		log.Fatalf("Connection failed: %v", err)
// 	}
// 	defer conn.Close()

// 	c := pb.NewChatServiceClient(conn)

// 	biDi(c)
// }

// func biDi(c pb.ChatServiceClient) {
// 	fmt.Println("Starting to do a Unary RPC...")

// 	req := &pb.AddRequest{
// 		Num1: 5,
// 		Num2: 10,
// 	}

// 	res, err := c.Add(context.Background(), req)
// 	if err != nil {
// 		log.Fatalf("Error while calling Add RPC: %v", err)
// 	}

// 	log.Printf("Response from Add: %v", res.Result)
// }

package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "chat/internal/pb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := pb.NewChatServiceClient(conn)

	stream, err := client.Chat(context.Background())
	if err != nil {
		log.Fatalf("Error opening chat stream: %v", err)
	}

	// Start a goroutine to listen for server messages
	go func() {
		for {
			msg, err := stream.Recv()
			if err != nil {
				log.Fatalf("Error receiving message from server: %v", err)
			}
			fmt.Printf("%s: %s\n", msg.Sender, msg.Message)
		}
	}()

	// Main loop to send messages to the server
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter your message (or 'exit' to quit): ")
		scanner.Scan()
		text := scanner.Text()

		if text == "exit" {
			break
		}

		msg := &pb.ChatMessage{
			Sender:    "Client",
			Message:   text,
			Timestamp: time.Now().Unix(),
		}

		if err := stream.Send(msg); err != nil {
			log.Fatalf("Failed to send message to server: %v", err)
		}
	}
}
