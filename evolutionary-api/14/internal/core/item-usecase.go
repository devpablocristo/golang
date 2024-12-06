package core

import (
	"fmt"

	"api/internal/core/item"
	"api/pkg/config"
)

type ItemUsecase struct {
	repo item.ItemRepositoryPort
}

func NewItemUsecase(repo item.ItemRepositoryPort) ItemUsecasePort {
	return &ItemUsecase{
		repo: repo,
	}
}

func (u *ItemUsecase) SaveItem(it item.Item) error {
	if err := u.repo.SaveItem(it); err != nil {
		return fmt.Errorf("error saving item: %w", err)
	}

	return nil
}

func (u *ItemUsecase) ListItems() (item.MapRepo, error) {
	its, err := u.repo.ListItems()
	if err != nil {
		return nil, fmt.Errorf("error in repository: %w", err)
	}

	if len(its) == 0 {
		return nil, config.ErrNotFound
	}

	return its, nil
}
