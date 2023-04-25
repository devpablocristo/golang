package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	// Se importa la librería Gin
	"github.com/gin-gonic/gin"
)

func main() {

	//  Se crea una instancia de `gin.Engine`
	// `gin.Default()` crea un enrutador con los middleware Logger y Recovery por defecto.
	router := gin.Default()

	r := newRepository()

	u := newitemUsecase(r)

	// creacion instacia del controller
	// es necesario inyectar en newController un repositorio
	h := newController(u)

	// Se definen las rutas
	router.GET("/", h.helloWorld)
	router.POST("/items", h.saveItem)
	router.GET("/items", h.getAllItems)

	log.Println("Server started at http://localhost:80808080/")

	// Se crea el servidor con el método `Run` de Gin:
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

// ////////////////////////////////////////////////////////////////////////////
// error global
// ////////////////////////////////////////////////////////////////////////////
var ErrNotFound = errors.New("not found")

//////////////////////////////////////////////////////////////////////////////
// Controller
//////////////////////////////////////////////////////////////////////////////

type controller struct {
	usecase *itemUsecase
}

// constructor de typo controller, en los parametros de entrada se inyencta el un repository
func newController(u *itemUsecase) *controller {
	return &controller{
		usecase: u, // aqui se carga el repostory inyectoado dentro del controller
	}
}

// como ahora la antigua funcion helloWorld, tiene un reciber de tipo controller,
// es un metodo de controller
func (h *controller) helloWorld(c *gin.Context) {
	c.String(http.StatusOK, "¡Hello World!")
}

func (h *controller) saveItem(c *gin.Context) {
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

func (h *controller) getAllItems(c *gin.Context) {
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
// Repositorio
/////////////////////////////////////////////////////////////////////////////

// entidad item
type item struct {
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

type mapRepo map[int]item

// creacion del tipo Repository
type repository struct {
	items mapRepo
}

// constructor del repositorio
func newRepository() *repository {
	return &repository{
		items: make(mapRepo), // ATENCION, aqui se satisface el campo items de Repository
	}
}

// este metodo sirve para guardar un item en la base de datos
// este metodo, si bien esta implementado, TODAVIA NO SE UTILIZA
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

func newitemUsecase(repo *repository) *itemUsecase {
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
