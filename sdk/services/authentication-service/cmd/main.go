package main

import (
	"log"

	sdkviper "github.com/devpablocristo/golang/sdk/pkg/configurators/viper"
	sdkmdb "github.com/devpablocristo/golang/sdk/pkg/databases/in-memory/mapdb"
	psdkmdb "github.com/devpablocristo/golang/sdk/pkg/databases/in-memory/mapdb/ports"
	sdkgm "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4"
	psdkgm "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/ports"
	sdkgin "github.com/devpablocristo/golang/sdk/pkg/rest/gin"
	psdkgin "github.com/devpablocristo/golang/sdk/pkg/rest/gin/ports"
	auth "github.com/devpablocristo/golang/sdk/services/authentication-service/internal/auth/core"
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

	auth.NewUseCases()

	userRepository := coreuser.NewMapDbRepository(mapdbService)
	userUsecases := coreuser.NewUseCases(userRepository)
	userHandler := gtwuser.NewGinHandler(userUsecases, ginServer)
	userHandler.Routes("secret", "v1")

	go func() {
		if err := gomicroService.StartRcpService(); err != nil {
			log.Fatalf("Error starting GoMicro RPC Service: %v", err)
		}
	}()

	gomicroService.GetWebService().Handle("/", ginServer.GetRouter())

	if err := gomicroService.StartWebService(); err != nil {
		log.Fatalf("Error starting GoMicro Web Service: %v", err)
	}
}
