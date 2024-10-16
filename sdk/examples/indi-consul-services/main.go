package main

import (
	"fmt"
	"log"

	"github.com/hashicorp/consul/api"
)

func main() {
	// Crear cliente Consul
	consulConfig := api.DefaultConfig()
	consulConfig.Address = "http://localhost:8500"
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		log.Fatalf("Error al conectar con Consul: %v", err)
	}

	// Obtener todos los servicios registrados
	services, _, err := consulClient.Catalog().Services(nil)
	if err != nil {
		log.Fatalf("Error obteniendo la lista de servicios: %v", err)
	}

	// Iterar sobre los servicios y obtener su chequeo de salud si lo tienen
	for serviceName := range services {
		fmt.Printf("Servicio: %s\n", serviceName)

		// Obtener chequeos de salud del servicio
		healthChecks, _, err := consulClient.Health().Checks(serviceName, nil)
		if err != nil {
			log.Printf("Error obteniendo chequeos de salud para %s: %v", serviceName, err)
			continue
		}

		if len(healthChecks) > 0 {
			for _, check := range healthChecks {
				fmt.Printf("  - Chequeo de salud: %s, Estado: %s\n", check.Name, check.Status)
			}
		} else {
			fmt.Println("  - Sin chequeo de salud")
		}
	}
}
