package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create an instance of `gin.Engine`.
	router := gin.Default()

	// Create an instance of the repository.
	r := newRepository()

	// Create an instance of the handler.
	// It is necessary to inject a repository into newHandler.
	h := newHandler(r)

	// Define the routes.
	router.GET("/", h.helloWorld)

	log.Println("Server started at http://localhost:8080/")

	// Create the server using Gin's `Run` method:
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

// Create the type or struct handler.
// It has a field of type repository.
type handler struct {
	repo *repository
}

// Constructor for the handler type.
// The repository is injected as a parameter.
func newHandler(r *repository) *handler {
	return &handler{
		repo: r, // Here, the injected repository is loaded into the handler.
	}
}

// As the previous helloWorld function now has a receiver of type handler,
// it becomes a method of the handler.
func (h handler) helloWorld(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!")
}

// Add an in-memory repository.

// Item entity.
type Item struct {
	ID          int
	Code        string
	Title       string
	Description string
	Price       float64
	Stock       int
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type mapRepo map[int]Item

// Create the Repository type.
type repository struct {
	items mapRepo
}

// Constructor for the repository.
func newRepository() *repository {
	return &repository{
		items: make(mapRepo), // ATTENTION, here the items field of Repository is satisfied.
	}
}

// This method is used to save an item in the database.
// Although this method is implemented, it is NOT YET USED.
func (r *repository) saveItem(item Item) error {
	if item.ID == 0 {
		return fmt.Errorf("item ID cannot be 0")
	}
	if _, exists := r.items[item.ID]; exists {
		return fmt.Errorf("item with ID %d already exists", item.ID)
	}
	r.items[item.ID] = item
	return nil
}
