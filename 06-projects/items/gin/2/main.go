package main

import (
	"log"
	"net/http"

	// Se importa la librería Gin
	"github.com/gin-gonic/gin"
)

func main() {

	//  Se crea una instancia de `gin.Engine`
	// `gin.Default()` crea un enrutador con los middleware Logger y Recovery por defecto.
	router := gin.Default()

	// creacion instacia del handler
	h := newHandler()

	// Se definen las rutas
	router.GET("/", h.helloWorld)

	log.Println("Servidor escuchando en http://localhost:8080/")

	// Se crea el servidor con el método `Run` de Gin:
	if err := router.Run(":8080"); err != nil {
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
func (h handler) helloWorld(c *gin.Context) {
	c.String(http.StatusOK, "¡Hello World!")
}
