package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	u := newItemUsecase()
	h := newHandler(u) // It is necessary to inject a itemUsecase into newHandler

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
// Usecases
/////////////////////////////////////////////////////////////////////////////

type itemUsecase struct {
}

func newItemUsecase() *itemUsecase {
	return &itemUsecase{}
}

func (u *itemUsecase) saveItem(item item) (item, error) {
	// business logic
	return item, nil
}

func (u *itemUsecase) getAllItems() (map[int]item, error) {
	items := make(map[int]item)
	// business logic
	return items, nil
}

/////////////////////////////////////////////////////////////////////////////
// Domain
/////////////////////////////////////////////////////////////////////////////

// Item entity.
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
