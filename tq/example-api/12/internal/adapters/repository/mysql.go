package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"

	entity "items/internal/entity"
)

type mysqlItemRepository struct {
	conn *sqlx.DB
}

func NewMySQLRepository(db *sqlx.DB) entity.ItemRepositoryPort {
	return &mysqlItemRepository{
		conn: db,
	}
}

func (r *mysqlItemRepository) GetItemByID(id entity.ID) (*entity.Item, error) {
	var itemDB ItemDAO
	err := r.conn.Get(&itemDB, "SELECT * FROM item WHERE id=?", id)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(errItemNotFound)
			return nil, errItemNotFound
		}
		log.Printf("error getting item: %v", err)
		return nil, fmt.Errorf("error getting item: %v", err)
	}
	return itemDB.dao2Item(), nil
}

func (r *mysqlItemRepository) CheckItemByCode(code string) (bool, error) {
	var exist bool
	err := r.conn.Get(&exist, `SELECT EXISTS(SELECT id FROM items WHERE code =  ?)`, code)
	if err != nil {
		log.Println(err)
		return exist, fmt.Errorf("error getting item: %v", err)
	}
	return exist, nil
}

func (r *mysqlItemRepository) SaveItem(item *entity.Item) (*entity.Item, error) {
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

	savedItem, err := r.GetItemByID(entity.ID(id))
	if err != nil {
		log.Printf("error saving item: %v", err)
		return nil, fmt.Errorf("error saving item: %v", err)
	}

	return savedItem, nil
}

func (r *mysqlItemRepository) GetAllItems() (entity.MapRepo, error) {
	items := make(entity.MapRepo)
	var itemsDB []ItemDAO

	err := r.conn.Select(&itemsDB, "SELECT * FROM item")
	if err != nil {
		return nil, fmt.Errorf("error getting all items: %v", err)
	}

	for _, dao := range itemsDB {
		items[entity.ID(dao.ID)] = dao.dao2Item()
	}

	return items, nil
}
