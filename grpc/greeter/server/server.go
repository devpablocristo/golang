/*
This code defines a simple gRPC server that listens for client requests on port 50051.

The server implements the GreetServiceServer interface, which is defined in the pb package that was generated from the gRPC proto file.

The SayHello method receives requests from clients and sends a "Hello" message back to the client in the response.

The server logs the received request and starts serving on the specified port.
*/
package main

import (
	"io"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "greeter/proto"
)

// Define the port the server will listen on
const (
	port = ":50051"
)

// Declare a struct that embeds the GreetServiceServer interface generated from the gRPC proto file
type greetServer struct {
	pb.GreetServiceServer
}

// Implement the SayHello method of the GreetServiceServer interface
func (s *greetServer) SayHello(stream pb.GreetService_SayHelloServer) error {
	// Continuously receive requests from the client
	for {
		// Receive a request from the client
		req, err := stream.Recv()
		// If the client has stopped sending requests, return nil to close the stream
		if err == io.EOF {
			return nil
		}
		// If there's an error receiving the request, return the error
		if err != nil {
			return err
		}
		// Log the received request
		log.Printf("got request with name : %v", req.Name)
		// Create a response to send back to the client
		res := &pb.HelloResponse{
			Message: "Hello " + req.Name,
		}
		// Send the response to the client
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}

// Main function to start the server
func main() {
	// Create a listener on the specified port
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to start server %v", err)
	}

	// Create a new gRPC server instance
	grpcServer := grpc.NewServer()

	// Register the greet service with the gRPC server
	pb.RegisterGreetServiceServer(grpcServer, &greetServer{})
	log.Printf("server started at %v", lis.Addr())

	// Start the gRPC server on the specified listener
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start: %v", err)
	}
}
