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

	// Se crea una instancia de `gin.Engine`
	// `gin.Default()` crea un enrutador con los middleware Logger y Recovery por defecto.
	router := gin.Default()

	r := newRepository()

	u := newItemUsecase(r)

	// Creación instancia del controller
	// es necesario inyectar en newController un usecase
	h := newController(u)

	// Se definen las rutas
	router.GET("/", h.helloWorld)
	router.POST("/items", h.saveItemController)
	router.GET("/items", h.getAllItemsController)

	log.Println("Server started at http://localhost:80808080/")

	// Se crea el servidor con el método `Run` de Gin:
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

// Error global
var ErrNotFound = errors.New("not found")

// Controller
type controller struct {
	usecase itemUsecaseInterface
}

// Constructor del tipo controller, en los parametros de entrada se inyecta el un usecase
func newController(u itemUsecaseInterface) *controller {
	return &controller{
		usecase: u, // Aquí se carga el usecase inyectado dentro del controller
	}
}

// La función helloWorld ahora es un método de controller
func (h *controller) helloWorld(c *gin.Context) {
	c.String(http.StatusOK, "¡Hello World!")
}

func (h *controller) saveItemController(c *gin.Context) {
	var item Item
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

func (h *controller) getAllItemsController(c *gin.Context) {
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

// Repositorio
type Item struct {
	ID          int
	Code        string
	Title       string
	Description string
	Price       float32
	Stock       int
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type MapRepo map[int]Item

type Repository struct {
	items MapRepo
}

func newRepository() *Repository {
	return &Repository{
		items: make(MapRepo),
	}
}

func (r *Repository) saveItem(item Item) error {
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

// Usecases
type itemUsecaseInterface interface {
	saveItem(Item) (Item, error)
	getAllItems() (MapRepo, error)
}

type itemUsecase struct {
	repo *Repository
}

func newItemUsecase(repo *Repository) itemUsecaseInterface {
	return &itemUsecase{
		repo: repo,
	}
}

func (u *itemUsecase) saveItem(item Item) (Item, error) {
	if err := u.repo.saveItem(item); err != nil {
		return item, fmt.Errorf("error saving Item: %w", err)
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
