package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"NimbleCin7Integration/internal/cin7/domain"
	"NimbleCin7Integration/internal/cin7/dto"
)

type Cin7Handler struct {
	useCase domain.Cin7UseCase
}

func NewCin7Handler(uc domain.Cin7UseCase) *Cin7Handler {
	return &Cin7Handler{useCase: uc}
}

func (h *Cin7Handler) HandleShipmentUpdate(c *gin.Context) {
	var shipmentDTO dto.Cin7Shipment
	if err := c.ShouldBindJSON(&shipmentDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shipment := domain.Shipment{
		ShipmentID:  shipmentDTO.ShipmentID,
		OrderID:     shipmentDTO.OrderID,
		ShippedDate: shipmentDTO.ShippedDate,
		Items:       convertDTOItemsToDomainItems(shipmentDTO.Items),
	}

	if err := h.useCase.UpdateShipment(shipment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func convertDTOItemsToDomainItems(items []dto.Item) []domain.Item {
	domainItems := make([]domain.Item, len(items))
	for i, item := range items {
		domainItems[i] = domain.Item{
			ItemID:   item.ItemID,
			Quantity: item.Quantity,
		}
	}
	return domainItems
}
