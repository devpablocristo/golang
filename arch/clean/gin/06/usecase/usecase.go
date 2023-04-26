package usecase

import (
	"fmt"

	entity "items/entity"
)

// Usecases
type ItemUsecaseInterface interface {
	SaveItem(entity.Item) (entity.Item, error)
	GetAllItems() (entity.MapRepo, error)
}

// el tipo de usecase es del tipo interface de repository
type ItemUsecase struct {
	repo entity.ItemRepository
}

// como parametro de salida se usar la interface de usecase
func NewItemUsecase(repo entity.ItemRepository) ItemUsecaseInterface {
	return &ItemUsecase{
		repo: repo,
	}
}

func (u *ItemUsecase) SaveItem(item entity.Item) (entity.Item, error) {
	if err := u.repo.SaveItem(item); err != nil {
		return entity.Item{}, fmt.Errorf("error saving entity.entity.Item: %w", err)
	}

	return item, nil
}

func (u *ItemUsecase) GetAllItems() (entity.MapRepo, error) {
	items, err := u.repo.GetAllItems()
	if err != nil {
		return items, fmt.Errorf("error in repository: %w", err)
	}

	if len(items) == 0 {
		return items, entity.ErrNotFound
	}

	return items, nil
}
