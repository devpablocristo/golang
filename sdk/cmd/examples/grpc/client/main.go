package main

import (
	"log"

	pb "github.com/devpablocristo/golang/sdk/cmd/gateways/user/pb"
	setup "github.com/devpablocristo/golang/sdk/internal/bootstrap/google-grpc"
)

func main() {
	client, err := setup.NewClientInstance()
	if err != nil {
		log.Fatalf("Failed to get gRPC server instance: %v", err)
	}

	_ = client 

	pb.NewUserServiceClient(client.)

}
