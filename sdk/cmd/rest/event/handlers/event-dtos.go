package handler

import (
	"os/user"
	"time"

	"github.com/devpablocristo/golang/sdk/internal/core/event"
	eve "github.com/devpablocristo/golang/sdk/internal/core/event"
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

type EventDTO struct {
	EventName   string      `json:"event_name"`
	Description string      `json:"description"`
	Date        string      `json:"date"`
	Location    LocationDTO `json:"location"`
	Organizer   []user.User `json:"organizer"`
	Attendees   []user.User `json:"attendees"`
	Planners    []user.User `json:"planners"`
	Tags        []string    `json:"tags"`
}

type LocationDTO struct {
	Address    string `json:"address"`
	City       string `json:"city"`
	State      string `json:"state"`
	Country    string `json:"country"`
	PostalCode string `json:"postalCode"`
}

func EventDtoToDomain(e *EventDTO) *event.Event {
	if e == nil {
		return nil
	}

	attendees := make([]user.User, 0, len(e.Attendees))
	for _, a := range e.Attendees {
		attendee := user.User{
			ID:   a.ID,
			Type: a.Type,
		}
		attendees = append(attendees, attendee)
	}

	planners := make([]user.User, 0, len(e.Planners))
	for _, p := range e.Planners {
		planner := user.User{
			ID:   p.ID,
			Type: p.Type,
		}
		planners = append(planners, planner)
	}

	organizers := make([]user.User, 0, len(e.Organizer))
	for _, o := range e.Organizer {
		organizer := user.User{
			ID:   o.ID,
			Type: o.Type,
		}
		organizers = append(organizers, organizer)
	}

	return &event.Event{
		EventName:   e.EventName,
		Description: e.Description,
		Date:        e.Date,
		Location: event.Location{
			Address:    e.Location.Address,
			City:       e.Location.City,
			State:      e.Location.State,
			Country:    e.Location.Country,
			PostalCode: e.Location.PostalCode,
		},
		Organizers: organizers, // Aquí cambiamos a la lista de organizadores
		Attendees:  attendees,
		Planners:   planners,
		Tags:       e.Tags,
	}
}

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
		//Location:    dto.Location,
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
