package usecase

import (
	"fmt"

	domain "items/domain"
)

type ItemUsecasePort interface {
	SaveItem(domain.Item) (domain.Item, error)
	GetAllItems() (domain.MapRepo, error)
}

type ItemUsecase struct {
	repo domain.ItemRepositoryPort
}

func NewItemUsecase(repo domain.ItemRepositoryPort) ItemUsecasePort {
	return &ItemUsecase{
		repo: repo,
	}
}

func (u *ItemUsecase) SaveItem(item domain.Item) (domain.Item, error) {
	if err := u.repo.SaveItem(item); err != nil {
		return domain.Item{}, fmt.Errorf("error saving domain.domain.Item: %w", err)
	}

	return item, nil
}

func (u *ItemUsecase) GetAllItems() (domain.MapRepo, error) {
	items, err := u.repo.GetAllItems()
	if err != nil {
		return items, fmt.Errorf("error in repository: %w", err)
	}

	if len(items) == 0 {
		return items, domain.ErrNotFound
	}

	return items, nil
}
