package main

import (
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

	// creacion instacia del handler
	// es necesario inyectar en newHandler un repositorio
	h := newHandler(r)

	// Se definen las rutas
	router.GET("/", h.helloWorld)

	log.Println("Servidor escuchando en http://localhost:8080/")

	// Se crea el servidor con el método `Run` de Gin:
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

// ATENCION: solo se implemento un repostorio pero todavia no utiliza en ningun lado

// se crea el tipo o type handler
// con un campo de tipo repository
type handler struct {
	repo *repository
}

// constructor de typo handler, en los parametros de entrada se inyencta el un repository
func newHandler(r *repository) *handler {
	return &handler{
		repo: r, // aqui se carga el repostory inyectoado dentro del handler
	}
}

// como ahora la antigua funcion helloWorld, tiene un reciber de tipo handler,
// es un metodo de handler
func (h handler) helloWorld(c *gin.Context) {
	c.String(http.StatusOK, "¡Hello World!")
}

// se añade un repositorio de tipo inmemory

// entidad Item
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
