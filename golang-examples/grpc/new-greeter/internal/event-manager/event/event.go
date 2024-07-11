package event

import (
	usrdom "github.com/devpablocristo/qh/internal/users/domain"
)

type Event struct {
	EventName   string
	Description string
	Date        string
	Location    Location
	Organizers  []usrdom.User
	Attendees   []usrdom.User
	Planners    []usrdom.User
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
