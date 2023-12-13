package grpchandler

import (
	io "io"
	log "log"

	port "github.com/devpablocristo/qh/internal/greeter/ports"
	pb "github.com/devpablocristo/qh/internal/greeter/proto"
)

type GreetHandler struct {
	greetService port.Service
	pb.GreeterServer
}

func NewGreetHandler(gs port.Service) pb.GreeterServer {
	return &GreetHandler{
		greetService: gs,
	}
}

// Implement the SayHello method of the GreeterServer interface
func (s *GreetHandler) SayHello(stream pb.Greeter_SayHelloServer) error {
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
