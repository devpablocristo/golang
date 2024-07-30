package mapdb

import (
	inventory "github.com/devpablocristo/interviews/b6/inventory/domain"
)

type MapDB struct {
	Inventory map[string]inventory.Book
}

func NewMapDB() *MapDB {
	var m MapDB
	m.Inventory = make(map[string]inventory.Book)
	return &m
}

func (m MapDB) SaveBook(book inventory.Book) error {
	m.Inventory[book.ISBN] = book
	return nil
}

func (m MapDB) GetBook(isbn string) (inventory.Book, error) {
	return m.Inventory[isbn], nil
}

func (m MapDB) ListInventory() ([]inventory.Book, error) {
	var results []inventory.Book
	for _, book := range m.Inventory {
		results = append(results, book)
	}
	return results, nil
}

func (m MapDB) DeleteBook(isbn string) error {
	delete(m.Inventory, isbn)
	return nil
}

func (m MapDB) Name() string {
	return "MapDB"
}
