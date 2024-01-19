package usecase

import (
	"fmt"

	"items/internal/domain"
	ctypes "items/internal/platform/custom-types"
)

const (
	noStock = iota
	inStock
)

const (
	activeStatus   = "ACTIVE"
	inactiveStatus = "INACTIVE"
)

// Usecases
// Esta es la inferface por donde se comunicara usescases con demas capas
type ItemUsecasePort interface {
	SaveItem(*domain.Item) (*domain.Item, error)
	GetAllItems() (domain.MapRepo, error)
	GetItemByID(domain.ID) (*domain.Item, error)
}

// el tipo de usecase es del tipo interface de repository
type ItemUsecase struct {
	repository domain.ItemRepository
}

// como parametro de salida se usar la interface de usecase
func NewItemUsecase(repo domain.ItemRepository) ItemUsecasePort {
	return &ItemUsecase{
		repository: repo,
	}
}

func (u *ItemUsecase) SaveItem(item *domain.Item) (*domain.Item, error) {
	_, err := u.repository.GetItemByCode(item.Code)
	if err != nil {
		return nil, fmt.Errorf("codes must be unique %s: %s", err.Error(), item.Code)
	}

	item.Status = inactiveStatus
	if item.Stock > noStock {
		item.Status = activeStatus
	}

	savedItem, err := u.repository.SaveItem(item)
	if err != nil {
		return nil, fmt.Errorf("error saving domain.domain.Item: %w", err)
	}

	return savedItem, nil
}

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

func (u *ItemUsecase) GetItemByID(id domain.ID) (*domain.Item, error) {
	item, err := u.repository.GetItemByID(id)
	if err != nil {
		return nil, ctypes.NewCustomError(1, ctypes.ErrItemNotFound)
	}
	return item, nil
}
