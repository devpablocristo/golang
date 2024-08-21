package cin7

import (
	"net/http"

	"github.com/gin-gonic/gin"

	core "github.com/devpablocristo/golang/sdk/internal/core"
)

// Handler representa el controlador para manejar las solicitudes relacionadas con envíos
type Handler struct {
	useCase core.Cin7UseCases
}

// NewCin7Handler crea un nuevo controlador para Cin7
func NewCin7Handler(uc core.Cin7UseCases) *Handler {
	return &Handler{useCase: uc}
}

// ShipmentUpdate maneja la actualización de envíos a través de una solicitud HTTP
func (h *Handler) ShipmentUpdate(c *gin.Context) {
	var dto ShipmentReq

	// Verifica que el JSON recibido se pueda enlazar correctamente al DTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	// Convierte el DTO a la entidad de dominio
	shipment := ToCin7Shipment(dto)

	// Llama al caso de uso para actualizar el envío
	if err := h.useCase.UpdateShipment(shipment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update shipment: " + err.Error()})
		return
	}

	// Responde con un estado de éxito si no hubo errores
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
