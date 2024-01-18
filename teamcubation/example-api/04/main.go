package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := newRepository()
	u := newItemUsecase(r) // It is necessary to inject a repository into newItemUsecase
	h := newHandler(u)     // It is necessary to inject an itemUsecase into newHandler

	router := gin.Default()

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

// /////////////////////////////////////////////////////////////////////////////
// Global error
// /////////////////////////////////////////////////////////////////////////////
var ErrNotFound = errors.New("not found")

//////////////////////////////////////////////////////////////////////////////
// Handler
//////////////////////////////////////////////////////////////////////////////

type handler struct {
	usecase *itemUsecase
}

// Handler type constructor; an itemUsecase is injected into the parameters
func newHandler(u *itemUsecase) *handler {
	return &handler{
		usecase: u, // Here, the injected itemUsecase is loaded into the handler
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
// Usecases
/////////////////////////////////////////////////////////////////////////////

type itemUsecase struct {
	repo *repository
}

// newItemUsecase is a constructor for itemUsecase; a repository is injected into the parameters
func newItemUsecase(repo *repository) *itemUsecase {
	return &itemUsecase{
		repo: repo,
	}
}

// saveItem saves an item and returns it along with any error encountered
func (u *itemUsecase) saveItem(item item) (item, error) {
	if err := u.repo.saveItem(item); err != nil {
		return item, fmt.Errorf("error saving item: %w", err)
	}

	return item, nil
}

// getAllItems retrieves all items from the repository
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

/////////////////////////////////////////////////////////////////////////////
// Repository
/////////////////////////////////////////////////////////////////////////////

type mapRepo map[int]item

// repository is a type representing an in-memory storage for items
type repository struct {
	items mapRepo
}

// newRepository is a constructor for the repository
func newRepository() *repository {
	return &repository{
		items: make(mapRepo), // ATTENTION, here the items field of the repository is satisfied.
	}
}

// saveItem saves an item to the repository
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

// getAllItems retrieves all items from the repository
func (r *repository) getAllItems() (mapRepo, error) {
	return r.items, nil
}

/////////////////////////////////////////////////////////////////////////////
// Domain
/////////////////////////////////////////////////////////////////////////////

// item is an entity representing an item
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
