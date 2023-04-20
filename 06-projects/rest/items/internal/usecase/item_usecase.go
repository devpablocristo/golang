package usecase

import (
	"fmt"

	"github.com/mercadolibre/items/internal/entity"
)

type ItemUsecase interface {
	GetAllItems() ([]entity.Item, error)
	GetItemByID(id int) (entity.Item, error)
	AddItem(item entity.Item) (entity.Item, error)
}

type itemUsecase struct {
	repo entity.ItemRepository
}

func NewItemUsecase(repo entity.ItemRepository) ItemUsecase {
	return &itemUsecase{
		repo: repo,
	}
}

func (u *itemUsecase) GetAllItems() ([]entity.Item, error) {
	items, err := u.repo.GetItems()
	if err != nil {
		return items, fmt.Errorf("error in repository: %w", err)
	}

	return items, nil
}

func (u *itemUsecase) GetItemByID(id int) (entity.Item, error) {
	item, err := u.repo.GetItemByID(uint(id))
	if err != nil {
		return entity.Item{}, fmt.Errorf("error in repository: %w", err)
	}

	return item, nil
}

func (u *itemUsecase) AddItem(item entity.Item) (entity.Item, error) {
	exist, err := u.repo.CheckItemByCode(item.Code)
	if err != nil {
		return item, fmt.Errorf("error in repository: %w", err)
	}

	if exist {
		return item, entity.ItemAlreadyExist{
			Message: "item already exist",
		}
	}

	if err := u.repo.AddItem(&item); err != nil {
		return entity.Item{}, fmt.Errorf("error in repository: %w", err)
	}

	return item, nil
}
