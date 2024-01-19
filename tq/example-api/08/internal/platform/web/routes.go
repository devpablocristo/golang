package web

import (
	"log"

	gin "github.com/gin-gonic/gin"

	handler "items/internal/adapters/handler"
)

const port = ":8080"

func NewHTTPServer(h *handler.ItemHandler) error {
	router := gin.Default()

	router.GET("/", h.HelloWorld)
	router.POST("/items", h.SaveItem)
	router.GET("/items", h.GetAllItems)

	log.Println("Server started at http://localhost" + port)

	err := router.Run(port)
	if err != nil {
		return err
	}
	return nil
}
