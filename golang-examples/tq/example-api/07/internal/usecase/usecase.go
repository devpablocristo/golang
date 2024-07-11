package usecase

import (
	"fmt"

	"items/internal/domain"
)

// ItemUsecasePort defines the methods that an item use case should implement.
type ItemUsecasePort interface {
	SaveItem(item domain.Item) (domain.Item, error)
	GetAllItems() (domain.MapRepo, error)
}

// ItemUsecase is an implementation of the ItemUsecasePort interface.
type ItemUsecase struct {
	repo domain.ItemRepositoryPort
}

// NewItemUsecase creates a new instance of the item use case.
func NewItemUsecase(repo domain.ItemRepositoryPort) ItemUsecasePort {
	return &ItemUsecase{
		repo: repo,
	}
}

// SaveItem saves an item using the repository and returns the saved item.
func (u *ItemUsecase) SaveItem(item domain.Item) (domain.Item, error) {
	if err := u.repo.SaveItem(item); err != nil {
		return domain.Item{}, fmt.Errorf("error saving item: %w", err)
	}

	return item, nil
}

// GetAllItems retrieves all items using the repository.
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
