package item

import (
	"fmt"
)

// MapRepository es una implementación en memoria del repositorio de elementos
type MapRepository struct {
	items MapRepo // Mapa de elementos
}

// NewMapRepository crea una nueva instancia de MapRepository
func NewMapMapRepository() ItemRepositoryPort {
	return &MapRepository{
		items: make(MapRepo),
	}
}

// SaveItem guarda un nuevo elemento en el repositorio
func (r *MapRepository) SaveItem(it *Item) error {
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
func (r *MapRepository) ListItems() (MapRepo, error) {
	return r.items, nil
}
