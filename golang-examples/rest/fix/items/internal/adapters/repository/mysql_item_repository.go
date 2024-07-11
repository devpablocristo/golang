package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/devpablocristo/golang/06-projects/items/gin/items-final/internal/entity"
	"github.com/jmoiron/sqlx"
)

type itemDAO struct {
	ID          uint
	Code        string
	Description string
	Title       string
	Price       int
	Stock       int
	Status      string
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func (b itemDAO) toItemDomain() entity.Item {
	return entity.Item{
		ID:          b.ID,
		Code:        b.Code,
		Description: b.Description,
		Title:       b.Title,
		Price:       b.Price,
		Stock:       b.Stock,
		Status:      b.Status,
		CreatedAt:   b.CreatedAt,
		UpdatedAt:   b.UpdatedAt,
	}
}

type mysqlItemRepository struct {
	conn *sqlx.DB
}

func NewMySQLItemRepository(db *sqlx.DB) entity.ItemRepository {
	return &mysqlItemRepository{
		conn: db,
	}
}

func (r *mysqlItemRepository) GetItemByID(id uint) (entity.Item, error) {
	var item entity.Item
	var itemDB itemDAO

	err := r.conn.Get(&itemDB, "SELECT * FROM items WHERE id=?", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, entity.ItemNotFound{
				Message: "item not found",
			}
		}
		return item, fmt.Errorf("error getting item: %w", err)
	}

	return itemDB.toItemDomain(), nil
}

func (r *mysqlItemRepository) CheckItemByCode(code string) (bool, error) {
	var exist bool
	err := r.conn.Get(&exist, `SELECT EXISTS(SELECT id FROM items WHERE code =  ?)`, code)
	if err != nil {
		return exist, fmt.Errorf("error getting item: %w", err)
	}

	return exist, nil
}

func (r *mysqlItemRepository) AddItem(item *entity.Item) error {
	createdAt := time.Now()

	result, err := r.conn.Exec(`INSERT INTO items 
		(code, title, author, price, stock, created_at, updated_at) 
		VALUES(?,?,?,?,?,?,?)`, item.Code, item.Title, item.Description, item.Price, item.Stock, createdAt, createdAt)

	if err != nil {
		return fmt.Errorf("error inserting item: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error saving item: %w", err)
	}

	item.ID = uint(id)
	item.CreatedAt = createdAt
	item.UpdatedAt = createdAt

	return nil
}

func (r *mysqlItemRepository) GetItems() ([]entity.Item, error) {
	var items []entity.Item
	var itemsDB []itemDAO

	err := r.conn.Select(&itemsDB, "SELECT * FROM items LIMIT 10")
	if err != nil {
		return items, fmt.Errorf("error getting all items: %w", err)
	}

	for _, b := range itemsDB {
		items = append(items, b.toItemDomain())
	}

	return items, nil
}
