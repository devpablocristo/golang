package main

import (
	"log"

	gtwuser "github.com/devpablocristo/golang/sdk/cmd/gateways/user"
	coreuser "github.com/devpablocristo/golang/sdk/internal/core/user"
	sdkviper "github.com/devpablocristo/golang/sdk/pkg/configurators/viper"
	sdkmapdb "github.com/devpablocristo/golang/sdk/pkg/databases/in-memory/mapdb"
	portsmdb "github.com/devpablocristo/golang/sdk/pkg/databases/in-memory/mapdb/ports"
	sdkgomicro "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4"
	portsgm "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4/ports"
	sdkgin "github.com/devpablocristo/golang/sdk/pkg/rest/gin"
	portsgin "github.com/devpablocristo/golang/sdk/pkg/rest/gin/ports"
)

func init() {
	if err := sdkviper.LoadConfig(); err != nil {
		log.Fatalf("Viper Service error: %v", err)
	}
}

func main() {

	gomicroService, ginServer, mapdbService := setupServices()

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

func setupServices() (portsgm.Service, portsgin.Server, portsmdb.Repository) {
	gomicroService, err := sdkgomicro.Bootstrap()
	if err != nil {
		log.Fatalf("GoMicro Service error: %v", err)
	}

	//NOTE: gin NO se lanza,
	//NOTE: go-micro webservice si,
	//NOTE: de esta forma gin maneje las solicitudes
	//NOTE: y go-micro el resto
	ginServer, err := sdkgin.Bootstrap()
	if err != nil {
		log.Fatalf("Gin Service error: %v", err)
	}

	mapdbService, err := sdkmapdb.Boostrap()
	if err != nil {
		log.Fatalf("MapDB Service error: %v", err)
	}

	return gomicroService, ginServer, mapdbService

}
