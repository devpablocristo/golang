package book

import "time"

type Repository interface {
	Save(d *Book) error
	Get(id string) (*Book, error)
	Delete(id string) error
	//SearchByStatus(status task.Status) ([]Developer, error)
}

type Book struct {
	ID        string
	Name      string
	Author    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// type Publisher interface {
// 	Publish(ownerId string, t Task) error
// }

// type Status int

// const (
// 	StatusNone Status = iota
// 	StatusPending
// 	StatusCancelled
// 	StatusComplete
// )

// func (s Status) String() string {
// 	return [...]string{"none", "pending", "cancelled", "completed"}[s]
// }
