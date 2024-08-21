package event

import (
	"time"

	"github.com/devpablocristo/golang/sdk/internal/core/location"
	"github.com/devpablocristo/golang/sdk/internal/core/user/entities"
)

type EventStatus string
type Category string

const (
	EventScheduled EventStatus = "scheduled"
	EventOngoing   EventStatus = "ongoing"
	EventCompleted EventStatus = "completed"
	EventCancelled EventStatus = "cancelled"
	EventPostponed EventStatus = "postponed"

	CategoryMusic         Category = "music"
	CategorySports        Category = "sports"
	CategoryEducation     Category = "education"
	CategoryEntertainment Category = "entertainment"
	CategoryHealth        Category = "health"
	CategoryBusiness      Category = "business"
	CategoryTechnology    Category = "technology"
	CategoryCharity       Category = "charity"
	CategoryReligion      Category = "religion"
	CategoryFamily        Category = "family"
	CategoryGovernment    Category = "government"
	CategoryPrivate       Category = "private"
)

type Event struct {
	ID          string
	Title       string
	Description string
	Location    location.Location
	StartTime   time.Time
	EndTime     time.Time
	Category    Category
	CreatorID   string
	IsPublic    bool
	IsRecurring bool
	SeriesID    string
	Status      EventStatus
	Organizers  []entities.User
	Attendees   []entities.User
	Planners    []entities.User
	Tags        []string
}

type InMemDB map[string]*Event

func ToInterface(events []Event) []any {
	result := make([]any, len(events))
	for i, v := range events {
		result[i] = v
	}
	return result
}
