package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// en este ejemplo se usara DE MOMENTO, solo la libreria estandar
func main() {

	// creacion instancia del router
	router := http.NewServeMux()

	// creacion instancia del repositorio
	repo := newRepository()

	// creacion instacia del handler
	h := newHandler(*repo)

	// rutas, en esta caso solo hay una
	router.HandleFunc("/", h.helloWorld)

	// creacion de servidor
	server := &http.Server{
		Addr:    ":8080", // el servidor necesita un puerto
		Handler: router,  // y un router para poder funcionar
	}

	// loguea el inicio del servidor
	log.Println("Servidor escuchando en http://localhost:8080/")

	// iniciar el servidor, y en caso de tener un error lo imprime y termina la ejecucion de programa
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

// se crea el tipo o type handler
type handler struct {
	repo repository
}

// constructor de typo handler, en los parametros de entrada de esta es donde
// se usara inyeccion de dependencias para crear el servicio, en esta caso un handler,
// con todo lo que necesite para funcionar
func newHandler(r repository) *handler {
	return &handler{
		repo: r,
	}
}

// como ahora la antigua funcion helloWorld, tiene un reciber de tipo handler,
// es un metodo de handler
func (h *handler) helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "¡Hello World!")
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
