package main

import (
	"log"

	handler "items/internal/adapters/handler"
	repository "items/internal/adapters/repository"
	web "items/internal/infra/web"
	usecase "items/internal/usecase"
)

func main() {
	r := repository.NewRepository()
	u := usecase.NewItemUsecase(r)
	h := handler.NewHandler(u)

	// se mueven ls rutas a otro archivo
	err := web.NewHTTPServer(h)
	if err != nil {
		log.Fatalln(err)
	}
}
