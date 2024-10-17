package rating

import (
	"github.com/devpablocristo/golang/sdk/qh/rating/internal/core/entities"
)

type RatingDTO struct {
	ID       int    `json:"id"`
	Item     string `json:"item"`
	Score    int    `json:"score"`
	Comments string `json:"comments"`
}

func ToRating(dto RatingDTO) *entities.Rating {
	return &entities.Rating{
		Item:     dto.Item,
		Score:    dto.Score,
		Comments: dto.Comments,
	}
}

func ToRatingDTO(r *entities.Rating) RatingDTO {
	return RatingDTO{
		Item:     r.Item,
		Score:    r.Score,
		Comments: r.Comments,
	}
}
