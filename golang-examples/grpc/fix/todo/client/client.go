package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "github.com/devpablocristo/golang-examples/grpc/todo/task"
)

const (
	port = ":9112"
)

func main() {

	conn, err := grpc.Dial(port, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Connection failed: %v", err)
	}
	defer conn.Close()
	todoList := pb.NewTodoListClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var newTasks = make(map[int]string)

	newTasks[1] = "Gym"
	newTasks[2] = "Buy apples"

	for _, t := range newTasks {
		r, err := todoList.CreateNewTask(ctx, &pb.NewTask{Task: t})
		if err != nil {
			log.Fatalf("Could not create task: %v", err)
		}
		log.Printf(`Task details:
		Task: %s
		Id: %d`, r.GetTask(), r.GetId())
	}
}
