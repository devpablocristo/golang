package usecase

import (
	"fmt"

	"github.com/mercadolibre/inventory/internal/entity"
)

type BookUsecase interface {
	GetAllBooks() ([]entity.Book, error)
	GetBookByID(id int) (entity.Book, error)
	AddBook(book entity.Book) (entity.Book, error)
}

type bookUsecase struct {
	repo entity.BookRepository
}

func NewBookUsecase(repo entity.BookRepository) BookUsecase {
	return &bookUsecase{
		repo: repo,
	}
}

func (u *bookUsecase) GetAllBooks() ([]entity.Book, error) {
	books, err := u.repo.GetBooks()
	if err != nil {
		return books, fmt.Errorf("error in repository: %w", err)
	}

	return books, nil
}

func (u *bookUsecase) GetBookByID(id int) (entity.Book, error) {
	book, err := u.repo.GetBookByID(uint(id))
	if err != nil {
		return entity.Book{}, fmt.Errorf("error in repository: %w", err)
	}

	return book, nil
}

func (u *bookUsecase) AddBook(book entity.Book) (entity.Book, error) {
	exist, err := u.repo.CheckBookByCode(book.Code)
	if err != nil {
		return book, fmt.Errorf("error in repository: %w", err)
	}

	if exist {
		return book, entity.BookAlreadyExist{
			Message: "book already exist",
		}
	}

	if err := u.repo.AddBook(&book); err != nil {
		return entity.Book{}, fmt.Errorf("error in repository: %w", err)
	}

	return book, nil
}
