package main

import (
	"log"

	// Se importa la librería Gin

	controller "items/internal/adapters/controller"
	repository "items/internal/adapters/repository"
	usecase "items/internal/usecase"
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
