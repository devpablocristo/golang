package main

import (
	"log"

	"items/internal/adapters/handler"
	"items/internal/adapters/repository"
	"items/internal/platform/web"
	"items/internal/usecase"
)

func main() {
	itemRepo := repository.NewRepository()
	itemUsecase := usecase.NewItemUsecase(itemRepo)
	itemHandler := handler.NewHandler(itemUsecase)

	err := web.NewHTTPServer(itemHandler)
	if err != nil {
		log.Fatalf("Failed to start the HTTP server: %v", err)
	}
}
