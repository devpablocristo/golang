package main

import (
	"Items/internal/adapters/controller"
	"Items/internal/application"
	"Items/internal/infrastructure/repository"
	"Items/internal/infrastructure/web"
	"fmt"
)

func main() {

	repositoryService := repository.NewRepository()
	useCaseService := application.NewUseCaseService(repositoryService)
	controllerService := controller.NewController(useCaseService)
	fmt.Println(controllerService)

	server := web.NewHTTPServer(controllerService)

	fmt.Println("llega?", server)

}
