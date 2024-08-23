package main

import (
	"net/http"

	pkglogger "github.com/devpablocristo/golang/sdk/pkg/configurators/logger"
	pkgviper "github.com/devpablocristo/golang/sdk/pkg/configurators/viper"
	pkgmapdb "github.com/devpablocristo/golang/sdk/pkg/databases/in-memory/mapdb"
	pkggomicro "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4"
	pkggin "github.com/devpablocristo/golang/sdk/pkg/rest/gin"

	"github.com/devpablocristo/golang/sdk/internal/core/user"
)

// NOTE: mover examples/go-micro
func main() {
	if err := pkgviper.LoadConfig(); err != nil {
		pkglogger.StdError("Viper Service error: %v", err)
	}

	gomicroService, err := pkggomicro.Bootstrap()
	if err != nil {
		pkglogger.StdError("GoMicro Service error: %v", err)
	}

	//NOTE: gin NO se arranca,
	//NOTE: go-micro webservice si,
	//NOTE: de esta forma gin maneje las solicitudes
	//NOTE: y go-micro el resto
	ginService, err := pkggin.Bootstrap()
	if err != nil {
		pkglogger.GmError("Gin Service error: %v", err)
	}

	r := ginService.GetRouter()

	mapdbService := pkgmapdb.NewService()
	userRepository := user.NewMapDbRepository(mapdbService)
	user.NewUserUseCases(userRepository)

	pkglogger.GmInfo("ESTO ES INFO")
	go func() {
		if err := gomicroService.StartRcpService(); err != nil {
			pkglogger.GmError("Error starting GoMicro RPC Service: %v", err)
		}
	}()

	gomicroService.GetWebService().Handle("/", r)
	gomicroService.GetWebService().HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	if err := gomicroService.StartWebService(); err != nil {
		pkglogger.GmError("Error starting GoMicro Web Service: %v", err)
	}
}
