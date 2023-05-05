package main

import (
	"context"
	"log"
	"time"

	pb "github.com/devpablocristo/golang-examples/grpc/bookstore/book"
	"google.golang.org/grpc"
)

const (
	port = ":9111"
)

func main() {
	conn, err := grpc.Dial(port, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Connection failed: %v", err)
	}
	defer conn.Close()
	c := pb.NewBookstoreInventoryClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var newBooks = make(map[string]string)

	newBooks["Book A"] = "Author A"
	newBooks["Book B"] = "Author B"

	for title, author := range newBooks {
		r, err := c.CreateNewBook(ctx, &pb.NewBook{Title: title, Author: author})
		if err != nil {
			log.Fatalf("Could not create book: %v", err)
		}
		log.Printf(`Book details:
		Title: %s
		Author: %s
		Id: %d`, r.GetTitle(), r.GetAuthor(), r.GetId())
	}
}
