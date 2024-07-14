package presenter

import (
	"time"

	domain "items/internal/domain"
)

/*
A DTO is used for transferring data without modification,
Presenter adapts and formats data for its final presentation.
The presenter adds an abstraction layer that separates business logic from the presentation layer, proving useful for maintaining a clear separation in the software architecture.
*/

type jsonItem struct {
	Code        string    `json:"code"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func Item(i *domain.Item) *jsonItem {
	var itemResponse jsonItem

	itemResponse.Code = i.Code
	itemResponse.Title = i.Title
	itemResponse.Description = i.Description
	itemResponse.Price = i.Price
	itemResponse.Stock = i.Stock
	itemResponse.Status = i.Status
	itemResponse.CreatedAt = i.CreatedAt
	itemResponse.UpdatedAt = i.UpdatedAt

	return &itemResponse
}

func Items(items domain.MapRepo) map[domain.ID]*jsonItem {
	mJson := make(map[domain.ID]*jsonItem)
	for id, i := range items {
		mJson[id] = Item(i)
	}

	return mJson
}
