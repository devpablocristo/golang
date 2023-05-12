package grpchandler

import (
	"io"
	"log"

	pb "greet/internal/proto"
	port "greet/internal/service/ports"
)

// Declare a struct that embeds the GreetServiceServer interface generated from the gRPC proto file
type GreetHandler struct {
	greetService port.GreetService
	pb.GreetServiceServer
}

func NewGreetHandler(gs port.GreetService) *GreetHandler {
	return &GreetHandler{
		greetService: gs,
	}
}

//HttpServer(port, rou)

// Implement the SayHello method of the GreetServiceServer interface
func (s *GreetHandler) SayHello(stream pb.GreetService_SayHelloServer) error {
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
