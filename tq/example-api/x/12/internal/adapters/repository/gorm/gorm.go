package repository

import (
	entity "items/internal/entity"

	"gorm.io/gorm"
)

type itemsGormRepository struct {
	DB *gorm.DB
}

// NewDatabase : intializes and returns mysql db
func NewEventRepository(db *gorm.DB) entity.ItemRepositoryPort {
	return &itemsGormRepository{
		DB: db,
	}
}

func (r *itemsGormRepository) GetItemByID(id entity.ID) (*entity.Item, error) {
	var itemDB ItemDAO
	return itemDB.dao2Item(), nil
}

func (r *itemsGormRepository) CheckItemByCode(code string) (bool, error) {
	var exist bool
	return exist, nil
}

func (r *itemsGormRepository) SaveItem(item *entity.Item) (*entity.Item, error) {
	var savedItem *entity.Item
	return savedItem, nil
}

func (r *itemsGormRepository) GetAllItems() (entity.MapRepo, error) {
	items := make(entity.MapRepo)
	return items, nil
}
