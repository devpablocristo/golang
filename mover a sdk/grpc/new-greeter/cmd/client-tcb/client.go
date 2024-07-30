package main

import (
	"context"
	"io"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/joho/godotenv"

	ctypes "github.com/devpablocristo/qh/internal/platform/custom-types"
	gconfig "github.com/devpablocristo/qh/internal/prompt/config"
	pb "github.com/devpablocristo/qh/internal/prompt/proto"
)

// Main function to run the client
func main() {
	rootDir, err := os.Getwd()
	if err != nil {
		ctypes.HandleFatalError("Error getting current working directory: %v", err)
	}

	envFile := rootDir + "/.env"
	err = godotenv.Load(envFile)
	if err != nil {
		ctypes.HandleFatalError("error loading .env file", err)
	}

	// Connect to the gRPC server using insecure credentials
	grGrpcConfig := gconfig.NewGrpcConfig()
	conn, err := grpc.Dial(grGrpcConfig.GetServerHost()+":"+grGrpcConfig.GetServerPort(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	// Close the connection when the main function returns
	defer conn.Close()

	// Create a new GreetServiceClient using the connection
	client := pb.NewTeamcuBotClient(conn)

	// Call the callGreetBi function with the client and names
	callGreetBi(client)
}

// callGreetBi function sends requests to the server and receives responses using bidirectional streaming
func callGreetBi(client pb.TeamcuBotClient) {
	// Log the start of the bidirectional streaming
	log.Printf("Bidirectional Streaming started")

	// Create a new stream with the server
	stream, err := client.AIInitProcess(context.Background())
	if err != nil {
		log.Fatalf("could not send names: %v", err)
	}

	// Create a channel to synchronize the completion of receiving messages
	waitc := make(chan struct{})

	// Start a goroutine to receive messages from the server
	go func() {
		for {
			// Receive a message from the server
			message, err := stream.Recv()
			// If the server has stopped sending messages, break out of the loop
			if err == io.EOF {
				break
			}
			// If there's an error receiving the message, log the error
			if err != nil {
				log.Fatalf("error while streaming %v", err)
			}
			// Log the received message
			log.Println(message)
		}
		// Close the channel when done receiving messages
		close(waitc)
	}()

	// Print the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}

	// Send requests to the server
	fileContent1, err := os.ReadFile("cmd/client-tcb/text1.txt")
	if err != nil {
		log.Printf("current working directory: %s", currentDir)
		log.Fatalf("error reading the file: %v", err)
		return
	}

	var files []*pb.File

	file1 := pb.File{
		Name: "texto",
		Type: "txt",
		File: fileContent1,
	}
	files = append(files, &file1)

	fileContent2, err := os.ReadFile("cmd/client-tcb/text2.txt")
	if err != nil {
		log.Printf("current working directory: %s", currentDir)
		log.Fatalf("error reading the file: %v", err)
		return
	}

	file2 := pb.File{
		Name: "text2",
		Type: "txt",
		File: fileContent2,
	}
	files = append(files, &file2)

	req := &pb.TQBRequest{
		Rookie: "Rookie Test",
		Files:  files,
	}

	// Send the request to the server
	if err := stream.Send(req); err != nil {
		log.Fatalf("error while sending %v", err)
	}
	// Wait for 2 seconds before sending the next request
	time.Sleep(2 * time.Second)

	// Close the stream when done sending requests
	stream.CloseSend()
	// Wait for the goroutine to finish receiving messages
	<-waitc
	// Log the end of the bidirectional streaming
	log.Printf("bidirectional Streaming finished")
}
