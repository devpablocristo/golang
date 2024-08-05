package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

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

	// Crear un cliente HTTP para el servicio de saludo
	req := service.Client().NewRequest("greeter.service", "/greeter/hello?name=Micro", http.MethodGet)

	// Ejecutar la solicitud
	rsp := service.Client().NewResponse()
	err := service.Client().Call(context.Background(), req, rsp)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Leer y mostrar la respuesta
	body, _ := ioutil.ReadAll(rsp.Body())
	fmt.Println(string(body))
}
