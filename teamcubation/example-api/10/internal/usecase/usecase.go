package usecase

import (
	"fmt"

	entity "items/internal/entity"
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
	SaveItem(*entity.Item) (*entity.Item, error)
	GetAllItems() (entity.MapRepo, error)
	GetItemByID(entity.ID) (*entity.Item, error)
}

// el tipo de usecase es del tipo interface de repository
type ItemUsecase struct {
	repository entity.ItemRepositoryPort
}

// como parametro de salida se usar la interface de usecase
func NewItemUsecase(repo entity.ItemRepositoryPort) ItemUsecasePort {
	return &ItemUsecase{
		repository: repo,
	}
}

func (u *ItemUsecase) SaveItem(item *entity.Item) (*entity.Item, error) {
	_, err := u.repository.GetItemByCode(item.Code)
	if err != nil {
		return nil, fmt.Errorf("codes must be unique %v: %s", err.Error(), item.Code)
	}

	item.Status = inactiveStatus
	if item.Stock > noStock {
		item.Status = activeStatus
	}

	savedItem, err := u.repository.SaveItem(item)
	if err != nil {
		return nil, fmt.Errorf("error saving entity.entity.Item: %w", err)
	}

	return savedItem, nil
}

func (u *ItemUsecase) GetAllItems() (entity.MapRepo, error) {
	items, err := u.repository.GetAllItems()
	if err != nil {
		return nil, fmt.Errorf("error in repository: %w", err)
	}

	if len(items) == 0 {
		return nil, errItemNotFound
	}

	return items, nil
}

func (u *ItemUsecase) GetItemByID(id entity.ID) (*entity.Item, error) {
	item, err := u.repository.GetItemByID(id)
	if err != nil {
		return nil, errItemNotFound
	}
	return item, nil
}
