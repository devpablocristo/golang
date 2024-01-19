package web

import (
	"log"

	gin "github.com/gin-gonic/gin"

	handler "items/internal/adapters/handler"
)

const port = ":8080"

// NewHTTPServer configures and starts an HTTP server using Gin.
// It accepts an item handler and defines routes for item-related operations.
func NewHTTPServer(h *handler.ItemHandler) error {
	router := gin.Default()

	// Defines a route group with the prefix "/v1".
	v1 := router.Group("/v1")

	// Defines routes for item operations.
	v1.POST("/items", h.SaveItem)        // Route to save an item
	v1.GET("/items", h.GetAllItems)      // Route to get all items
	v1.GET("/items/:id", h.GetItemsByID) // Route to get an item by ID

	// Logs the server start.
	log.Println("Server started at http://localhost" + port)

	// Starts the server on the specified port.
	err := router.Run(port)
	if err != nil {
		return err
	}
	return nil
}
