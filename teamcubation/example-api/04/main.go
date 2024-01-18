package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	// Import the Gin library
	"github.com/gin-gonic/gin"
)

func main() {

	// Create an instance of `gin.Engine`
	// `gin.Default()` creates a router with the default Logger and Recovery middleware.
	router := gin.Default()

	r := newRepository()

	u := newItemUsecase(r)

	// Create an instance of the handler
	// It is necessary to inject a repository into newHandler
	h := newHandler(u)

	// Define the routes
	router.GET("/", h.helloWorld)
	router.POST("/items", h.saveItem)
	router.GET("/items", h.getAllItems)

	log.Println("Server started at http://localhost:8080/")

	// Create the server using Gin's `Run` method
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

// ////////////////////////////////////////////////////////////////////////////
// Global error
// ////////////////////////////////////////////////////////////////////////////
var ErrNotFound = errors.New("not found")

//////////////////////////////////////////////////////////////////////////////
// Handler
//////////////////////////////////////////////////////////////////////////////

type handler struct {
	usecase *itemUsecase
}

// Handler type constructor; a repository is injected into the parameters
func newHandler(u *itemUsecase) *handler {
	return &handler{
		usecase: u, // Here, the injected repository is loaded into the handler
	}
}

// As the previous helloWorld function now has a receiver of type handler,
// it becomes a method of the handler
func (h *handler) helloWorld(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

func (h *handler) saveItem(c *gin.Context) {
	var item item
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savedItem, err := h.usecase.saveItem(item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, savedItem)
}

func (h *handler) getAllItems(c *gin.Context) {
	items, err := h.usecase.getAllItems()
	if err != nil {
		if err == ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, items)
}

/////////////////////////////////////////////////////////////////////////////
// Repository
/////////////////////////////////////////////////////////////////////////////

// Item entity
type item struct {
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

type mapRepo map[int]item

// Repository type creation
type repository struct {
	items mapRepo
}

// Repository constructor
func newRepository() *repository {
	return &repository{
		items: make(mapRepo), // ATTENTION, here the items field of Repository is satisfied.
	}
}

// This method is used to save an item in the database.
// Although this method is implemented, it is NOT YET USED.
func (r *repository) saveItem(item item) error {
	if item.ID == 0 {
		return fmt.Errorf("item ID cannot be 0")
	}
	if _, exists := r.items[item.ID]; exists {
		return fmt.Errorf("item with ID %d already exists", item.ID)
	}
	r.items[item.ID] = item
	return nil
}

func (r *repository) getAllItems() (mapRepo, error) {
	return r.items, nil
}

/////////////////////////////////////////////////////////////////////////////
// Usecases
/////////////////////////////////////////////////////////////////////////////

type itemUsecase struct {
	repo *repository
}

func newItemUsecase(repo *repository) *itemUsecase {
	return &itemUsecase{
		repo: repo,
	}
}

func (u *itemUsecase) saveItem(item item) (item, error) {
	if err := u.repo.saveItem(item); err != nil {
		return item, fmt.Errorf("error saving item: %w", err)
	}

	return item, nil
}

func (u *itemUsecase) getAllItems() (mapRepo, error) {
	items, err := u.repo.getAllItems()
	if err != nil {
		return items, fmt.Errorf("error in repository: %w", err)
	}

	if len(items) == 0 {
		return items, ErrNotFound
	}

	return items, nil
}
