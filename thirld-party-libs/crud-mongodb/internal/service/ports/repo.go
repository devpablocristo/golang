package port

import (
	"crudmongodb/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
)

type Repo interface {
	Create(listing domain.Listing) error
	ReadAll() ([]domain.Listing, error)
	ReadByID(filter bson.M) ([]bson.M, error)
	Update(filter bson.M, update bson.M) error
	Delete(filter bson.M) error
}
