package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	pb "github.com/devpablocristo/golang-examples/grpc/bookstore/book"
	"google.golang.org/grpc"
)

const (
	port = ":9111"
)

type BookServer struct {
	pb.UnimplementedBookstoreInventoryServer
}

func (b *BookServer) CreateNewBook(ctx context.Context, nb *pb.NewBook) (*pb.Book, error) {
	log.Printf("Recibed: %v", nb.GetTitle())

	book := pb.Book{
		Title:  nb.GetTitle(),
		Author: nb.GetAuthor(),
		Id:     rand.Int31n(1000),
	}

	return &book, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	s := grpc.NewServer()
	pb.RegisterBookstoreInventoryServer(s, &BookServer{})
	log.Printf("Server listening at %v", lis.Addr())
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve over port %v: %v", port, err)
	}
}
