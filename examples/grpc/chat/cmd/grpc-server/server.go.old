package main

import (
	"io"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "chat/chatpb"
)

const (
	port = ":50051"
)

type chatServer struct {
	pb.ChatServiceServer
}

func (b *chatServer) Chat(stream pb.ChatService_ChatServer) error {

	for {

		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != io.EOF {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}

		stream.Send(&pb.ChatResponse{
			Content: "Hello " + req.User + "!",
		})
		if err != nil {
			log.Fatalf("Error while sending data to client: %v", err)
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	s := grpc.NewServer()
	pb.RegisterChatServiceServer(s, &chatServer{})
	log.Printf("Server listening at %v", lis.Addr())
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve over port %v: %v", port, err)
	}
}
