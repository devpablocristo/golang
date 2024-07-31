package event

import (
	"time"
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
	// Location    Location
	StartTime   time.Time
	EndTime     time.Time
	Category    Category
	CreatorID   string
	IsPublic    bool
	IsRecurring bool
	SeriesID    string
	Status      EventStatus
	// Organizers  []User
	// Attendees   []User
	// Planners    []User
	Tags []string
}

type InMemDB map[string]*Event

func EventToInterface(events []Event) []interface{} {
	result := make([]interface{}, len(events))
	for i, v := range events {
		result[i] = v
	}
	return result
}
