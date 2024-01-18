package web

import (
	"log"
	"net/http"

	ctrl "github.com/devpablocristo/golang/06-projects/items/gin/items-final/internal/adapters/handler"
	"github.com/gin-gonic/gin"
)

const port = ":9000"

func NewHTTPServer(ItemCtrl ctrl.ItemHandler) error {
	r := gin.Default()

	basePath := "/api/v1/inventory"
	publicRouter := r.Group(basePath)

	publicRouter.GET("/Items", ItemCtrl.GetAllItems)
	publicRouter.POST("/Items", ItemCtrl.AddItem)
	publicRouter.GET("/Items/:id", ItemCtrl.GetItem)

	log.Println("Server listening on port", port)

	return http.ListenAndServe(port, r)
}
