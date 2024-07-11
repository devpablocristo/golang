package inventory

import (
	"github.com/devpablocristo/clean-pod-2/internal/inventory/book"
)

/*
por lo q veo aqui es todo domain con una interface (port) de algo relacionado con la infraestructura
*/

// interface adapter
// type Repository interface {
// 	Save(d *Developer) error
// 	Get(id string) (*Developer, error)
// 	Delete(id string) error
// 	//SearchByStatus(status task.Status) ([]Developer, error)
// }

// domain
type Developer struct {
	ID    string
	Books []book.Book
	//Task      task.Task
}

// domain func
// func (d Developer) IsBusy() bool {
// 	return d.Task.Status == task.StatusPending
// }

// domain alias
//type Seniority int

// domain func
// func SeniorityFromString(value string) (Seniority, error) {
// 	switch value {
// 	case "senior":
// 		return Senior, nil
// 	case "semi_senior":
// 		return SemiSenior, nil
// 	case "analyst":
// 		return Analyst, nil
// 	case "junior":
// 		return Junior, nil
// 	default:
// 		return 0, fmt.Errorf("%s is not a valid seniority", value)
// 	}
// }

// domain consts
// const (
// 	Senior Seniority = iota
// 	SemiSenior
// 	Analyst
// 	Junior
// )

// domain func
// func (s Seniority) String() string {
// 	return [...]string{"senior", "semi_senior", "analyst", "junior"}[s]
// }
