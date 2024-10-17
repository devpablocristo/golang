// package main

// import (
// 	"fmt"
// 	"log"

// 	"github.com/go-micro/plugins/v4/registry/consul"
// 	"go-micro.dev/v4/registry"
// )

// func main() {
// 	// Dirección de Consul
// 	consulAddress := "localhost:8500"
// 	// Nombre del servicio gRPC registrado en Consul
// 	serviceName := "my-grpc-service"

// 	// Crear un registro de Consul
// 	consulRegistry := consul.NewRegistry(registry.Addrs(consulAddress))

// 	// Obtener las instancias del servicio registrado
// 	services, err := consulRegistry.GetService(serviceName)
// 	if err != nil {
// 		log.Fatalf("Error obteniendo el servicio %s desde Consul: %v", serviceName, err)
// 	}

// 	// Si no hay instancias del servicio
// 	if len(services) == 0 {
// 		log.Fatalf("No se encontraron instancias del servicio %s", serviceName)
// 	}

// 	// Iterar sobre los servicios
// 	for _, service := range services {
// 		fmt.Printf("Servicio: %s\n", service.Name)

// 		// Iterar sobre las instancias (nodos) del servicio
// 		for _, node := range service.Nodes {
// 			fmt.Printf("  - Instancia ID: %s\n", node.Id)
// 			fmt.Printf("  - Dirección: %s\n", node.Address)

// 			// Intentar acceder al puerto si está disponible en los metadatos
// 			if port, ok := node.Metadata["port"]; ok {
// 				fmt.Printf("  - Puerto: %s\n", port)
// 			} else {
// 				fmt.Println("  - Puerto: No disponible en los metadatos")
// 			}
// 		}
// 	}
// }

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-micro/plugins/v4/client/grpc"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/registry"
)

type HelloRequest struct {
	Name string
}

type HelloResponse struct {
	Message string
}

func main() {
	// Crear un registro Consul
	consulRegistry := consul.NewRegistry(registry.Addrs("localhost:8500"))

	// Crear un cliente gRPC usando go-micro
	grpcClient := grpc.NewClient(
		client.Registry(consulRegistry), // Usar Consul para resolver el servicio
	)

	// Crear una nueva solicitud gRPC
	request := grpcClient.NewRequest("example-service", "ExampleService.SayHello", &HelloRequest{
		Name: "Pablo",
	})

	// Crear una respuesta
	var response HelloResponse

	// Crear un contexto
	ctx := context.Background()

	// Hacer la llamada gRPC al servidor
	err := grpcClient.Call(ctx, request, &response)
	if err != nil {
		log.Fatalf("Error calling gRPC service: %v", err)
	}

	// Imprimir la respuesta
	fmt.Printf("Server response: %s\n", response.Message)
}
