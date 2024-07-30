package usecases

import (
	inventory "github.com/devpablocristo/interviews/b6/inventory/domain"
)

type UseCasesInteractor struct {
	handler inventory.InventoryRepository
}

func NewUseCasesInteractor(handler inventory.InventoryRepository) *UseCasesInteractor {
	return &UseCasesInteractor{handler}
}

func MakeUseCasesInteractor(handler inventory.InventoryRepository) UseCasesInteractor {
	return UseCasesInteractor{handler}
}

func (u UseCasesInteractor) SaveBook(book inventory.Book) error {
	return nil
}

func (u UseCasesInteractor) ListInventory() ([]inventory.Book, error) {
	results, _ := u.handler.ListInventory()
	return results, nil
}
