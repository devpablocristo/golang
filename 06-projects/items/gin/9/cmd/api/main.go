package main

import (
	"log"

	// Se importa la librería Gin

	controller "github.com/devpablocristo/golang/06-projects/items/gin/9/internal/adapters/controller"
	repository "github.com/devpablocristo/golang/06-projects/items/gin/9/internal/adapters/repository"
	web "github.com/devpablocristo/golang/06-projects/items/gin/9/internal/infra/web"
	usecase "github.com/devpablocristo/golang/06-projects/items/gin/9/internal/usecase"
)

func main() {
	r := repository.NewRepository()
	u := usecase.NewItemUsecase(r)
	h := controller.NewController(u)

	// se mueven ls rutas a otro archivo
	err := web.NewHTTPServer(h)
	if err != nil {
		log.Fatalln(err)
	}
}
