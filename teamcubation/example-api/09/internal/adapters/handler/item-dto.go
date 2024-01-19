package handler

import "items/internal/domain"

/*
DTO (Data Transfer Object)
Es un patrón de diseño con estructuras y una estructura para transferir datos esenciales entre componentes de un sistema y minimiza el acoplamiento entre ellos.
*/

// itemDTO es un DTO para transferir datos esenciales de ítems.
type itemDTO struct {
	Code        string  `json:"code"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Status      string  `json:"status"`
}

// dtoToItem convierte itemDTO a domain.Item.
func dtoToItem(dto *itemDTO) *domain.Item {
	return &domain.Item{
		Code:        dto.Code,
		Title:       dto.Title,
		Description: dto.Description,
		Price:       dto.Price,
		Stock:       dto.Stock,
		Status:      dto.Status,
	}
}
