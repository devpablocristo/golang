package main

import (
	"log"
	"net/http"

	gtwuser "github.com/devpablocristo/golang/sdk/cmd/gateways/user"
	coreuser "github.com/devpablocristo/golang/sdk/internal/core/user"
	pkgviper "github.com/devpablocristo/golang/sdk/pkg/configurators/viper"
	pkgmapdb "github.com/devpablocristo/golang/sdk/pkg/databases/in-memory/mapdb"
	pkggomicro "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4"
	pkggin "github.com/devpablocristo/golang/sdk/pkg/rest/gin"
)

// NOTE: mover examples/go-micro
func init() {
	if err := pkgviper.LoadConfig(); err != nil {
		log.Fatalf("Viper Service error: %v", err)
	}
}

func main() {
	gomicroService, err := pkggomicro.Bootstrap()
	if err != nil {
		log.Fatalf("GoMicro Service error: %v", err)
	}

	//NOTE: gin NO se lanza,
	//NOTE: go-micro webservice si,
	//NOTE: de esta forma gin maneje las solicitudes
	//NOTE: y go-micro el resto
	ginService, err := pkggin.Bootstrap()
	if err != nil {
		log.Fatalf("Gin Service error: %v", err)
	}

	r := ginService.GetRouter()

	mapdbService := pkgmapdb.NewService()
	userRepository := coreuser.NewMapDbRepository(mapdbService)
	userUsecases := coreuser.NewUserUseCases(userRepository)
	userHandler := gtwuser.NewGinHandler(userUsecases)

	gtwuser.Routes(r, userHandler)

	go func() {
		if err := gomicroService.StartRcpService(); err != nil {
			log.Fatalf("Error starting GoMicro RPC Service: %v", err)
		}
	}()

	gomicroService.GetWebService().Handle("/", r)
	gomicroService.GetWebService().HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	if err := gomicroService.StartWebService(); err != nil {
		log.Fatalf("Error starting GoMicro Web Service: %v", err)
	}
}
