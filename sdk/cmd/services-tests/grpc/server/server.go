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

	// TODO: probar y crear el generics
	grpcInst, err := ggrpcsetup.NewGgrpcServerInstance()
	if err != nil {
		log.Fatalf("Failed to get gRPC grpcInst instance: %v", err)
	}

	// users := []entities.User{
	// 	{
	// 		UUID: "user1-uuid",
	// 		Credentials: entities.Credentials{
	// 			Username:     "user1",
	// 			PasswordHash: "password1",
	// 		},
	// 	},
	// 	{
	// 		UUID: "user2-uuid",
	// 		Credentials: entities.Credentials{
	// 			Username:     "user2",
	// 			PasswordHash: "password2",
	// 		},
	// 	},
	// }

	// mapDbClient, err := mapdbsetup.NewMapDbInstance(true, users)
	// if err != nil {
	// 	log.Fatalf("Failed to initialize MapDb: %v", err)
	// }

	mapDbClient, err := mapdbsetup.NewMapDbInstance()
	if err != nil {
		log.Fatalf("Failed to initialize MapDb: %v", err)
	}

	// xxx, err := mysqlsetup.NewMySQLInstance()
	// if err != nil {
	// 	log.Fatalf("Failed to initialize MapDb: %v", err)
	// }

	//mapDbClient := user.NewMapDbRepository()
	userService := user.NewUserUseCases(mapDbClient)
	userGrpc := usergtw.NewGgrpcServer(userService)

	grpcInst.RegisterService(&pb.UserService_ServiceDesc, userGrpc)

	log.Println("gRPC grpcInst is running on port 50051")
	if err := grpcInst.Start(); err != nil {
		log.Fatalf("Failed to start gRPC grpcInst: %v", err)
	}
}
