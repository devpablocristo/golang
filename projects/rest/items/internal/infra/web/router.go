package web

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	ctrl "github.com/mercadolibre/items/internal/adapters/controller"
)

const port = ":9000"

func NewHTTPServer(ItemCtrl ctrl.ItemController) error {
	r := gin.Default()

	basePath := "/api/v1/inventory"
	publicRouter := r.Group(basePath)

	publicRouter.GET("/Items", ItemCtrl.GetItems)
	publicRouter.POST("/Items", ItemCtrl.AddItem)
	publicRouter.GET("/Items/:id", ItemCtrl.GetItem)

	log.Println("Server listening on port", port)

	return http.ListenAndServe(port, r)
}
