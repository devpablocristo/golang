package port

import "crudmongodb/internal/domain"

type Service interface {
	Create(domain.Listing) error
	Read() error
	Update() error
	Delete() error
}
