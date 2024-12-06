package item

import (
	"fmt"
)

// MapRepository es una implementación en memoria del repositorio de elementos
type MapRepository struct {
	items MapRepo // Mapa de elementos
}

// NewMapRepository crea una nueva instancia de MapRepository
func NewMapRepository() ItemRepositoryPort {
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

// UpdateItem actualiza un elemento existente en el repositorio
func (r *MapRepository) UpdateItem(it *Item) error {
	if it.ID == 0 {
		return fmt.Errorf("item ID cannot be 0")
	}
	if _, exists := r.items[it.ID]; !exists {
		return fmt.Errorf("item with ID %d does not exist", it.ID)
	}
	r.items[it.ID] = *it
	return nil
}

// DeleteItem elimina un elemento del repositorio
func (r *MapRepository) DeleteItem(id int) error {
	if id == 0 {
		return fmt.Errorf("item ID cannot be 0")
	}
	if _, exists := r.items[id]; !exists {
		return fmt.Errorf("item with ID %d does not exist", id)
	}
	delete(r.items, id)
	return nil
}
