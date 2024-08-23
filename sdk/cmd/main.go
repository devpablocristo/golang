package main

import (
	"fmt"

	loggerpkg "github.com/devpablocristo/golang/sdk/pkg/configurators/logger"
	viperpkg "github.com/devpablocristo/golang/sdk/pkg/configurators/viper"
	pkgmapdb "github.com/devpablocristo/golang/sdk/pkg/databases/in-memory/mapdb"
	pkggomicro "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4"
	pkggin "github.com/devpablocristo/golang/sdk/pkg/rest/gin"

	"github.com/devpablocristo/golang/sdk/internal/core/user"
)

// NOTE: mover examples/go-micro
func main() {
	if err := viperpkg.LoadConfig("../"); err != nil {
		loggerpkg.StdError("Viper Service error: %v", err)
	}

	gomicroService, err := pkggomicro.Bootstrap()
	if err != nil {
		loggerpkg.StdError("GoMicro Service error: %v", err)
	}

	fmt.Println("Starting service with config:", gomicroService)

	//NOTE: gin NO se arranca,
	//NOTE: go-micro webservice si,
	//NOTE: de esta forma gin maneje las solicitudes
	//NOTE: y go-micro el resto
	ginService, err := pkggin.Bootstrap()
	if err != nil {
		loggerpkg.GmError("Gin Service error: %v", err)
	}

	r := ginService.GetRouter()

	mapdbService := pkgmapdb.NewService()
	userRepository := user.NewMapDbRepository(mapdbService)
	user.NewUserUseCases(userRepository)

	go func() {
		if err := gomicroService.StartRcpService(); err != nil {
			loggerpkg.GmError("Error starting GoMicro Service: %v", err)
		}
	}()

	gomicroService.GetWebService().Handle("/", r)

	if err := gomicroService.StartWebService(); err != nil {
		loggerpkg.GmError("Error starting GoMicro Service: %v", err)
	}
}
