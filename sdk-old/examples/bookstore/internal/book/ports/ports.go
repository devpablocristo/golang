package ports

import (
	"context"

	"github.com/devpablocristo/golang/sdk/services/bookstore/internal/book/entities"
)

type Repository interface {
	GetBook(context.Context, *entities.Book, int) (*entities.Book, error)
	AddBook(context.Context, *entities.Book) (int, error)
	UpdateBook(context.Context, *entities.Book) (int64, error)
	RemoveBook(context.Context, int) (int64, error)
}
