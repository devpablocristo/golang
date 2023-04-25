package repository

import (
	"time"

	entity "items/internal/entity"
)

type itemDAO struct {
	ID          uint      `db:"id"`
	Code        string    `db:"code"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Price       float64   `db:"price"`
	Stock       int       `db:"stock"`
	Status      string    `db:"is_active"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func dao2Item(dao itemDAO) entity.Item {
	return entity.Item{
		ID:          dao.ID,
		Code:        dao.Code,
		Description: dao.Description,
		Title:       dao.Title,
		Price:       dao.Price,
		Stock:       dao.Stock,
		Status:      dao.Status,
		CreatedAt:   dao.CreatedAt,
		UpdatedAt:   dao.UpdatedAt,
	}
}
