package cin7

import (
	"net/http"

	"github.com/gin-gonic/gin"

	core "github.com/devpablocristo/golang/sdk/internal/core/nimble-cin7"
)

type Handler struct {
	useCase core.Cin7UseCasePort
}

func NewCin7Handler(uc core.Cin7UseCasePort) *Handler {
	return &Handler{useCase: uc}
}

func (h *Handler) ShipmentUpdate(c *gin.Context) {
	var dto ShipmentReq
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shipment := ToCin7Shipment(dto)

	if err := h.useCase.UpdateShipment(shipment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
