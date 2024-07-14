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
	router := gin.Default()

	r := newRepository()
	u := newItemUsecase(r)
	h := newHandler(u)

	// Define routes.
	router.GET("/", h.helloWorld)
	router.POST("/items", h.saveItemHandler)
	router.GET("/items", h.getAllItemsHandler)

	log.Println("Server started at http://localhost:8080/")

	// Create the server using Gin's `Run` method:
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

/////////////////////////////////////////////////////////////////////////////
// Global error
/////////////////////////////////////////////////////////////////////////////

var ErrNotFound = errors.New("not found")

// ///////////////////////////////////////////////////////////////////////////
// Handler
// ///////////////////////////////////////////////////////////////////////////
type handler struct {
	usecase ItemUsecasePort
}

// Constructor for the handler type, injecting a usecase in the input parameters.
func newHandler(u ItemUsecasePort) *handler {
	return &handler{
		usecase: u, // Inject the usecase into the handler here.
	}
}

// The helloWorld function is now a method of the handler.
func (h *handler) helloWorld(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

func (h *handler) saveItemHandler(c *gin.Context) {
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

func (h *handler) getAllItemsHandler(c *gin.Context) {
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

// Repo

type MapRepo map[int]item

type Repository struct {
	items MapRepo
}

func newRepository() *Repository {
	return &Repository{
		items: make(MapRepo),
	}
}

func (r *Repository) saveItem(item item) error {
	if item.ID == 0 {
		return fmt.Errorf("item ID cannot be 0")
	}
	if _, exists := r.items[item.ID]; exists {
		return fmt.Errorf("item with ID %d already exists", item.ID)
	}
	r.items[item.ID] = item
	return nil
}

func (r *Repository) getAllItems() (MapRepo, error) {
	return r.items, nil
}

/////////////////////////////////////////////////////////////////////////////
// Usecases
/////////////////////////////////////////////////////////////////////////////

type ItemUsecasePort interface {
	saveItem(item) (item, error)
	getAllItems() (MapRepo, error)
}

type itemUsecase struct {
	repo *Repository
}

func newItemUsecase(repo *Repository) ItemUsecasePort {
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

func (u *itemUsecase) getAllItems() (MapRepo, error) {
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
// Domain
/////////////////////////////////////////////////////////////////////////////

// item entity.
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
