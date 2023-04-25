package web

import (
	"log"

	gin "github.com/gin-gonic/gin"

	controller "github.com/devpablocristo/golang/06-projects/items/gin/9/internal/adapters/controller"
)

const port = ":8080"

func NewHTTPServer(h *controller.ItemController) error {
	// Se crea una instancia de `gin.Engine`
	// `gin.Default()` crea un enrutador con los middleware Logger y Recovery por defecto.
	router := gin.Default()

	// La función Group de Gin se utiliza para definir un grupo de rutas.
	// Toma una cadena como primer argumento que especifica el prefijo común para todas las rutas en el grupo.
	v1 := router.Group("/v1")

	// Se definen las rutas
	v1.POST("/items", h.SaveItem)
	v1.GET("/items", h.GetAllItems)
	v1.GET("/items/:id", h.GetItemsByID)

	// PUT v1/items/{id}
	// DELETE v1/items/{id}

	log.Println("Server started at http://localhost" + port)

	// Se crea el servidor con el método `Run` de Gin:
	err := router.Run(port)
	if err != nil {
		return err
	}
	return nil
}
