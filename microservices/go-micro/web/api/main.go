package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-micro/micro/v4"
	"github.com/go-micro/plugins/v4/client/http"
	"github.com/go-micro/plugins/v4/registry/consul"
)

func main() {
	// Crear un nuevo registro Consul
	registry := consul.NewRegistry()

	// Crear un nuevo servicio cliente
	service := micro.NewService(
		micro.Client(http.NewClient()), // Cliente HTTP
		micro.Registry(registry),       // Usar Consul para el descubrimiento
	)

	// Inicializar el servicio
	service.Init()

	// Crear un router Gin para el API Gateway
	router := gin.Default()
	router.GET("/api/greeter", func(c *gin.Context) {
		name := c.Query("name")

		// Crear una solicitud al servicio de saludo
		req := service.Client().NewRequest("greeter.service", "/greeter/hello?name="+name, http.MethodGet)
		rsp := service.Client().NewResponse()

		// Ejecutar la solicitud
		if err := service.Client().Call(c.Request.Context(), req, rsp); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Leer y mostrar la respuesta
		body, _ := ioutil.ReadAll(rsp.Body())
		c.String(http.StatusOK, string(body))
	})

	// Ejecutar el API Gateway
	if err := router.Run(":8080"); err != nil {
		fmt.Println(err)
	}
}
