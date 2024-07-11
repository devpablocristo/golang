package main

import (
	"log"

	handler "items/internal/adapters/handler"
	repository "items/internal/adapters/repository"
	usecase "items/internal/usecase"
)

func main() {
	r := repository.NewRepository()
	u := usecase.NewItemUsecase(r)
	h := handler.NewHandler(u)

	// Routes moved to routes.go
	err := NewHTTPServer(h)
	if err != nil {
		log.Fatalln(err)
	}
}
