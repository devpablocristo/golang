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

// client new greater
// func main() {
// 	rootDir, err := os.Getwd()
// 	if err != nil {
// 		ctyoes.HandleFatalError("Error getting current working directory: %v", err)
// 	}

// 	envFile := rootDir + "/.env"
// 	err = godotenv.Load(envFile)
// 	if err != nil {
// 		ctyoes.HandleFatalError("error loading .env file", err)
// 	}

// 	// Connect to the gRPC server using insecure credentials
// 	grGrpcConfig := gconfig.NewGrpcConfig()
// 	conn, err := grpc.Dial(grGrpcConfig.GetServerHost()+":"+grGrpcConfig.GetServerPort(), grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		log.Fatalf("failed to connect: %v", err)
// 	}
// 	// Close the connection when the main function returns
// 	defer conn.Close()
