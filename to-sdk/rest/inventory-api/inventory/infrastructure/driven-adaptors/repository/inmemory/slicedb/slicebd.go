package slicedb

import (
	"errors"
	"time"

	"github.com/devpablocristo/golang/06-apps/bookstore/inventory/domain"
)

type SliceDB struct {
	Inventory []domain.BookStock
}

func NewSliceDB() *SliceDB {
	return &SliceDB{}
}

func (s *SliceDB) SaveBook(book *domain.Book) error {
	_, found := s.searchBook(book.ISBN)
	if !found {
		newBookStock := &domain.BookStock{
			Book:      book,
			Stock:     1,
			CreatedAt: time.Now(),
		}
		s.Inventory = append(s.Inventory, *newBookStock)
		return nil
	}
	return errors.New("book found")
}

func (s *SliceDB) GetBook(ISBN string) (*domain.Book, error) {
	invIndex, found := s.searchBook(ISBN)
	if !found {
		return &domain.Book{}, errors.New("book not found")
	}
	return s.Inventory[invIndex].Book, nil
}

func (s *SliceDB) GetInventory() (map[string]*domain.BookStock, error) {
	invMap, err := slice2Map(s.Inventory)
	if err != nil {
		return make(map[string]*domain.BookStock), err
	}
	return invMap, nil
}

func (s *SliceDB) UpdateBook(book *domain.Book) error {
	invIndex, found := s.searchBook(book.ISBN)
	if !found {
		return errors.New("book not found")
	}
	s.Inventory[invIndex].Book = book
	s.Inventory[invIndex].Stock = s.Inventory[invIndex].Stock + 1
	return nil
}

func (s *SliceDB) DeleteBook(ISBN string) error {
	invIndex, found := s.searchBook(ISBN)
	if !found {
		return errors.New("book not found")
	}
	s.Inventory = append(s.Inventory[:invIndex], s.Inventory[invIndex+1:]...)

	return nil
}

func (s *SliceDB) searchBook(ISBN string) (int, bool) {
	var found bool
	for inventoryIndex, book := range s.Inventory {
		if ISBN == book.Book.ISBN {
			found = true
			return inventoryIndex, found
		}
	}
	found = false
	return 0, found
}

func slice2Map(s []domain.BookStock) (map[string]*domain.BookStock, error) {
	invMap := make(map[string]*domain.BookStock)
	for i := 0; i < len(s); i++ {
		invMap[s[i].Book.ISBN] = &s[i]
	}
	return invMap, nil
}
