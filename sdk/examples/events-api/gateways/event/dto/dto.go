package dto

import (
	"time"

	event "github.com/devpablocristo/golang/sdk/examples/events-api/internal/event/entities"
	location "github.com/devpablocristo/golang/sdk/examples/locations-api/internal/location/entities"
	user "github.com/devpablocristo/golang/sdk/examples/users-api/internal/user/entities"
)

type EventRequest struct {
	ID          string            `json:"id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Location    location.Location `json:"location"`
	StartTime   time.Time         `json:"start_time"`
	EndTime     time.Time         `json:"end_time"`
	Category    string            `json:"category"`
	CreatorID   string            `json:"creator_id"`
	IsPublic    bool              `json:"is_public"`
	IsRecurring bool              `json:"is_recurring"`
	SeriesID    string            `json:"series_id"`
	Status      string            `json:"status"`
	Organizer   []user.User       `json:"organizer"`
	Attendees   []user.User       `json:"attendees"`
	Planners    []user.User       `json:"planners"`
	Tags        []string          `json:"tags"`
}

func (dto *EventRequest) ToDomain() *event.Event {
	// Convertir lista de asistentes
	attendees := make([]user.User, len(dto.Attendees))
	for i, a := range dto.Attendees {
		attendees[i] = user.User{
			UUID: a.UUID,
			Credentials: user.Credentials{
				Username:     a.Credentials.Username,
				PasswordHash: a.Credentials.Username,
			},
			CreatedAt: a.CreatedAt,
		}
	}

	// Convertir lista de planificadores
	planners := make([]user.User, len(dto.Planners))
	for i, p := range dto.Planners {
		planners[i] = user.User{
			UUID: p.UUID,
			Credentials: user.Credentials{
				Username:     p.Credentials.Username,
				PasswordHash: p.Credentials.PasswordHash,
			},
			CreatedAt: p.CreatedAt,
		}
	}

	// Convertir lista de organizadores
	organizers := make([]user.User, len(dto.Organizer))
	for i, o := range dto.Organizer {
		organizers[i] = user.User{
			UUID: o.UUID,
			Credentials: user.Credentials{
				Username:     o.Credentials.Username,
				PasswordHash: o.Credentials.PasswordHash,
			},
			CreatedAt: o.CreatedAt,
		}
	}

	// Retornar el objeto del dominio con todos los datos convertidos
	return &event.Event{
		ID:          dto.ID,
		Title:       dto.Title,
		Description: dto.Description,
		Location:    dto.Location,
		StartTime:   dto.StartTime,
		EndTime:     dto.EndTime,
		Category:    event.Category(dto.Category),
		CreatorID:   dto.CreatorID,
		IsPublic:    dto.IsPublic,
		IsRecurring: dto.IsRecurring,
		SeriesID:    dto.SeriesID,
		Status:      event.EventStatus(dto.Status),
		Organizers:  organizers,
		Attendees:   attendees,
		Planners:    planners,
		Tags:        dto.Tags,
	}
}
