package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	core "github.com/devpablocristo/qh/events/internal/core/integration-nimble-cin7"
)

type NimbleHandler struct {
	useCase core.NimbleUseCasePort
}

func NewNimbleHandler(uc core.NimbleUseCasePort) *NimbleHandler {
	return &NimbleHandler{useCase: uc}
}

func (h *NimbleHandler) OrderShipment(c *gin.Context) {
	var dto OrderReq
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order := ToNimbleOrder(dto)

	if err := h.useCase.ProcessOrder(order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
