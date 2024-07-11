package main

import (
	"log"

	"items/internal/adapters/handler"
	"items/internal/adapters/repository"
	"items/internal/platform/web"
	"items/internal/usecase"
)

func main() {
	r := repository.NewRepository()
	u := usecase.NewItemUsecase(r)
	h := handler.NewHandler(u)

	err := web.NewHTTPServer(h)
	if err != nil {
		log.Fatalln(err)
	}
}
