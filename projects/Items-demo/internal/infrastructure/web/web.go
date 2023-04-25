package web

import (
	"Items/internal/adapters/controller"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//routes
//implementar gin

func NewHTTPServer(controller controller.Controller) error {
	r := gin.Default()

	basePath := "/api/v1/inventory"
	publicRouter := r.Group(basePath)

	publicRouter.GET("/items", controller.AddItem)

	log.Println("Server listening on port", 8080)

	return http.ListenAndServe(":8080", r)
}
