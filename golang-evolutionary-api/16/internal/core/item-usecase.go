package core

import (
	"fmt"

	"api/internal/core/item"
	"api/pkg/config"
)

// ItemUsecase representa el caso de uso para los elementos
type ItemUsecase struct {
	repo item.ItemRepositoryPort // Repositorio de elementos
}

// NewItemUsecase crea una nueva instancia de ItemUsecase
func NewItemUsecase(repo item.ItemRepositoryPort) ItemUsecasePort {
	return &ItemUsecase{
		repo: repo,
	}
}

// SaveItem guarda un nuevo elemento en el repositorio
func (u *ItemUsecase) SaveItem(it item.Item) error {
	if err := u.repo.SaveItem(&it); err != nil {
		return fmt.Errorf("error saving item: %w", err)
	}
	return nil
}

// ListItems lista todos los elementos del repositorio
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

// UpdateItem actualiza un elemento existente en el repositorio
func (u *ItemUsecase) UpdateItem(it item.Item) error {
	if err := u.repo.UpdateItem(&it); err != nil {
		return fmt.Errorf("error updating item: %w", err)
	}
	return nil
}

// DeleteItem elimina un elemento del repositorio
func (u *ItemUsecase) DeleteItem(id int) error {
	if err := u.repo.DeleteItem(id); err != nil {
		return fmt.Errorf("error deleting item: %w", err)
	}
	return nil
}
