package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	gtwgreeter "github.com/devpablocristo/golang/sdk/cmd/gateways/greeter-server"
	greeter "github.com/devpablocristo/golang/sdk/internal/core/greeter-server"
	sdklogger "github.com/devpablocristo/golang/sdk/pkg/configurators/logger"
	sdkviper "github.com/devpablocristo/golang/sdk/pkg/configurators/viper"
	sdkgrpcserver "github.com/devpablocristo/golang/sdk/pkg/grpc/server"
	sdkgrpcserverport "github.com/devpablocristo/golang/sdk/pkg/grpc/server/ports"
)

func init() {
	if err := sdkviper.LoadConfig(); err != nil {
		log.Fatalf("Viper Service error: %v", err)
	}
}

func main() {
	// Canal para recibir señales del sistema operativo
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Inicializar los servicios
	gServer := setupServices()

	sdklogger.Info("Estandar Info!")
	sdklogger.Mwarning("Este es un warning!")

	// Inicializar y arrancar el servidor gRPC
	greeterUseCases := greeter.NewUseCases()
	greeterGrpcServer := gtwgreeter.NewGrpcServer(greeterUseCases, gServer)

	// Comenzar a escuchar en el servidor gRPC
	if err := greeterGrpcServer.Start(); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}

	// Esperar una señal para finalizar
	<-sigs
	sdklogger.Info("Recibida señal de terminación, cerrando el servidor gRPC...")

	// Aquí podrías añadir la lógica para cerrar conexiones o limpiar recursos si es necesario
}

func setupServices() sdkgrpcserverport.Server {
	s, err := sdkgrpcserver.Bootstrap()
	if err != nil {
		log.Fatalf("failed to initialize gRPC server: %v", err)
	}

	return s
}
