package repodao

import (
	"time"

	"items/internal/domain"
	entity "items/internal/domain"
)

type ItemDAO struct {
	ID          uint      `db:"id"`
	Code        string    `db:"code"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Price       float64   `db:"price"`
	Stock       int       `db:"stock"`
	Status      string    `db:"status"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func (dao *ItemDAO) DaoToItem() *entity.Item {
	return &entity.Item{
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

func ItemToDao(item *domain.Item) *ItemDAO {
	if item == nil {
		return nil
	}

	return &ItemDAO{
		Code:        item.Code,
		Description: item.Description,
		Title:       item.Title,
		Price:       item.Price,
		Stock:       item.Stock,
		Status:      item.Status,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
	}
}
