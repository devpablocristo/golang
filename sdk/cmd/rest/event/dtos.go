package handler

import (
	"time"

	eve "github.com/devpablocristo/golang-sdk/internal/core/event"
)

type EventRequest struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
	Category    string    `json:"category"`
	CreatorID   string    `json:"creatorId"`
	IsPublic    bool      `json:"isPublic"`
	IsRecurring bool      `json:"isRecurring"`
	SeriesID    string    `json:"seriesId"`
	Status      string    `json:"status"`
}

func (dto *EventRequest) ToDomain() *eve.Event {
	return &eve.Event{
		ID:          dto.ID,
		Title:       dto.Title,
		Description: dto.Description,
		Location:    dto.Location,
		StartTime:   dto.StartTime,
		EndTime:     dto.EndTime,
		Category:    eve.Category(dto.Category),
		CreatorID:   dto.CreatorID,
		IsPublic:    dto.IsPublic,
		IsRecurring: dto.IsRecurring,
		SeriesID:    dto.SeriesID,
		Status:      eve.EventStatus(dto.Status),
	}
}
