package main

import (
	"log"

	sdkviper "github.com/devpablocristo/golang/sdk/pkg/configurators/viper"
	sdkgm "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4"

	// sdkgrpclientport "github.com/devpablocristo/golang/sdk/pkg/grpc/client/ports"
	// sdkgrpserverport "github.com/devpablocristo/golang/sdk/pkg/grpc/server/ports"
	// sdkginport "github.com/devpablocristo/golang/sdk/pkg/rest/gin/ports"

	authconn "github.com/devpablocristo/golang/sdk/services/authentication-service/internal/auth/adapters/connectors" // Adaptador de conexión
	authgtw "github.com/devpablocristo/golang/sdk/services/authentication-service/internal/auth/adapters/gateways"    // Adaptador de gateway
	auth "github.com/devpablocristo/golang/sdk/services/authentication-service/internal/auth/core"                    // Casos de uso
)

func init() {
	if err := sdkviper.LoadConfig(); err != nil {
		log.Fatalf("Viper Service error: %v", err)
	}
}

func main() {
	//
	grpcClient := authconn.NewGrpcClient()
	redisService := authconn.NewRedisService()
	jwtService := authconn.NewJwtService()
	//
	authUsecases := auth.NewUseCases(grpcClient, jwtService, redisService)
	//
	grpcServer := authgtw.NewGrpcServer(authUsecases)
	authHandler := authgtw.NewGinHandler(authUsecases)
	//

	gomicroService, err := sdkgm.Bootstrap(grpcClient.GetClient(), grpcServer.GetServer(), authHandler.GetServer())

	if err != nil {
		log.Fatalf("GoMicro Service error: %v", err)
	}

	go func() {
		if err := gomicroService.StartGrpcService(); err != nil {
			log.Fatalf("Error starting GoMicro RPC Service: %v", err)
		}
	}()

	// Configurar el servicio web en GoMicro
	//gomicroService.GetRestServer().Handle("/", authHandler.GetRouter())

	// Iniciar el servicio web
	if err := gomicroService.StartRestServer(); err != nil {
		log.Fatalf("Error starting GoMicro Web Service: %v", err)
	}
}
