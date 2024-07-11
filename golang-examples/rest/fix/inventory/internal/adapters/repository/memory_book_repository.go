package repository

import (
	"errors"
	"time"

	"github.com/mercadolibre/inventory/internal/entity"
)

var newID uint = 0

type bookRepository struct {
	db []entity.Book
}

func NewBookRepository() entity.BookRepository {
	return &bookRepository{}
}

func (r *bookRepository) GetBooks() ([]entity.Book, error) {
	return r.db, nil
}

func (r *bookRepository) GetBookByID(id uint) (entity.Book, error) {
	for _, book := range r.db {
		if book.ID == id {
			return book, nil
		}
	}

	return entity.Book{}, errors.New("book not found")
}

func (r *bookRepository) CheckBookByCode(code string) (bool, error) {
	for _, book := range r.db {
		if book.Code == code {
			return true, nil
		}
	}

	return false, nil
}

func (r *bookRepository) AddBook(book *entity.Book) error {
	createdAt := time.Now()
	newID = newID + 1

	book.ID = newID
	book.CreatedAt = createdAt
	book.UpdatedAt = createdAt
	r.db = append(r.db, *book)

	return nil
}
