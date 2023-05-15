package port

import (
	"crudmongodb/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
)

type Repo interface {
	Create(domain.Listing) error
	Read(filter bson.M) ([]bson.M, error)
	Update(filter bson.M, update bson.M) error
	Delete(filter bson.M) error
}
