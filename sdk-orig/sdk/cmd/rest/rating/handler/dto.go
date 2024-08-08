package rating

import (
	"github.com/devpablocristo/golang/sdk/internal/core/rating"
)

type RatingDTO struct {
	ID       int    `json:"id"`
	Item     string `json:"item"`
	Score    int    `json:"score"`
	Comments string `json:"comments"`
}

func ToRating(dto RatingDTO) *rating.Rating {
	return &rating.Rating{
		Item:     dto.Item,
		Score:    dto.Score,
		Comments: dto.Comments,
	}
}

func ToRatingDTO(r *rating.Rating) RatingDTO {
	return RatingDTO{
		Item:     r.Item,
		Score:    r.Score,
		Comments: r.Comments,
	}
}
