package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create an instance of the handler.
	h := newHandler()

	// Create an instance of `gin.Engine`.
	router := gin.Default()

	// Define the route for the handler.
	router.GET("/", h.helloWorld)

	log.Println("Server started at http://localhost:8080/")

	// Start the server using Gin's `Run` method:
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

// Create the type or struct handler.
type handler struct{}

// Constructor for the handler type.
// Input parameters will be used for dependency injection
// to create the necessary service for the handler to function.
func newHandler() *handler {
	return &handler{}
}

// Since the helloWorld function now has a receiver of type handler,
// it becomes a method of the handler.
func (h *handler) helloWorld(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!")
}
