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
	u := newItemUsecase(r)
	h := newHandler(u)

	router := gin.Default()

	router.GET("/", h.home)
	router.GET("/hello", h.hello)
	router.POST("/bye", h.bye)

	router.POST("/items", h.saveItem)
	router.GET("/items", h.listItems)

	log.Println("Server started at http://localhost:8080/")

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

// /////////////////////////////////////////////////////////////////////////////
// Global error
// /////////////////////////////////////////////////////////////////////////////
var errNotFound = errors.New("not found")

// ////////////////////////////////////////////////////////////////////////////
// Handler
// ////////////////////////////////////////////////////////////////////////////
type handler struct {
	usecase *itemUsecase
}

func newHandler(u *itemUsecase) *handler {
	return &handler{
		usecase: u,
	}
}

func (h *handler) home(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to the home page!")
}

func (h *handler) hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello, world!")
}

func (h *handler) bye(c *gin.Context) {
	var msg map[string]string
	if err := c.BindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON"})
		return
	}
	message, exists := msg["message"]
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Message field is missing"})
		return
	}
	c.String(http.StatusOK, "Received POST request with message: %s", message)
}

func (h *handler) saveItem(c *gin.Context) {
	var it item
	err := c.BindJSON(&it)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.usecase.saveItem(it); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "item saved successfully")
}

func (h *handler) listItems(c *gin.Context) {
	items, err := h.usecase.listItems()
	if err != nil {
		if err == errNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, items)
}

// ///////////////////////////////////////////////////////////////////////////
// Usecases
// ///////////////////////////////////////////////////////////////////////////
type itemUsecase struct {
	repo *repository
}

func newItemUsecase(repo *repository) *itemUsecase {
	return &itemUsecase{
		repo: repo,
	}
}

func (u *itemUsecase) saveItem(it item) error {
	if err := u.repo.saveItem(it); err != nil {
		return fmt.Errorf("error saving item: %w", err)
	}
	return nil
}

func (u *itemUsecase) listItems() (mapRepo, error) {
	its, err := u.repo.listItems()
	if err != nil {
		return nil, fmt.Errorf("error in repository: %w", err)
	}
	if len(its) == 0 {
		return nil, errNotFound
	}
	return its, nil
}

// ///////////////////////////////////////////////////////////////////////////
// Repository
// ///////////////////////////////////////////////////////////////////////////
type repository struct {
	items mapRepo
}

func newRepository() *repository {
	return &repository{
		items: make(mapRepo),
	}
}

func (r *repository) saveItem(it item) error {
	if it.ID == 0 {
		return fmt.Errorf("item ID cannot be 0")
	}
	if _, exists := r.items[it.ID]; exists {
		return fmt.Errorf("item with ID %d already exists", it.ID)
	}
	r.items[it.ID] = it
	return nil
}

func (r *repository) listItems() (mapRepo, error) {
	return r.items, nil
}

// ///////////////////////////////////////////////////////////////////////////
// Domain
// ///////////////////////////////////////////////////////////////////////////
// Item entity
type item struct {
	ID          int       `json:"id"`
	Code        string    `json:"code"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type mapRepo map[int]item
