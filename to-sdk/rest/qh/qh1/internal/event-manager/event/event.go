package event

import (
	usr "github.com/devpablocristo/qh/internal/user-manager/user"
)

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
