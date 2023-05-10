package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	grpchandler "greet/internal/adapters/driver/handlers/grpc"
	cng "greet/internal/config"
	pb "greet/internal/proto"
)

func grpcConn(gh *grpchandler.GreetHandler) error {
	// Create a listener on the specified port
	lis, err := net.Listen("tcp", cng.GrpcServerPort)
	if err != nil {
		customErr := fmt.Errorf("failed to start server %v", err)
		return customErr
	}

	// Create a new gRPC server instance
	grpcServer := grpc.NewServer()

	// Register the greet service with the gRPC server
	pb.RegisterGreetServiceServer(grpcServer, gh)
	log.Printf("server started at %v", lis.Addr())

	// Start the gRPC server on the specified listener
	if err := grpcServer.Serve(lis); err != nil {
		customErr := fmt.Errorf("failed to start: %v", err)
		return customErr
	}

	return nil
}
