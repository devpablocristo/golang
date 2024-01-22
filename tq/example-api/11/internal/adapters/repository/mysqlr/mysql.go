package mysqlr

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"

	repodao "items/internal/adapters/repository"
	domain "items/internal/domain"
	ctypes "items/internal/platform/custom-types"
)

type mysqlItemRepository struct {
	conn *sqlx.DB
}

func NewMySQLRepository(db *sqlx.DB) domain.ItemRepositoryPort {
	return &mysqlItemRepository{
		conn: db,
	}
}

// Implementar GetItemByCode
func (r *mysqlItemRepository) GetItemByCode(code string) (*domain.Item, error) {
	var itemDB repodao.ItemDAO
	err := r.conn.Get(&itemDB, "SELECT * FROM item WHERE code=?", code)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(ctypes.ErrItemNotFound)
			return nil, errors.New(ctypes.ErrItemNotFound)
		}
		log.Printf("error getting item by code: %v", err)
		return nil, fmt.Errorf("error getting item by code: %v", err)
	}
	return itemDB.DaoToItem(), nil
}

// Los demás métodos se mantienen sin cambios

func (r *mysqlItemRepository) GetItemByID(id domain.ID) (*domain.Item, error) {
	var itemDB repodao.ItemDAO
	err := r.conn.Get(&itemDB, "SELECT * FROM item WHERE id=?", id)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(ctypes.ErrItemNotFound)
			return nil, errors.New(ctypes.ErrItemNotFound)
		}
		log.Printf("error getting item: %v", err)
		return nil, fmt.Errorf("error getting item: %v", err)
	}
	return itemDB.DaoToItem(), nil
}

func (r *mysqlItemRepository) SaveItem(item *domain.Item) (*domain.Item, error) {
	createdAt := time.Now()
	updatedAt := createdAt

	result, err := r.conn.Exec(`INSERT INTO item 
		(code, title, description, price, stock, created_at, updated_at) 
		VALUES(?,?,?,?,?,?,?)`, item.Code, item.Title, item.Description, item.Price, item.Stock, createdAt, updatedAt)

	if err != nil {
		log.Printf("error inserting item: %v", err)
		return nil, fmt.Errorf("error inserting item: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("error saving item: %v", err)
		return nil, fmt.Errorf("error saving item: %v", err)
	}

	savedItem, err := r.GetItemByID(domain.ID(id))
	if err != nil {
		log.Printf("error saving item: %v", err)
		return nil, fmt.Errorf("error saving item: %v", err)
	}

	return savedItem, nil
}

func (r *mysqlItemRepository) GetAllItems() (domain.MapRepo, error) {
	items := make(domain.MapRepo)
	var itemsDB []repodao.ItemDAO

	err := r.conn.Select(&itemsDB, "SELECT * FROM item")
	if err != nil {
		return nil, fmt.Errorf("error getting all items: %v", err)
	}

	for _, dao := range itemsDB {
		items[domain.ID(dao.ID)] = dao.DaoToItem()
	}

	return items, nil
}
