package nimble

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, handler *Handler) {
	r.GET("/nimble-ping", handler.NimblePing)
	r.POST("/order-shipment", handler.OrderShipment)
}
