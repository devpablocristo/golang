### Data Transfer Object (DTO)

A Data Transfer Object (DTO) is a design pattern used to transfer data between systems or layers of an application in a simplified manner. DTOs encapsulate data and are commonly used to transfer relevant information between client and server, or between different parts of a system, without exposing internal details of how those data are stored or managed. This promotes greater flexibility and decoupling between the layers of an application, being especially useful in the development of APIs to model both incoming requests and outgoing responses, allowing to control exactly what data is exposed through the API.

#### Example of DTO in Go for a book management application:

```go
package dto

// LibroCreateDTO is used for creating books, defining what is necessary to add a new book.
type LibroCreateDTO struct {
    Title  string `json:"title"`
    Author string `json:"author"`
    Summary string `json:"summary,omitempty"` // Optional
    Year    int    `json:"year"`
}

// LibroResponseDTO is used to send book data to the client, defining how a book is presented in responses.
type LibroResponseDTO struct {
    ID      int64  `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
    Summary string `json:"summary,omitempty"` // Optional
    Year    int    `json:"year"`
}
```

### Presenter

The Presenter is part of the Model-View-Presenter (MVP) pattern, acting as an intermediary between the view (UI) and the model (business data). Its main goal is to prepare the model data for presentation in the view, containing specific presentation logic that decides how the data should be shown to the user.

#### Example of Presenter in Go:

```go
package presenter

import (
    "fmt"
    "strings"
    "yourProject/dto"
)

// LibroPresenter is responsible for preparing book data for the view.
type LibroPresenter struct {
    Book dto.LibroDTO
}

// FormatTitle converts the book title to uppercase.
func (lp *LibroPresenter) FormatTitle() string {
    return strings.ToUpper(lp.Book.Title)
}

// BookDetails composes a string with the formatted details of the book for presentation.
func (lp *LibroPresenter) BookDetails() string {
    return fmt.Sprintf("Title: %s, Author: %s, Publication Year: %d",
        lp.FormatTitle(), lp.Book.Author, lp.Book.Publication)
}
```

### Data Access Object (DAO)

The DAO (Data Access Object) is a design pattern that provides an abstract interface to access data stored in a database, file, or any other persistence medium. Its purpose is to separate the data access logic from the business logic of the application, allowing the latter to be independent of the underlying storage mechanism.

#### Example of DAO in Go for the Book entity:

```go
package dao

import (
    "errors"
    "yourProject/models"
)

// LibroDAO defines the data access operations for books.
type LibroDAO interface {
    Create(book models.Libro) error
    GetById(id string) (models.Libro, error)
    Update(book models.Libro) error
    Delete(id string) error
}

// LibroDAOImpl implements LibroDAO using a map as in-memory storage.
type LibroDAOImpl struct {
    books map[string]models.Libro
}

func NewLibroDAOImpl() *LibroDAOImpl {
    return &LibroDAOImpl{books: make(map[string]models.Libro)}
}

// Implementation of the Create, GetById, Update, and Delete methods...
```

These examples illustrate how the DTO is used for efficient data transfer between layers or services, the Presenter to adapt and format data for presentation, and the DAO to encapsulate access and manipulation of data, demonstrating the importance of each pattern in software design and architecture.