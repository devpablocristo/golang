package mapdb

import (
	"errors"
	"time"

	"github.com/devpablocristo/golang/06-apps/bookstore/inventory/domain"
)

type MapDB struct {
	Inventory map[string]*domain.BookStock
}

func NewMapDB() *MapDB {
	return &MapDB{}
}

func (m *MapDB) SaveBook(book *domain.Book) error {
	found := m.searchBook(book.ISBN)
	if !found {
		newBookStock := domain.BookStock{
			Book:      book,
			Stock:     1,
			CreatedAt: time.Now(),
		}

		m.Inventory[book.ISBN] = &newBookStock
		return nil
	}

	return errors.New("book found")
}

func (m *MapDB) GetBook(ISBN string) (*domain.Book, error) {
	return m.Inventory[ISBN].Book, nil
}

func (m *MapDB) GetInventory() (map[string]*domain.BookStock, error) {
	return m.Inventory, nil
}

func (m *MapDB) UpdateBook(book *domain.Book) error {
	found := m.searchBook(book.ISBN)
	if !found {
		return errors.New("book not found")
	}

	m.Inventory[book.ISBN].Book = book
	m.Inventory[book.ISBN].Stock = m.Inventory[book.ISBN].Stock + 1
	return nil
}

func (m *MapDB) DeleteBook(ISBN string) error {
	delete(m.Inventory, ISBN)
	return nil
}

func (m *MapDB) searchBook(ISBN string) bool {
	_, found := m.Inventory[ISBN]
	return found
}
