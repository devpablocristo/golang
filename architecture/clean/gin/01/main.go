/*
En este código se usa `gin.Default()` para crear un enrutador con los middleware Logger
y Recovery por defecto. En la función `helloWorld`,
se usa `c.String` en lugar de `fmt.Fprintf` para enviar la respuesta al cliente.
Además, se usa el contexto `c` de Gin en lugar de los parámetros `w` y `r` de la función
`http.HandleFunc`. Por último, se utiliza `log` para mostrar información en la consola.
Por ejemplo, se usa `log.Println` para imprimir un mensaje de inicio del servidor.
Además, Gin tiene un middleware Logger que registra cada solicitud HTTP y su resultado en el
registro del servidor.
*/

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

	// Se definen las rutas
	router.GET("/qwerty", helloWorld)

	// esto es simplemente un texto que sale por la salida estandar avisando que el servidor esta online
	log.Println("Server started at http://localhost:8080/")

	// Se crea el servidor con el método `Run` de Gin:
	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

// funcion hello world
func helloWorld(c *gin.Context) {
	c.String(http.StatusOK, "¡Hello World!")
}
