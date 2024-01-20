package handler

import "items/internal/domain"

/*
DTO (Data Transfer Object)
It's a design pattern using structures to transfer essential data between system components, minimizing coupling between them.
*/

// itemDTO is a DTO for transferring essential item data.
type itemDTO struct {
	Code        string  `json:"code"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Status      string  `json:"status"`
}

// dtoToItem converts itemDTO to domain.Item.
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
