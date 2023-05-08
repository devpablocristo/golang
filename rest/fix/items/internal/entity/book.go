package entity

import (
	"fmt"
	"time"
)

type Book struct {
	ID        uint
	Code      string
	Author    string
	Title     string
	Price     int
	Stock     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type BookRepository interface {
	GetBooks() ([]Book, error)
	GetBookByID(id uint) (Book, error)
	CheckBookByCode(code string) (bool, error)
	AddBook(book *Book) error
}

type BookNotFound struct {
	Message string
}

func (e BookNotFound) Error() string {
	return fmt.Sprintf("error: '%s'", e.Message)
}

type BookAlreadyExist struct {
	Message string
}

func (e BookAlreadyExist) Error() string {
	return fmt.Sprintf("error: '%s'", e.Message)
}
