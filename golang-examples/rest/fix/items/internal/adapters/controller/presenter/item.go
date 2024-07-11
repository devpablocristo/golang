package presenter

import (
	"time"

	"github.com/mercadolibre/items/internal/entity"
)

type jsonItem struct {
	ID          uint      `json:"id"`
	Code        string    `json:"code"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	Stock       int       `json:"stock"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func Item(i entity.Item) jsonItem {
	var itemResponse jsonItem

	itemResponse.ID = i.ID
	itemResponse.Code = i.Code
	itemResponse.Title = i.Title
	itemResponse.Description = i.Description
	itemResponse.Price = i.Price
	itemResponse.Stock = i.Stock
	itemResponse.CreatedAt = i.CreatedAt
	itemResponse.UpdatedAt = i.UpdatedAt

	return itemResponse
}

func Items(items []entity.Item) []jsonItem {
	var itemResponse []jsonItem

	for _, val := range items {
		itemResponse = append(itemResponse, Item(val))
	}

	return itemResponse
}
