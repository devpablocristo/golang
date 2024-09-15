package main

import (
	"log"

	sdkviper "github.com/devpablocristo/golang/sdk/pkg/configurators/viper"
	sdkgm "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4"

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
	gomicroService, err := sdkgm.Bootstrap()
	if err != nil {
		log.Fatalf("GoMicro Service error: %v", err)
	}

	grpcClient := authconn.NewGrpcClient()
	redisService := authconn.NewRedisService()
	jwtService := authconn.NewJwtService()
	authUsecases := auth.NewUseCases(grpcClient, jwtService, redisService)
	userHandler := authgtw.NewGinHandler(authUsecases)

	go func() {
		if err := gomicroService.StartRcpService(); err != nil {
			log.Fatalf("Error starting GoMicro RPC Service: %v", err)
		}
	}()

	// Configurar el servicio web en GoMicro
	gomicroService.GetWebService().Handle("/", userHandler.GetRouter())

	// Iniciar el servicio web
	if err := gomicroService.StartWebService(); err != nil {
		log.Fatalf("Error starting GoMicro Web Service: %v", err)
	}
}
