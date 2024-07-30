package application

import (
	"github.com/devpablocristo/golang/06-apps/bookstore/inventory/application/port"
	"github.com/devpablocristo/golang/06-apps/bookstore/inventory/domain"
)

type InventoryService struct {
	storage port.Repository
}

func NewInventoryService(st port.Repository) *InventoryService {
	return &InventoryService{
		storage: st,
	}
}

func (i *InventoryService) SaveBook(book *domain.Book) error {
	err := i.storage.SaveBook(book)
	if err != nil {
		return err
	}
	return nil
}

func (i *InventoryService) GetBook(ISBN string) (*domain.Book, error) {
	book, err := i.storage.GetBook(ISBN)
	if err != nil {
		return &domain.Book{}, err
	}
	return book, nil
}

func (i *InventoryService) GetInventory() (map[string]*domain.BookStock, error) {
	inventory, err := i.storage.GetInventory()
	if err != nil {
		return nil, err
	}
	return inventory, nil
}

func (i *InventoryService) UpdateBook(book *domain.Book) error {
	err := i.storage.UpdateBook(book)
	if err != nil {
		return err
	}
	return nil
}

func (i *InventoryService) DeleteBook(ISBN string) error {
	err := i.storage.DeleteBook(ISBN)
	if err != nil {
		return err
	}
	return nil
}
