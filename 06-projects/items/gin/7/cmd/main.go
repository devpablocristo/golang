package main

import (
	"log"

	// Se importa la librería Gin

	controller "github.com/devpablocristo/golang/06-projects/items/gin/7/internal/adapters/controller"
	repository "github.com/devpablocristo/golang/06-projects/items/gin/7/internal/adapters/repository"
	usecase "github.com/devpablocristo/golang/06-projects/items/gin/7/internal/usecase"
)

func main() {
	r := repository.NewRepository()
	u := usecase.NewItemUsecase(r)
	h := controller.NewController(u)

	// se mueven ls rutas a otro archivo
	err := NewHTTPServer(h)
	if err != nil {
		log.Fatalln(err)
	}
}
