package main

import (
	"log"

	// Se importa la librería Gin

	handler "github.com/devpablocristo/golang/06-projects/items/gin/7/internal/adapters/handler"
	repository "github.com/devpablocristo/golang/06-projects/items/gin/7/internal/adapters/repository"
	usecase "github.com/devpablocristo/golang/06-projects/items/gin/7/internal/usecase"
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
