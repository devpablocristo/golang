package main

import (
	"log"

	gin "github.com/gin-gonic/gin"

	controller "items/internal/adapters/controller"
)

const port = ":8080"

func NewHTTPServer(h *controller.ItemController) error {
	// Se crea una instancia de `gin.Engine`
	// `gin.Default()` crea un enrutador con los middleware Logger y Recovery por defecto.
	router := gin.Default()

	// Se definen las rutas
	router.GET("/", h.HelloWorld)
	router.POST("/items", h.SaveItem)
	router.GET("/items", h.GetAllItems)

	log.Println("Server started at http://localhost" + port)

	// Se crea el servidor con el método `Run` de Gin:
	err := router.Run(port)
	if err != nil {
		return err
	}
	return nil
}
