package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	core "github.com/devpablocristo/qh/events/internal/core/integration-nimble-c7"
)

type NimbleHandler struct {
	useCase core.NimbleUseCasePort
}

func NewNimbleHandler(uc core.NimbleUseCasePort) *NimbleHandler {
	return &NimbleHandler{useCase: uc}
}

func (h *NimbleHandler) HandleOrderShipment(c *gin.Context) {
	var orderDTO order
	if err := c.ShouldBindJSON(&orderDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order := order{
		OrderID:      orderDTO.OrderID,
		CustomerName: orderDTO.CustomerName,
		Items:        toItems(orderDTO.Items),
	}

	if err := h.useCase.ProcessOrder(order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
