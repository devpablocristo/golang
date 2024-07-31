package event

import (
	"time"
)

const (
	EventScheduled EventStatus = "scheduled"
	EventOngoing   EventStatus = "ongoing"
	EventCompleted EventStatus = "completed"
	EventCancelled EventStatus = "cancelled"
	EventPostponed EventStatus = "postponed"
)

type EventStatus string

const (
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

type Category string

type Event struct {
	ID          string
	Title       string
	Description string
	Location    string
	StartTime   time.Time
	EndTime     time.Time
	Category    Category
	CreatorID   string
	IsPublic    bool
	IsRecurring bool
	SeriesID    string
	Status      EventStatus
}

type InMemDB map[string]*Event

type Event struct {
	EventName   string
	Description string
	Date        string
	Location    Location
	Organizers  []usr.User
	Attendees   []usr.User
	Planners    []usr.User
	Tags        []string
}

type Location struct {
	Address    string
	City       string
	State      string
	Country    string
	PostalCode string
}

func EventToInterface(events []Event) []interface{} {
	result := make([]interface{}, len(events))
	for i, v := range events {
		result[i] = v
	}
	return result
}
