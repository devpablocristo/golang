package book

type RepositoryPort interface {
	GetBooks(Book, []Book) ([]Book, error)
	GetBook(Book, int) (Book, error)
	AddBook(Book) (int, error)
	UpdateBook(Book) (int64, error)
	RemoveBook(int) (int64, error)
}
