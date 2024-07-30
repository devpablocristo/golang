package event

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	usr "github.com/devpablocristo/qh/internal/user-manager/user"
)

type EventDAO struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
	DeletedAt   time.Time          `bson:"deleted_at"`
	Deleted     bool               `bson:"deleted"`
	Version     int                `bson:"version"`
	EventName   string             `bson:"event_name"`
	Description string             `bson:"description"`
	Date        string             `bson:"date"`
	Location    *LocationDAO       `bson:"location"`
	Organizers  []usr.User         `bson:"organizer"`
	Attendees   []usr.User         `bson:"attendees"`
	Planners    []usr.User         `bson:"planners"`
	Tags        []string           `bson:"tags"`
}

type LocationDAO struct {
	Address    string `bson:"address"`
	City       string `bson:"city"`
	State      string `bson:"state"`
	Country    string `bson:"country"`
	PostalCode string `bson:"postal_code"`
}

func EventDaoToDomain(dao *EventDAO) *Event {
	if dao == nil {
		return nil
	}

	return &Event{
		EventName:   dao.EventName,
		Description: dao.Description,
		Date:        dao.Date,
		Location: Location{
			Address:    dao.Location.Address,
			City:       dao.Location.City,
			State:      dao.Location.State,
			Country:    dao.Location.Country,
			PostalCode: dao.Location.PostalCode,
		},
		Organizers: dao.Organizers,
		Attendees:  dao.Attendees,
		Planners:   dao.Planners,
		Tags:       dao.Tags,
	}
}

func EventDomainToDao(domain *Event) *EventDAO {
	if domain == nil {
		return nil
	}

	return &EventDAO{
		EventName:   domain.EventName,
		Description: domain.Description,
		Date:        domain.Date,
		Location: &LocationDAO{
			Address:    domain.Location.Address,
			City:       domain.Location.City,
			State:      domain.Location.State,
			Country:    domain.Location.Country,
			PostalCode: domain.Location.PostalCode,
		},
		Organizers: domain.Organizers,
		Attendees:  domain.Attendees,
		Planners:   domain.Planners,
		Tags:       domain.Tags,
	}
}
