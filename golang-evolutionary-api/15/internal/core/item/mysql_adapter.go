package item

import (
	"database/sql"
)

type mysqlRepository struct {
	db *sql.DB
}

func NewMySqlRepository(instance *sql.DB) ItemRepositoryPort {
	return &mysqlRepository{
		db: instance,
	}
}

func (r *mysqlRepository) SaveItem(it *Item) error {
	query := `INSERT INTO items (code, title, description, price, stock, status, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.Exec(query, it.Code, it.Title, it.Description, it.Price, it.Stock, it.Status, it.CreatedAt, it.UpdatedAt)
	return err
}

func (r *mysqlRepository) ListItems() (MapRepo, error) {
	query := `SELECT id, code, title, description, price, stock, status, created_at, updated_at FROM items`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make(MapRepo)
	for rows.Next() {
		var it Item
		if err := rows.Scan(&it.ID, &it.Code, &it.Title, &it.Description, &it.Price, &it.Stock, &it.Status, &it.CreatedAt, &it.UpdatedAt); err != nil {
			return nil, err
		}
		items[it.ID] = it
	}

	return items, nil
}
