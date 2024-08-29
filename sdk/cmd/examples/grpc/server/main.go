package main

import (
	"log"

	usergtw "github.com/devpablocristo/golang/sdk/cmd/gateways/user"
	pb "github.com/devpablocristo/golang/sdk/cmd/gateways/user/pb"

	user "github.com/devpablocristo/golang/sdk/internal/core/user"

	setup "github.com/devpablocristo/golang/sdk/internal/bootstrap/google-grpc"

	mapdbsetup "github.com/devpablocristo/golang/sdk/internal/bootstrap/mapdb"
)

func main() {
	userService := user.NewUserUseCases(mapDbRepo)
	userGrpc := usergtw.NewServer(userService)

	grpcInst, err := setup.NewServerInstance()
	if err != nil {
		log.Fatalf("Failed to get gRPC grpcInst instance: %v", err)
	}
	grpcInst.RegisterService(&pb.UserService_ServiceDesc, userGrpc)

	log.Println("gRPC grpcInst is running on port 50051")
	if err := grpcInst.Start(); err != nil {
		log.Fatalf("Failed to start gRPC grpcInst: %v", err)
	}
}


	// Create a new GreetServiceClient using the connection
	client := pb.NewGreeterClient(conn)

	// Define the names to be sent in the requests
	names := &pb.NamesList{
		Names: []string{"Emma", "Phoebe", "Gordita"},
	}

	// Call the callGreetBi function with the client and names
	callGreetBi(client, names)
}

// callGreetBi function sends requests to the server and receives responses using bidirectional streaming
func callGreetBi(client pb.GreeterClient, names *pb.NamesList) {
	// Log the start of the bidirectional streaming
	log.Printf("Bidirectional Streaming started")

	// Create a new stream with the server
	stream, err := client.SayHello(context.Background())
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

	// Send requests to the server
	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		// Send the request to the server
		if err := stream.Send(req); err != nil {
			log.Fatalf("error while sending %v", err)
		}
		// Wait for 2 seconds before sending the next request
		time.Sleep(2 * time.Second)
	}

	// Close the stream when done sending requests
	stream.CloseSend()
	// Wait for the goroutine to finish receiving messages
	<-waitc
	// Log the end of the bidirectional streaming
	log.Printf("bidirectional Streaming finished")
}

server new greeter grpc

func main() {
	rootDir, err := os.Getwd()
	if err != nil {
		ctypes.HandleFatalError("error getting current working directory: %v", err)
	}

	envFile := rootDir + "/.env"
	err = godotenv.Load(envFile)
	if err != nil {
		ctypes.HandleFatalError("error loading .env file", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	Start(ctx)
}

func Start(ctx context.Context) {
	app := NewAppLauncher()
	app.Setup(ctx)
	var wg sync.WaitGroup

	// Start the REST service in a goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := app.InitEventService(ctx); err != nil {
			ctypes.HandleFatalError("error starting the REST service", err)
		}
	}()

	// Start the gRPC service in a goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := app.InitGreeterServer(ctx); err != nil {
			ctypes.HandleFatalError("error starting gRPC services", err)
		}
	}()

	wg.Wait()
}
