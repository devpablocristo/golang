package main

import (
	"log"

	usergtw "github.com/devpablocristo/golang/sdk/cmd/gateways/user"
	pb "github.com/devpablocristo/golang/sdk/cmd/gateways/user/pb"

	user "github.com/devpablocristo/golang/sdk/internal/core/user"

	ggrpcsetup "github.com/devpablocristo/golang/sdk/internal/bootstrap/google-grpc"

	mapdbsetup "github.com/devpablocristo/golang/sdk/internal/bootstrap/mapdb"
)

func main() {

	mapDbInst, err := mapdbsetup.NewMapDbInstance()
	if err != nil {
		log.Fatalf("Failed to initialize MapDb: %v", err)
	}
	mapDbRepo := user.NewMapDbRepository(mapDbInst)

	userService := user.NewUserUseCases(mapDbRepo)
	userGrpc := usergtw.NewGgrpcServer(userService)

	grpcInst, err := ggrpcsetup.NewGgrpcServerInstance()
	if err != nil {
		log.Fatalf("Failed to get gRPC grpcInst instance: %v", err)
	}
	grpcInst.RegisterService(&pb.UserService_ServiceDesc, userGrpc)

	log.Println("gRPC grpcInst is running on port 50051")
	if err := grpcInst.Start(); err != nil {
		log.Fatalf("Failed to start gRPC grpcInst: %v", err)
	}
}
