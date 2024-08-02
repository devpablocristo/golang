package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	core "github.com/devpablocristo/golang-sdk/internal/core/nimble-cin7"
)

type Cin7Handler struct {
	useCase core.Cin7UseCasePort
}

func NewCin7Handler(uc core.Cin7UseCasePort) *Cin7Handler {
	return &Cin7Handler{useCase: uc}
}

func (h *Cin7Handler) ShipmentUpdate(c *gin.Context) {
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
