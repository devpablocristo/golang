package nimble

import (
	"net/http"

	"github.com/devpablocristo/golang/sdk/internal/core"
	"github.com/gin-gonic/gin"
)

// Handler representa el controlador para manejar las solicitudes relacionadas con órdenes
type Handler struct {
	useCase core.NimbleUseCasesPort
}

// NewNimbleHandler crea un nuevo controlador para Nimble
func NewNimbleHandler(uc core.NimbleUseCasesPort) *Handler {
	return &Handler{useCase: uc}
}

// OrderShipment maneja la creación de envíos a través de una solicitud HTTP
func (h *Handler) OrderShipment(c *gin.Context) {
	var dto OrderReq
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	order := ToNimbleOrder(dto)

	if err := h.useCase.ProcessOrder(order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process order: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

// NimblePing es un endpoint de prueba que responde con "nimble pong"
func (h *Handler) NimblePing(c *gin.Context) {
	c.String(http.StatusOK, "nimble pong")
}

