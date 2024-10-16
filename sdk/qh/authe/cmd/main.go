package main

import (
	"log"
	"sync"

	sdkviper "github.com/devpablocristo/golang/sdk/pkg/configurators/viper"
	sdkgmgs "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/grpc-service"
	sdkgmws "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/web-server"

	authconn "github.com/devpablocristo/golang/sdk/qh/authe/internal/adapters/connectors"
	authgtw "github.com/devpablocristo/golang/sdk/qh/authe/internal/adapters/gateways"
	authe "github.com/devpablocristo/golang/sdk/qh/authe/internal/core"
)

func init() {
	if err := sdkviper.LoadConfig(); err != nil {
		log.Fatalf("Viper Service error: %v", err)
	}
}

func main() {
	grpcClient, err := authconn.NewGrpcClient()
	if err != nil {
		log.Fatalf("Go Micro gRPC Client error: %v", err)
	}

	redisService, err := authconn.NewRedisService()
	if err != nil {
		log.Fatalf("Redis error: %v", err)
	}

	jwtService, err := authconn.NewJwtService()
	if err != nil {
		log.Fatalf("JWT Service error: %v", err)
	}

	authUsecases := authe.NewUseCases(grpcClient, redisService, jwtService)

	grpcServer, err := authgtw.NewGrpcServer(authUsecases)
	if err != nil {
		log.Fatalf("Go Micro gRPC Server error: %v", err)
	}

	authHandler, err := authgtw.NewGinHandler(authUsecases)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	grpcService, err := sdkgmgs.Bootstrap(grpcServer.GetServer(), grpcClient.GetClient())
	if err != nil {
		log.Fatalf("Go Micro service Boostrap error: %v", err)
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
