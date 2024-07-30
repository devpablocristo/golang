package main

import (
	"log"

	"github.com/gin-gonic/gin"

	handler "api/cmd/rest/handlers"
	"api/internal/core"
	"api/internal/core/item"
)

func main() {
	r := item.NewRepository()
	u := core.NewItemUsecase(r)
	h := handler.NewHandler(u)

	router := gin.Default()

	router.POST("/items", h.SaveItem)
	router.GET("/items", h.ListItems)

	log.Println("Server started at http://localhost:8080")

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
