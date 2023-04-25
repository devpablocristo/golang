package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/mercadolibre/inventory/internal/entity"
)

type bookDAO struct {
	ID        uint
	Code      string
	Author    string
	Title     string
	Price     int
	Stock     int
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (b bookDAO) toBookDomain() entity.Book {
	return entity.Book{
		ID:        b.ID,
		Code:      b.Code,
		Author:    b.Author,
		Title:     b.Title,
		Price:     b.Price,
		Stock:     b.Stock,
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
	}
}

type mysqlBookRepository struct {
	conn *sqlx.DB
}

func NewMySQLBookRepository(db *sqlx.DB) entity.BookRepository {
	return &mysqlBookRepository{
		conn: db,
	}
}

func (r *mysqlBookRepository) GetBookByID(id uint) (entity.Book, error) {
	var book entity.Book
	var bookDB bookDAO

	err := r.conn.Get(&bookDB, "SELECT * FROM books WHERE id=?", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return book, entity.BookNotFound{
				Message: "book not found",
			}
		}
		return book, fmt.Errorf("error getting book: %w", err)
	}

	return bookDB.toBookDomain(), nil
}

func (r *mysqlBookRepository) CheckBookByCode(code string) (bool, error) {
	var exist bool
	err := r.conn.Get(&exist, `SELECT EXISTS(SELECT id FROM books WHERE code =  ?)`, code)
	if err != nil {
		return exist, fmt.Errorf("error getting book: %w", err)
	}

	return exist, nil
}

func (r *mysqlBookRepository) AddBook(book *entity.Book) error {
	createdAt := time.Now()

	result, err := r.conn.Exec(`INSERT INTO books 
		(code, title, author, price, stock, created_at, updated_at) 
		VALUES(?,?,?,?,?,?,?)`, book.Code, book.Title, book.Author, book.Price, book.Stock, createdAt, createdAt)

	if err != nil {
		return fmt.Errorf("error inserting book: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error saving book: %w", err)
	}

	book.ID = uint(id)
	book.CreatedAt = createdAt
	book.UpdatedAt = createdAt

	return nil
}

func (r *mysqlBookRepository) GetBooks() ([]entity.Book, error) {
	var books []entity.Book
	var booksDB []bookDAO

	err := r.conn.Select(&booksDB, "SELECT * FROM books LIMIT 10")
	if err != nil {
		return books, fmt.Errorf("error getting all books: %w", err)
	}

	for _, b := range booksDB {
		books = append(books, b.toBookDomain())
	}

	return books, nil
}
