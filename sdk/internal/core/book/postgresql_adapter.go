package book

import (
	pqsql "github.com/devpablocristo/golang-sdk/pkg/postgresql/pq"
)

// Repository represents the repository structure.
type Repository struct {
	pgInst pqsql.PostgreSQLClientPort
}

// NewRepository initializes a new book repository.
func NewRepository(inst pqsql.PostgreSQLClientPort) RepositoryPort {
	return &Repository{
		pgInst: inst,
	}
}

// GetBooks retrieves all books from the database.
func (b *Repository) GetBooks(book Book, books []Book) ([]Book, error) {
	rows, err := b.pgInst.DB().Query("SELECT * FROM books")
	if err != nil {
		return []Book{}, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		if err != nil {
			return []Book{}, err
		}
		books = append(books, book)
	}

	return books, nil
}

// GetBook retrieves a single book by its ID.
func (b *Repository) GetBook(book Book, id int) (Book, error) {
	row := b.pgInst.DB().QueryRow("SELECT * FROM books WHERE id=$1", id)
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	return book, err
}

// AddBook inserts a new book into the database.
func (b *Repository) AddBook(book Book) (int, error) {
	var id int
	err := b.pgInst.DB().QueryRow("INSERT INTO books (title, author, year) VALUES($1, $2, $3) RETURNING id;",
		book.Title, book.Author, book.Year).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// UpdateBook updates an existing book in the database.
func (b *Repository) UpdateBook(book Book) (int64, error) {
	result, err := b.pgInst.DB().Exec("UPDATE books SET title=$1, author=$2, year=$3 WHERE id=$4",
		book.Title, book.Author, book.Year, book.ID)
	if err != nil {
		return 0, err
	}
	rowsUpdated, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsUpdated, nil
}

// RemoveBook deletes a book from the database.
func (b *Repository) RemoveBook(id int) (int64, error) {
	result, err := b.pgInst.DB().Exec("DELETE FROM books WHERE id=$1", id)
	if err != nil {
		return 0, err
	}
	rowsDeleted, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsDeleted, nil
}
