package nimroutes

import (
	"github.com/gin-gonic/gin"

	wire "github.com/devpablocristo/qh/events/cmd/rest"
	is "github.com/devpablocristo/qh/events/pkg/init-setup"
)

func NimRoutes(r *gin.Engine) {
	nimHandler, err := wire.InitializeNimbleHandler()
	if err != nil {
		is.MicroLogError("nimHandler error: %v", err)
	}

	r.GET("/nimble-ping", nimHandler.NimblePing)
	r.POST("/order-shipment", nimHandler.OrderShipment)
}
