package port

import "crudmongodb/internal/domain"

type Service interface {
	Create(domain.Listing) error
	ReadAll() error
	ReadByID() error
	Update() error
	Delete() error
}
