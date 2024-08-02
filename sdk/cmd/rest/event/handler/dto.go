package handler

import (
	"time"

	"github.com/devpablocristo/golang/sdk/internal/core/event"
	"github.com/devpablocristo/golang/sdk/internal/core/user"
)

// EventRequest representa la solicitud para crear o actualizar un evento
type EventRequest struct {
	ID          string          `json:"id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Location    LocationRequest `json:"location"`
	StartTime   time.Time       `json:"start_time"`
	EndTime     time.Time       `json:"end_time"`
	Category    string          `json:"category"`
	CreatorID   string          `json:"creator_id"`
	IsPublic    bool            `json:"is_public"`
	IsRecurring bool            `json:"is_recurring"`
	SeriesID    string          `json:"series_id"`
	Status      string          `json:"status"`
	Organizer   []user.User     `json:"organizer"`
	Attendees   []user.User     `json:"attendees"`
	Planners    []user.User     `json:"planners"`
	Tags        []string        `json:"tags"`
}

// LocationRequest representa la solicitud para una ubicación
type LocationRequest struct {
	Address    string `json:"address"`
	City       string `json:"city"`
	State      string `json:"state"`
	Country    string `json:"country"`
	PostalCode string `json:"postal_code"`
}

	// ToDomain convierte un EventRequest a un objeto del dominio Event
func (dto *EventRequest) ToDomain() *event.Event {
	return &event.Event{
		ID:          dto.ID,
		Title:       dto.Title,
		Description: dto.Description,
		// Location: event.Location{
		// 	Address:    dto.Location.Address,
		// 	City:       dto.Location.City,
		// 	State:      dto.Location.State,
		// 	Country:    dto.Location.Country,
		// 	PostalCode: dto.Location.PostalCode,
		// },
		StartTime:   dto.StartTime,
		EndTime:     dto.EndTime,
		Category:    event.Category(dto.Category),
		CreatorID:   dto.CreatorID,
		IsPublic:    dto.IsPublic,
		IsRecurring: dto.IsRecurring,
		SeriesID:    dto.SeriesID,
		Status:      event.EventStatus(dto.Status),
		// Organizers:  dto.Organizer,
		// Attendees:   dto.Attendees,
		// Planners:    dto.Planners,
		Tags: dto.Tags,
	}
}

// ToDomain convierte un EventDTO a un objeto del dominio Event
// func ToDomain(e *EventDTO) *event.Event {
// 	if e == nil {
// 		return nil
// 	}

// 	attendees := make([]user.User, len(e.Attendees))
// 	for i, a := range e.Attendees {
// 		attendees[i] = user.User{
// 			ID:   a.ID,
// 			Name: a.Name,
// 			Type: a.Type,
// 		}
// 	}

// 	planners := make([]user.User, len(e.Planners))
// 	for i, p := range e.Planners {
// 		planners[i] = user.User{
// 			ID:   p.ID,
// 			Name: p.Name,
// 			Type: p.Type,
// 		}
// 	}

// 	organizers := make([]user.User, len(e.Organizer))
// 	for i, o := range e.Organizer {
// 		organizers[i] = user.User{
// 			ID:   o.ID,
// 			Name: o.Name,
// 			Type: o.Type,
// 		}
// 	}

// 	return &event.Event{
// 		Title:       e.EventName,
// 		Description: e.Description,
// 		Location: event.Location{
// 			Address:    e.Location.Address,
// 			City:       e.Location.City,
// 			State:      e.Location.State,
// 			Country:    e.Location.Country,
// 			PostalCode: e.Location.PostalCode,
// 		},
// 		StartTime:  e.StartTime,
// 		EndTime:    e.EndTime,
// 		Organizers: organizers,
// 		Attendees:  attendees,
// 		Planners:   planners,
// 		Tags:       e.Tags,
// 	}
// }
