package main

import (
	"log"
	"net/http"

	gtwuser "github.com/devpablocristo/golang/sdk/cmd/gateways/user"
	coreuser "github.com/devpablocristo/golang/sdk/internal/core/user"
	sdkviper "github.com/devpablocristo/golang/sdk/pkg/configurators/viper"
	sdkmapdb "github.com/devpablocristo/golang/sdk/pkg/databases/in-memory/mapdb"
	sdkgomicro "github.com/devpablocristo/golang/sdk/pkg/microservices/go-micro/v4"
	sdkgin "github.com/devpablocristo/golang/sdk/pkg/rest/gin"
)

// NOTE: mover examples/go-micro
func init() {
	if err := sdkviper.LoadConfig(); err != nil {
		log.Fatalf("Viper Service error: %v", err)
	}
}

func main() {
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

	r := ginServer.GetRouter()

	mapdbService, err := sdkmapdb.Boostrap()
	if err != nil {
		log.Fatalf("MapDB Service error: %v", err)
	}
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

// package main

// import (
// 	"log"
// 	"context"

// 	coreuser "github.com/devpablocristo/golang/sdk/internal/core/user"
// 	gtwuser "github.com/devpablocristo/golang/sdk/cmd/gateways/user"
// 	sdkmapdb "github.com/devpablocristo/golang/sdk/pkg/databases/in-memory/mapdb"
// 	sdkgin "github.com/devpablocristo/golang/sdk/pkg/rest/gin"
// 	pkgrabbitmq "github.com/devpablocristo/golang/sdk/pkg/messaging/rabbitmq"
// 	pkgggrpc "github.com/devpablocristo/golang/sdk/pkg/grpc"
// )

// func main() {
// 	// Adaptadores Driven (Producers, Repositories, Clients)
// 	rabbitProducer, err := pkgrabbitmq.BootstrapProducer()
// 	if err != nil {
// 		log.Fatalf("RabbitMQ Producer error: %v", err)
// 	}

// 	mapdbRepository, err := sdkmapdb.Bootstrap()
// 	if err != nil {
// 		log.Fatalf("MapDB Repository error: %v", err)
// 	}

// 	ggrpcClient, err := auth.NewGrpcClient("user-service:50051") // Dirección del servicio gRPC configurada
// 	if err != nil {
// 		log.Fatalf("gRPC Client error: %v", err)
// 	}

// 	// Adaptadores Driver (Consumers, Servers)
// 	rabbitConsumer, err := pkgrabbitmq.BootstrapConsumer()
// 	if err != nil {
// 		log.Fatalf("RabbitMQ Consumer error: %v", err)
// 	}

// 	ginServer, err := sdkgin.Bootstrap()
// 	if err != nil {
// 		log.Fatalf("Gin Server error: %v", err)
// 	}

// 	ggrpcServer, err := pkgggrpc.BootstrapServer()
// 	if err != nil {
// 		log.Fatalf("gRPC Server error: %v", err)
// 	}

// 	// Creación de Repositorios y Casos de Uso
// 	userRepository := coreuser.NewMapDbRepository(mapdbRepository)
// 	userMessengerProducer := coreuser.NewRabbitMqProducer(rabbitProducer)

// 	// Inyección del cliente gRPC en los casos de uso
// 	userUsecases := coreuser.NewUserUseCases(userRepository, userMessengerProducer, ggrpcClient)

// 	// Creación de Handlers y Configuración de Servidores
// 	userMessengerConsumer := gtwuser.NewRabbitMqConsumer(userUsecases, rabbitConsumer)
// 	userGrpcServer := gtwuser.NewGrpcServer(userUsecases, ggrpcServer)
// 	userRestHandler := gtwuser.NewGinHandler(userUsecases, ginServer, "v1", "your-jwt-secret")

// 	// Inicialización de los adaptadores driver
// 	go func() {
// 		if err := userMessengerConsumer.StartConsuming(); err != nil {
// 			log.Fatalf("Failed to start RabbitMQ Consumer: %v", err)
// 		}
// 	}()

// 	go func() {
// 		if err := userGrpcServer.Start(); err != nil {
// 			log.Fatalf("Failed to start gRPC Server: %v", err)
// 		}
// 	}()

// 	if err := userRestHandler.Start(); err != nil {
// 		log.Fatalf("Failed to start Gin Server: %v", err)
// 	}
// }
