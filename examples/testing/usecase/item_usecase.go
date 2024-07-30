package usecase

import (
	"fmt"

	"github.com/teamcubation/neocamp-meli/testing/repository"
)

type ItemUsecase interface {
	CreateItem(name string, stock int) error
	GetItemByID(id uint) error
}

type itemUsecase struct {
	repo repository.ItemRepository
}

func NewItemUsecase(repo repository.ItemRepository) ItemUsecase {
	return &itemUsecase{
		repo: repo,
	}
}

func (svc *itemUsecase) CreateItem(name string, stock int) error {
	if name == "" {
		return fmt.Errorf("item name could not be empty")
	}

	if stock == 0 {
		return fmt.Errorf("stock could not be zero")
	}

	err := svc.repo.SaveItem(name, stock)
	if err != nil {
		return fmt.Errorf("error in repository: %w", err)
	}

	return nil
}

func (svc *itemUsecase) GetItemByID(id uint) error {
	if id == 0 {
		return fmt.Errorf("item id cannot be zero")
	}

	err := svc.repo.GetItemByID(id)
	if err != nil {
		return fmt.Errorf("error in repository: %w", err)
	}

	return nil
}
