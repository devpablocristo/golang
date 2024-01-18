package main

import (
	"log"

	// Se importa la librería Gin

	handler "items/internal/adapters/handler"
	repository "items/internal/adapters/repository"
	usecase "items/internal/usecase"
)

func main() {
	r := repository.NewRepository()
	u := usecase.NewItemUsecase(r)
	h := handler.NewHandler(u)

	// se mueven ls rutas a otro archivo
	err := NewHTTPServer(h)
	if err != nil {
		log.Fatalln(err)
	}
}
