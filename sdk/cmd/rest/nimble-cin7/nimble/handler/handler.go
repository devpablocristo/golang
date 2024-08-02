package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	core "github.com/devpablocristo/golang/sdk/internal/core/nimble-cin7"
)

type Handler struct {
	useCase core.NimbleUseCasePort
}

func NewNimbleHandler(uc core.NimbleUseCasePort) *Handler {
	return &Handler{useCase: uc}
}

func (h *Handler) OrderShipment(c *gin.Context) {
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

func (h *Handler) NimblePing(c *gin.Context) {
	c.String(http.StatusOK, "nimble pong")
}
