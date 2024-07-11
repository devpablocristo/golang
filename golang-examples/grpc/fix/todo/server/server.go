package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	pb "github.com/devpablocristo/golang-examples/grpc/todo/task"
	"google.golang.org/grpc"
)

const (
	port = ":9113"
)

type TaskServer struct {
	pb.UnimplementedTodoListServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	s := grpc.NewServer()
	pb.RegisterTodoListServer(s, &TaskServer{})
	log.Printf("Server listening at %v", lis.Addr())
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve over port %v: %v", port, err)
	}
}

func (b *TaskServer) CreateNewTask(ctx context.Context, nt *pb.NewTask) (*pb.Task, error) {
	log.Printf("Recibed: %v", nt.GetTask())

	task := pb.Task{
		Task: nt.GetTask(),
		Id:   rand.Int31n(1000),
	}

	return &task, nil
}
