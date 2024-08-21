package main

import (
	loggerpkg "github.com/devpablocristo/golang/sdk/pkg/configurators/logger"
	viperpkg "github.com/devpablocristo/golang/sdk/pkg/configurators/viper"
	mapdbpkg "github.com/devpablocristo/golang/sdk/pkg/databases/in-memory/mapdb"
	gomicropkg "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4"
	ginpkg "github.com/devpablocristo/golang/sdk/pkg/rest/gin"

	"github.com/devpablocristo/golang/sdk/internal/core/user"
)

// NOTE: mover examples/go-micro
func main() {
	if err := viperpkg.LoadConfig("../"); err != nil {
		loggerpkg.StdError("Viper Service error: %v", err)
	}

	gomicroService, err := gomicropkg.Bootstrap()
	if err != nil {
		loggerpkg.StdError("GoMicro Service error: %v", err)
	}

	ginService, err := ginpkg.Bootstrap()
	if err != nil {
		loggerpkg.GmError("Gin Service error: %v", err)
	}

	r := ginService.GetRouter()

	mapdbService := mapdbpkg.NewService()
	userRepository := user.NewMapDbRepository(mapdbService)
	user.NewUserUseCases(userRepository)

	//NOTE: gin NO se arranca, pq go-micro maneja esa parte, se pasa de esta forma para que gin maneje las solicitudes
	gomicroService.GetWebService().Handle("/", r)

	if err := gomicroService.Start(); err != nil {
		loggerpkg.GmError("Error starting GoMicro Service: %v", err)
	}
}
