package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"

	entity "items/internal/entity"
)

type mysqlItemRepository struct {
	conn *sqlx.DB
}

func NewMySQLItemRepository(db *sqlx.DB) entity.ItemRepositoryPort {
	return &mysqlItemRepository{
		conn: db,
	}
}

func (r *mysqlItemRepository) GetItemByID(id entity.ID) (*entity.Item, error) {
	var itemDB itemDAO
	err := r.conn.Get(&itemDB, "SELECT * FROM items WHERE id=?", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errItemNotFound
		}
		return nil, fmt.Errorf("error getting item: %v", err)
	}
	return itemDB.dao2Item(), nil
}

func (r *mysqlItemRepository) GetItemByCode(code string) (*entity.Item, error) {
	var itemDB itemDAO
	err := r.conn.Get(&itemDB, `SELECT EXISTS(SELECT id FROM items WHERE code =  ?)`, code)
	if err != nil {
		return nil, fmt.Errorf("error getting item: %v", err)
	}
	return itemDB.dao2Item(), nil
}

func (r *mysqlItemRepository) SaveItem(item *entity.Item) (*entity.Item, error) {
	createdAt := time.Now()
	updatedAt := createdAt

	result, err := r.conn.Exec(`INSERT INTO items 
		(code, title, author, price, stock, created_at, updated_at) 
		VALUES(?,?,?,?,?,?,?)`, item.Code, item.Title, item.Description, item.Price, item.Stock, createdAt, updatedAt)

	if err != nil {
		return nil, fmt.Errorf("error inserting item: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error saving item: %v", err)
	}

	savedItem, err := r.GetItemByID(entity.ID(id))
	if err != nil {
		return nil, fmt.Errorf("error saving item: %v", err)
	}

	return savedItem, nil
}

func (r *mysqlItemRepository) GetAllItems() (entity.MapRepo, error) {
	items := make(entity.MapRepo)
	var itemsDB []itemDAO

	err := r.conn.Select(&itemsDB, "SELECT * FROM items")
	if err != nil {
		return nil, fmt.Errorf("error getting all items: %v", err)
	}

	for id, dao := range itemsDB {
		items[entity.ID(id)] = dao.dao2Item()
	}

	return items, nil
}
