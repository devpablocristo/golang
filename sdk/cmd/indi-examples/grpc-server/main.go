package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/devpablocristo/golang/sdk/pb"
)

// server is used to implement pb.GreeterServer
type server struct {
	pb.UnimplementedGreeterServer
}

// GreetUnary implements pb.GreeterServer
func (s *server) GreetUnary(ctx context.Context, in *pb.GreetUnaryRequest) (*pb.GreetUnaryResponse, error) {
	log.Printf("Received: %v %v", in.Greeting.FirstName, in.Greeting.LastName)
	return &pb.GreetUnaryResponse{Result: "Hello " + in.Greeting.FirstName + " " + in.Greeting.LastName}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Println("Server is running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
