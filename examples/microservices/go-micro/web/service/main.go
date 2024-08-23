package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-micro/micro/v4"
	"github.com/go-micro/plugins/v4/registry/consul"
	"github.com/go-micro/plugins/v4/server/http"
)

// GreeterService es una estructura que implementa el servicio de saludo
type GreeterService struct{}

// Hello es un método que maneja las solicitudes de saludo
func (g *GreeterService) Hello(ctx context.Context, req *http.Request, rsp *http.ResponseWriter) error {
	name := req.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	response := fmt.Sprintf("Hello, %s!", name)
	(*rsp).Write([]byte(response))
	return nil
}

func main() {
	// Crear un nuevo registro Consul
	registry := consul.NewRegistry()

	// Crear un nuevo servicio HTTP con Go Micro
	service := micro.NewService(
		micro.Server(http.NewServer()), // Servidor HTTP
		micro.Registry(registry),       // Usar Consul para el registro
		micro.Name("greeter.service"),  // Nombre del servicio
		micro.Address(":8081"),         // Dirección del servicio
	)

	// Inicializar el servicio
	service.Init()

	// Crear un router Gin para manejar las rutas
	router := gin.Default()
	router.GET("/greeter/hello", func(c *gin.Context) {
		name := c.Query("name")
		if name == "" {
			name = "World"
		}
		c.String(http.StatusOK, "Hello, %s!", name)
	})

	// Registrar el handler del servicio
	httpServer := service.Server().Options().Server.(*http.Server)
	httpServer.Handler = router

	// Ejecutar el servicio
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
