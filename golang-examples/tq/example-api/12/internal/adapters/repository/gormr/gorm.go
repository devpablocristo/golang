package gormr

import (
	"errors"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"

	repodao "items/internal/adapters/repository"
	domain "items/internal/domain"
	ctypes "items/internal/platform/custom-types"
)

type mysqlItemRepository struct {
	DB *gorm.DB
}

// NewEventRepository initializes and returns a new GORM repository for items
func NewEventRepository(db *gorm.DB) domain.ItemRepositoryPort {
	return &mysqlItemRepository{
		DB: db,
	}
}

// Implementar GetItemByCode
func (r *mysqlItemRepository) GetItemByCode(code string) (*domain.Item, error) {
	var itemDB repodao.ItemDAO
	result := r.DB.Where("code = ?", code).First(&itemDB)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Println(ctypes.ErrItemNotFound)
			return nil, errors.New(ctypes.ErrItemNotFound)
		}
		log.Printf("error getting item by code: %v", result.Error)
		return nil, fmt.Errorf("error getting item by code: %v", result.Error)
	}
	return itemDB.DaoToItem(), nil
}

// GetItem gets an item by its ID
func (r *mysqlItemRepository) GetItem(id domain.ID) (*domain.Item, error) {
	var itemDB repodao.ItemDAO
	result := r.DB.First(&itemDB, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Println(ctypes.ErrItemNotFound)
			return nil, errors.New(ctypes.ErrItemNotFound)
		}
		log.Printf("error getting item: %v", result.Error)
		return nil, fmt.Errorf("error getting item: %v", result.Error)
	}
	return itemDB.DaoToItem(), nil
}

// SaveItem saves an item to the database
func (r *mysqlItemRepository) SaveItem(item *domain.Item) (*domain.Item, error) {
	itemDB := repodao.ItemToDao(item)
	itemDB.CreatedAt = time.Now()
	itemDB.UpdatedAt = itemDB.CreatedAt

	result := r.DB.Create(&itemDB)
	if result.Error != nil {
		log.Printf("error inserting item: %v", result.Error)
		return nil, fmt.Errorf("error inserting item: %v", result.Error)
	}

	return itemDB.DaoToItem(), nil
}

// GetAllItems retrieves all items
func (r *mysqlItemRepository) GetAllItems() (domain.MapRepo, error) {
	var itemsDB []repodao.ItemDAO
	result := r.DB.Find(&itemsDB)
	if result.Error != nil {
		return nil, fmt.Errorf("error getting all items: %v", result.Error)
	}

	items := make(domain.MapRepo)
	for _, dao := range itemsDB {
		items[domain.ID(dao.ID)] = dao.DaoToItem()
	}

	return items, nil
}
