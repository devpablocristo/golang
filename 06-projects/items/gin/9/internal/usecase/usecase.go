package usecase

import (
	"fmt"

	entity "github.com/devpablocristo/golang/06-projects/items/gin/9/internal/entity"
)

// Usecases
type ItemUsecaseInterface interface {
	SaveItem(entity.Item) (entity.Item, error)
	GetItems() (entity.MapRepo, error)
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
	savedItem, err := u.repo.SaveItem(item)
	if err != nil {
		return entity.Item{}, fmt.Errorf("error saving entity.entity.Item: %w", err)
	}

	return savedItem, nil
}

func (u *ItemUsecase) GetItems() (entity.MapRepo, error) {
	items, err := u.repo.GetItems()
	if err != nil {
		return items, fmt.Errorf("error in repository: %w", err)
	}

	if len(items) == 0 {
		return items, errNotFound
	}

	return items, nil
}
