package item

import (
	"fmt"
)

// Repository es una implementación en memoria del repositorio de elementos
type Repository struct {
	items MapRepo // Mapa de elementos
}

// NewRepository crea una nueva instancia de Repository
func NewRepository() ItemRepositoryPort {
	return &Repository{
		items: make(MapRepo),
	}
}

// SaveItem guarda un nuevo elemento en el repositorio
func (r *Repository) SaveItem(it *Item) error {
	if it.ID == 0 {
		return fmt.Errorf("item ID cannot be 0")
	}
	if _, exists := r.items[it.ID]; exists {
		return fmt.Errorf("item with ID %d already exists", it.ID)
	}
	r.items[it.ID] = *it
	return nil
}

// ListItems lista todos los elementos en el repositorio
func (r *Repository) ListItems() (MapRepo, error) {
	return r.items, nil
}
