package port

import "github.com/devpablocristo/golang/06-apps/bookstore/inventory/domain"

type Service interface {
	SaveBook(*domain.Book) error
	GetBook(string) (*domain.Book, error)
	GetInventory() (map[string]*domain.BookStock, error)
	UpdateBook(*domain.Book) error
	DeleteBook(string) error
}
