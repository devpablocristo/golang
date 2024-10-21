package ports

import "github.com/devpablocristo/golang/sdk/qh/rating/internal/core/entities"

type Repository interface {
	CreateRating(r *entities.Rating) (*entities.Rating, error)
	GetRatingByUUID(UUID string) (*entities.Rating, error)
	GetRatingByTarget(targetID string, targetType entities.TargetType) ([]*entities.Rating, error)
	UpdateRating(r *entities.Rating) (*entities.Rating, error)
	GetRatingByRaterAndTarget(raterID, targetID string, targetType entities.TargetType) (*entities.Rating, error) // Para verificar si ya existe una calificación de un usuario a un evento
}

type GrpcClient interface{}

type UseCases interface{}
