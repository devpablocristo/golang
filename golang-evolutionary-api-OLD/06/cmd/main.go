package main

import (
	"log"

	gin "github.com/gin-gonic/gin"

	handler "items/handler"
	repository "items/repository"
	usecase "items/usecase"
)

func main() {
	r := repository.NewRepository()
	u := usecase.NewItemUsecase(r)
	h := handler.NewHandler(u)

	router := gin.Default()

	router.GET("/", h.HelloWorld)
	router.POST("/items", h.SaveItem)
	router.GET("/items", h.GetAllItems)

	log.Println("Server started at http://localhost:80808080/")

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
