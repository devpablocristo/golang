package main

import (
	"fmt"
	"log"
	"net/http"
)

// en este ejemplo se usara DE MOMENTO, solo la libreria estandar
func main() {

	// creacion instancia del router
	router := http.NewServeMux()

	// creacion instacia del handler
	h := newHandler()

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
type handler struct{}

// constructor de typo handler, en los parametros de entrada de esta es donde
// se usara inyeccion de dependencias para crear el servicio, en esta caso un handler,
// con todo lo que necesite para funcionar
func newHandler() *handler {
	return &handler{}
}

// como ahora la antigua funcion helloWorld, tiene un reciber de tipo handler,
// es un metodo de handler
func (h *handler) helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "¡Hello World!")
}
