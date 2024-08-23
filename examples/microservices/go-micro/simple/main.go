package main

import (
	"net/http"

	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/web"
)

func main() {
	logger.DefaultLogger = logger.NewLogger(logger.WithLevel(logger.DebugLevel))

	// Configura el registro de Consul
	consulReg := consul.NewRegistry(func(opts *registry.Options) {
		opts.Addrs = []string{"localhost:8500"}
	})

	// Crea un micro.Service para RPC
	rpcService := micro.NewService(
		micro.Name("example.rpc.service"),
		micro.Registry(consulReg),
	)

	// Inicializa el servicio RPC
	rpcService.Init()

	// Crea un micro.WebService para HTTP
	webService := web.NewService(
		web.Name("example.web.service"),
		web.Registry(consulReg),
		web.Address(":8088"),
	)

	// Configura un handler HTTP simple
	webService.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Ejecuta ambos servicios
	go func() {
		if err := rpcService.Run(); err != nil {
			logger.Fatal(err)
		}
	}()

	if err := webService.Run(); err != nil {
		logger.Fatal(err)
	}
}
