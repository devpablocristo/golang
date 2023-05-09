package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	cng "greeter/internal/config"
	pb "greeter/internal/proto"
	service "greeter/internal/service"
)

// Main function to start the server
func main() {

	// a child context is being created to limit the duration a operation,
	// since when calling the cancellation function,
	// any routine waiting for the closed context will immediately return with an error.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// repo()
	// app(repo)
	// handler(app)

	greeterService := service.NewGreetServer()

	// Create a listener on the specified port
	lis, err := net.Listen("tcp", cng.GrpcServerPort)
	if err != nil {
		log.Fatalf("failed to start server %v", err)
	}

	// Create a new gRPC server instance
	grpcServer := grpc.NewServer()

	// Register the greet service with the gRPC server
	pb.RegisterGreetServiceServer(grpcServer, &application.GreetServer{})
	log.Printf("server started at %v", lis.Addr())

	// Start the gRPC server on the specified listener
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start: %v", err)
	}
}
