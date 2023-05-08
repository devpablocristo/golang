package main

import (
	"io"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/devpablocristo/grpc/greeter/proto"
)

const (
	port = ":50051"
)

type greetServer struct {
	pb.GreetServiceServer
}

func (s *greetServer) GreetBi(stream pb.GreetService_SayHelloServer) error {
	for {
		req, err := stream.Recv()
		if err != io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("got request with name: %s", req.Name)
		res := &pb.HelloResponse{
			Message: "Hello " + req.Name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to start listener: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &greetServer{})
	log.Printf("server started at %v", port)

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
