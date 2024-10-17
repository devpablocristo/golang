package main

import (
	"context"
	"log"
	"time"

	pb "github.com/devpablocristo/golang/sdk/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Crear un contexto con un tiempo de espera adecuado
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Preparar la solicitud
	request := &pb.GreetUnaryRequest{
		Greeting: &pb.Greeting{
			FirstName: "Emma",
			LastName:  "Watson",
		},
	}

	// client := pb.NewGreeterClient(conn)

	// // Realizar la llamada gRPC
	// response, err := client.GreetUnary(ctx, request)
	// if err != nil {
	// 	log.Fatalf("Failed to call GreetUnary: %v", err)
	// }

	// log.Printf("Response from server: %s", response.Result)

	response2 := &pb.GreetUnaryResponse{}
	method := "/greeter.Greeter/GreetUnary"
	err = conn.Invoke(ctx, method, request, response2)
	if err != nil {
		log.Fatalf("Failed to invoke GreetUnary method: %v", err)
	}

	log.Printf("Response from server: %s", response2.Result)
}
