package event

import (
	"time"

	"github.com/devpablocristo/golang/sdk/internal/core/event"
	"github.com/devpablocristo/golang/sdk/internal/core/location"
	user "github.com/devpablocristo/golang/sdk/internal/core/user/entities"
)

// EventRequest representa la solicitud para crear o actualizar un evento
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

// ToDomain convierte un EventRequest a un objeto del dominio Event
func (dto *EventRequest) ToDomain() *event.Event {
	// Convertir lista de asistentes
	attendees := make([]user.User, len(dto.Attendees))
	for i, a := range dto.Attendees {
		attendees[i] = user.User{
			UUID:         a.UUID,
			Username:     a.Username,
			PasswordHash: a.PasswordHash,
			CreatedAt:    a.CreatedAt,
		}
	}

	// Convertir lista de planificadores
	planners := make([]user.User, len(dto.Planners))
	for i, p := range dto.Planners {
		planners[i] = user.User{
			UUID:         p.UUID,
			Username:     p.Username,
			PasswordHash: p.PasswordHash,
			CreatedAt:    p.CreatedAt,
		}
	}

	// Convertir lista de organizadores
	organizers := make([]user.User, len(dto.Organizer))
	for i, o := range dto.Organizer {
		organizers[i] = user.User{
			UUID:      o.UUID,
			Username:  o.Username,
			PasswordHash:  o.PasswordHash,
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
