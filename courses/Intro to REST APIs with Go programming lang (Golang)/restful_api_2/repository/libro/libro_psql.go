package libroRepository

import (
	"database/sql"
	"restful_api_2/models"
)

type LibroRepository struct{}

func (l LibroRepository) ObtenerLibros(db *sql.DB, libro models.Libro, libros []models.Libro) ([]models.Libro, error) {
	rows, err := db.Query("select * from libros")

	if err != nil {
		return []models.Libro{}, err
	}

	for rows.Next() {
		err = rows.Scan(&libro.ID, &libro.Titulo, &libro.Autor, &libro.Año)
		libros = append(libros, libro)
	}

	if err != nil {
		return []models.Libro{}, err
	}

	return libros, nil
}

func (l LibroRepository) ObtenerLibro(db *sql.DB, lib models.Libro, id int) (models.Libro, error) {
	rows := db.QueryRow("SELECT * FROM libros WHERE id=$1;", id)
	err := rows.Scan(&lib.ID, &lib.Titulo, &lib.Autor, &lib.Año)

	return lib, err
}

func (l LibroRepository) AñadirLibro(db *sql.DB, lib models.Libro) (int, error) {
	err := db.QueryRow("INSERT INTO libros (titulo, autor, año) VALUES ($1, $2, $3) RETURNING id;", lib.Titulo, lib.Autor, lib.Año).Scan(&lib.ID)

	if err != nil {
		return 0, err
	}

	return lib.ID, nil
}

func (l LibroRepository) ActualizarLibro(db *sql.DB, lib models.Libro) (int64, error) {
	result, err := db.Exec("UPDATE libros SET titulo=$1, autor=$2, año=$3 WHERE id=$4;", lib.Titulo, lib.Autor, lib.Año, lib.ID)

	if err != nil {
		return 0, err
	}

	rowsUpdated, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsUpdated, nil
}

func (l LibroRepository) EliminarLibro(db *sql.DB, id int) (int64, error) {
	result, err := db.Exec("DELETE FROM libros WHERE id=$1", id)

	if err != nil {
		return 0, err
	}

	rowsDeleted, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsDeleted, nil
}
