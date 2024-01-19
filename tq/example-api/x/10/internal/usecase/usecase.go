package usecase

import (
	"fmt"

	"items/internal/domain"

	ctypes "items/internal/platform/custom-types"
)

const (
	noStock = iota // Represent no stock for an item
	inStock        // Represent in stock for an item
)

const (
	activeStatus   = "ACTIVE"   // Item status when it is active
	inactiveStatus = "INACTIVE" // Item status when it is inactive
)

// ItemUsecasePort defines the interface for item use case operations.
type ItemUsecasePort interface {
	SaveItem(*domain.Item) (*domain.Item, error)
	GetAllItems() (domain.MapRepo, error)
	GetItemByID(domain.ID) (*domain.Item, error)
}

type ItemUsecase struct {
	repository domain.ItemRepositoryPort
}

// NewItemUsecase creates a new item use case with the given repository.
func NewItemUsecase(repo domain.ItemRepositoryPort) ItemUsecasePort {
	return &ItemUsecase{
		repository: repo,
	}
}

// SaveItem saves an item, ensuring unique code and setting its status based on stock.
func (u *ItemUsecase) SaveItem(item *domain.Item) (*domain.Item, error) {
	if existingItem, err := u.repository.GetItemByCode(item.Code); err == nil && existingItem != nil {
		// If an item with the same code already exists, return an error.
		return nil, fmt.Errorf("code must be unique %s", item.Code)
	}

	item.Status = inactiveStatus
	if item.Stock > noStock {
		item.Status = activeStatus
	}

	savedItem, err := u.repository.SaveItem(item)
	if err != nil {
		return nil, fmt.Errorf("error saving domain.Item: %w", err)
	}

	return savedItem, nil
}

// GetAllItems retrieves all items, returning an error if none are found.
func (u *ItemUsecase) GetAllItems() (domain.MapRepo, error) {
	items, err := u.repository.GetAllItems()
	if err != nil {
		return nil, fmt.Errorf("error in repository: %w", err)
	}

	if len(items) == 0 {
		return nil, ctypes.NewCustomError(1, ctypes.ErrItemNotFound)
	}
	return items, nil
}

// GetItemByID retrieves an item by its ID, returning a custom error if not found.
func (u *ItemUsecase) GetItemByID(id domain.ID) (*domain.Item, error) {
	item, err := u.repository.GetItemByID(id)
	if err != nil {
		return nil, ctypes.NewCustomError(1, ctypes.ErrItemNotFound)
	}
	return item, nil
}
