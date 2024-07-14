package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
)

func main() {
	// Obtener la dirección de Consul desde la variable de entorno
	consulAddress := os.Getenv("CONSUL_ADDRESS")
	if consulAddress == "" {
		consulAddress = "127.0.0.1:8500"
	}

	// Configurar Consul
	config := api.DefaultConfig()
	config.Address = consulAddress
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Consul client: %v", err)
	}

	// Registrar el servicio en Consul
	registration := &api.AgentServiceRegistration{
		ID:      "consul-example",
		Name:    "app",
		Address: "app", // Usar el nombre del servicio aquí
		Port:    8081,
		Check: &api.AgentServiceCheck{
			HTTP:     "http://app:8081/health", // Usar el nombre del servicio aquí
			Interval: "10s",
			Timeout:  "1s",
		},
	}

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatalf("Failed to register service with Consul: %v", err)
	}

	// Configurar el router Gin
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})

	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Get Users"})
	})

	err = r.Run(":8081")
	if err != nil {
		log.Fatalf("Failed to run Gin server: %v", err)
	}
}
