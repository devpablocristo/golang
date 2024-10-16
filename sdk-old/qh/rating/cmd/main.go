package main

import (
	"log"
	"sync"

	sdkviper "github.com/devpablocristo/golang/sdk/pkg/configurators/viper"
	sdkgmgs "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/grpc-service"
	sdkgmws "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/web-server"

	ratconn "github.com/devpablocristo/golang/sdk/qh/rating/internal/adapters/connectors"
	ratgtw "github.com/devpablocristo/golang/sdk/qh/rating/internal/adapters/gateways"
	rating "github.com/devpablocristo/golang/sdk/qh/rating/internal/core"
)

func init() {
	if err := sdkviper.LoadConfig(); err != nil {
		log.Fatalf("Viper Service error: %v", err)
	}
}

func main() {
	ratGrpcClient, err := ratconn.NewGrpcClient()
	if err != nil {
		log.Fatalf("Go Micro gRPC Client error: %v", err)
	}

	ratUsecases := rating.NewUseCases(ratGrpcClient)

	ratGrpcServer, err := ratgtw.NewGrpcServer(ratUsecases)
	if err != nil {
		log.Fatalf("Go Micro gRPC Server error: %v", err)
	}

	ratHandler, err := ratgtw.NewGinHandler(ratUsecases)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	ratGrpcService, err := sdkgmgs.Bootstrap(ratGrpcServer.GetServer(), ratGrpcClient.GetClient())
	if err != nil {
		log.Fatalf("Go Micro service Boostrap error: %v", err)
	}

	ratWebServer, err := sdkgmws.Bootstrap(ratHandler.GetRouter())
	if err != nil {
		log.Fatalf("GoMicro Service error: %v", err)
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		if err := ratWebServer.Run(); err != nil {
			log.Fatalf("Error starting Web Server: %v", err)
		}
	}()

	go func() {
		defer wg.Done()
		if err := ratGrpcService.Run(); err != nil {
			log.Fatalf("Error starting gRPC Service: %v", err)
		}
	}()

	wg.Wait()
}
