package item

import (
	"database/sql"
)

// mysqlRepository es una implementación del repositorio de elementos utilizando MySQL
type mysqlRepository struct {
	db *sql.DB // Conexión a la base de datos MySQL
}

// NewMySqlRepository crea una nueva instancia de mysqlRepository
func NewMySqlRepository(db *sql.DB) ItemRepositoryPort {
	return &mysqlRepository{
		db: db,
	}
}

// SaveItem guarda un nuevo elemento en la base de datos MySQL
func (r *mysqlRepository) SaveItem(it *Item) error {
	query := `INSERT INTO items (code, title, description, price, stock, status, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.Exec(query, it.Code, it.Title, it.Description, it.Price, it.Stock, it.Status, it.CreatedAt, it.UpdatedAt)
	return err
}

// ListItems lista todos los elementos de la base de datos MySQL
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

// UpdateItem actualiza un elemento existente en la base de datos MySQL
func (r *mysqlRepository) UpdateItem(it *Item) error {
	query := `UPDATE items SET code=?, title=?, description=?, price=?, stock=?, status=?, updated_at=? WHERE id=?`
	_, err := r.db.Exec(query, it.Code, it.Title, it.Description, it.Price, it.Stock, it.Status, it.UpdatedAt, it.ID)
	return err
}

// DeleteItem elimina un elemento de la base de datos MySQL
func (r *mysqlRepository) DeleteItem(id int) error {
	query := `DELETE FROM items WHERE id=?`
	_, err := r.db.Exec(query, id)
	return err
}
