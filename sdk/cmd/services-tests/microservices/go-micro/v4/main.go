package main

import (
	"log"

	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
)

func main() {
	// Crear un registro Consul
	consulReg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{"http://consul:8500"} // Dirección de Consul
	})

	// Crear un nuevo servicio Micro con Consul como registro
	ms := micro.NewService(
		micro.Name("minimal-ms"),
		micro.Registry(consulReg),
	)

	// Inicializar el servicio
	ms.Init()

	// Ejecutar el servicio, pero no hace nada útil
	if err := ms.Run(); err != nil {
		log.Fatal(err)
	}
}
