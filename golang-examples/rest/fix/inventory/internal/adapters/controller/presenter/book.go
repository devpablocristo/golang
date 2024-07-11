package presenter

import (
	"time"

	"github.com/mercadolibre/inventory/internal/entity"
)

type jsonBook struct {
	ID        uint      `json:"id"`
	Code      string    `json:"code"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Price     int       `json:"price"`
	Stock     int       `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func Book(i entity.Book) jsonBook {
	var bookResponse jsonBook

	bookResponse.ID = i.ID
	bookResponse.Code = i.Code
	bookResponse.Title = i.Title
	bookResponse.Author = i.Author
	bookResponse.Price = i.Price
	bookResponse.Stock = i.Stock
	bookResponse.CreatedAt = i.CreatedAt
	bookResponse.UpdatedAt = i.UpdatedAt

	return bookResponse
}

func Books(books []entity.Book) []jsonBook {
	var bookResponse []jsonBook

	for _, val := range books {
		bookResponse = append(bookResponse, Book(val))
	}

	return bookResponse
}
