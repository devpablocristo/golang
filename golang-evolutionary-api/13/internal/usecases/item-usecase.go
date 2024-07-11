package usecase

import (
	"fmt"

	"api/internal/domain"
)

type ItemUsecasePort interface {
	SaveItem(domain.Item) error
	ListItems() (domain.MapRepo, error)
}

type ItemUsecase struct {
	repo domain.ItemRepositoryPort
}

func NewItemUsecase(repo domain.ItemRepositoryPort) ItemUsecasePort {
	return &ItemUsecase{
		repo: repo,
	}
}

func (u *ItemUsecase) SaveItem(it domain.Item) error {
	if err := u.repo.SaveItem(it); err != nil {
		return fmt.Errorf("error saving domain.domain.Item: %w", err)
	}

	return nil
}

func (u *ItemUsecase) ListItems() (domain.MapRepo, error) {
	its, err := u.repo.ListItems()
	if err != nil {
		return nil, fmt.Errorf("error in repository: %w", err)
	}

	if len(its) == 0 {
		return nil, domain.ErrNotFound
	}

	return its, nil
}
