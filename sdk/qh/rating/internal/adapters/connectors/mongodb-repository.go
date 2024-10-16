package ratconn

import (
	"fmt"

	sdk "github.com/devpablocristo/golang/sdk/pkg/databases/nosql/mongodb/mongo-driver"
	entities "github.com/devpablocristo/golang/sdk/qh/rating/internal/core/entities"
	ports "github.com/devpablocristo/golang/sdk/qh/rating/internal/core/ports"
)

type MongoDbRepository struct {
	repository ports.Repository
}

func NewMongoDbRepository() (ports.Repository, error) {
	r, err := sdk.Bootstrap()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize MongoDB client: %w", err)
	}

	return &MongoDbRepository{
		repository: r,
	}, nil
}

func (rat *MongoDbRepository) CreateRating(r *entities.Rating) (*entities.Rating, error) {
	return nil, nil
}

func (rat *MongoDbRepository) GetRatingByUUID(UUID string) (*entities.Rating, error) {
	return nil, nil
}

func (rat *MongoDbRepository) GetRatingByTarget(targetID string, targetType entities.TargetType) ([]*entities.Rating, error) {
	return nil, nil
}

func (rat *MongoDbRepository) UpdateRating(r *entities.Rating) (*entities.Rating, error) {
	return nil, nil
}

func (rat *MongoDbRepository) GetRatingByRaterAndTarget(raterID, targetID string, targetType entities.TargetType) (*entities.Rating, error) {
	return nil, nil
}
