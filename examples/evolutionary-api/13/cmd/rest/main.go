package main

import (
	"log"

	"github.com/gin-gonic/gin"

	handler "api/internal/adapters/handlers"
	repository "api/internal/adapters/repositories"
	usecase "api/internal/usecases"
)

func main() {
	r := repository.NewRepository()
	u := usecase.NewItemUsecase(r)
	h := handler.NewHandler(u)

	router := gin.Default()

	router.POST("/items", h.SaveItem)
	router.GET("/items", h.ListItems)

	log.Println("Server started at http://localhost:8080")

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
