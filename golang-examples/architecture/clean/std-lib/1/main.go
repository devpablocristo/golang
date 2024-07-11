package main

import (
	"fmt"
	"log"
	"net/http"
)

// en este ejemplo se usara DE MOMENTO, solo la libreria estandar
func main() {

	// creacion del router
	router := http.NewServeMux()

	// rutas, en esta caso solo hay una
	router.HandleFunc("/", helloWorld)

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

// funcion helloWorld
func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "¡Hello World!")
}
