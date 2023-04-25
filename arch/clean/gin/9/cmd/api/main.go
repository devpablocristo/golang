package main

import (
	"log"

	controller "items/internal/adapters/controller"
	repository "items/internal/adapters/repository"
	web "items/internal/infra/web"
	usecase "items/internal/usecase"
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
