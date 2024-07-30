package repository

import (
	inventory "github.com/devpablocristo/interviews/b6/inventory/domain"
)

type RepositoryInteractorRespository interface {
	SaveBook(book inventory.Book) error
	ListInventory() ([]inventory.Book, error)
}

type RepositoryInteractor struct {
	handler RepositoryInteractorRespository
}

func NewRepositoryInteractor(handler RepositoryInteractorRespository) *RepositoryInteractor {
	return &RepositoryInteractor{handler}
}

func (r RepositoryInteractor) SaveBook(book inventory.Book) error {
	return r.handler.SaveBook(book)
}

func (r RepositoryInteractor) ListInventory() ([]inventory.Book, error) {
	results, _ := r.handler.ListInventory()
	return results, nil
}
