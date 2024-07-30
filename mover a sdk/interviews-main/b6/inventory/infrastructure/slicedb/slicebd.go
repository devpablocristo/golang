package slicedb

import (
	inventory "github.com/devpablocristo/interviews/b6/inventory/domain"
)

type SliceDB struct {
	Inventory []inventory.Book
}

func NewSliceDB() *SliceDB {
	return &SliceDB{}
}

func (s SliceDB) SaveBook(book inventory.Book) error {
	s.Inventory = append(s.Inventory, book)
	return nil
}

func (s SliceDB) GetBook(isbn string) (inventory.Book, error) {
	for _, book := range s.Inventory {
		if book.ISBN == isbn {
			return book, nil
		}
	}
	return inventory.Book{}, nil
}
func (s SliceDB) ListInventory() ([]inventory.Book, error) {
	return s.Inventory, nil
}

func (s SliceDB) DeleteBook(isbn string) error {
	for i, book := range s.Inventory {
		if book.ISBN == isbn {
			s.Inventory = append(s.Inventory[:i], s.Inventory[i+1:]...)
			return nil
		}
	}
	return nil
}
