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

	// creacion instacia del controller
	h := newController()

	// Se definen las rutas
	router.GET("/", h.helloWorld)

	log.Println("Server started at http://localhost:8080/")

	// Se crea el servidor con el método `Run` de Gin:
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

// se crea el tipo o type controller
type controller struct{}

// constructor de typo controller, en los parametros de entrada de esta es donde
// se usara inyeccion de dependencias para crear el servicio, en esta caso un controller,
// con todo lo que necesite para funcionar
func newController( /*paramentros de entrada*/ ) *controller {
	return &controller{}
}

// como ahora la antigua funcion helloWorld, tiene un reciber de tipo controller,
// es un metodo de controller
func (h *controller) helloWorld(c *gin.Context) {
	c.String(http.StatusOK, "¡Hello World!")
}
