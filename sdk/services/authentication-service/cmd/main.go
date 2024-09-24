package main

import (
	"log"
	"sync"

	sdkviper "github.com/devpablocristo/golang/sdk/pkg/configurators/viper"
	sdkgmgs "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/grpc-service"
	sdkgmws "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/web-server"

	authconn "github.com/devpablocristo/golang/sdk/services/authentication-service/internal/auth/adapters/connectors"
	authgtw "github.com/devpablocristo/golang/sdk/services/authentication-service/internal/auth/adapters/gateways"
	auth "github.com/devpablocristo/golang/sdk/services/authentication-service/internal/auth/core"
)

func init() {
	if err := sdkviper.LoadConfig(); err != nil {
		log.Fatalf("Viper Service error: %v", err)
	}
}

func main() {

	grpcClient, err := authconn.NewGrpcClient()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	redisService, err := authconn.NewRedisService()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	jwtService := authconn.NewJwtService()

	authUsecases := auth.NewUseCases(grpcClient, jwtService, redisService)

	grpcServer, err := authgtw.NewGrpcServer(authUsecases)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	authHandler, err := authgtw.NewGinHandler(authUsecases)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	grpcService, err := sdkgmgs.Bootstrap(grpcServer.GetServer(), grpcClient.GetClient())
	if err != nil {
		log.Fatalf("GoMicro Service error: %v", err)
	}

	webServer, err := sdkgmws.Bootstrap(authHandler.GetRouter())
	if err != nil {
		log.Fatalf("GoMicro Service error: %v", err)
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		if err := webServer.Run(); err != nil {
			log.Fatalf("Error starting Web Server: %v", err)
		}
	}()

	go func() {
		defer wg.Done()
		if err := grpcService.Run(); err != nil {
			log.Fatalf("Error starting gRPC Service: %v", err)
		}
	}()

	wg.Wait()
}
