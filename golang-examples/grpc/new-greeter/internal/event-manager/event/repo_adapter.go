package event

import (
	"context"
	"log"

	usrdom "github.com/devpablocristo/qh/internal/users/domain"
)

type Repo struct {
	dao MongoEventDAOPort
}

func NewRepo(dao MongoEventDAOPort) RepoPort {
	return &Repo{
		dao: dao,
	}
}

func (r *Repo) CreateEvent(ctx context.Context, event *Event) (*Event, error) {
	createdEvent, err := r.dao.Create(ctx, event)
	if err != nil {
		log.Printf("Error creating event in repository: %v", err)
		return nil, err
	}
	return createdEvent, nil
}

func (r *Repo) DeleteEvent(ctx context.Context, eventID string) (*Event, error) {
	deletedEvent, err := r.dao.SoftDelete(ctx, eventID)
	if err != nil {
		log.Printf("Error soft deleting event with ID %s in repository: %v", eventID, err)
		return nil, err
	}
	return deletedEvent, nil
}

func (r *Repo) HardDeleteEvent(ctx context.Context, eventID string) (*Event, error) {
	deletedEvent, err := r.dao.HardDelete(ctx, eventID)
	if err != nil {
		log.Printf("Error soft deleting event with ID %s in repository: %v", eventID, err)
		return nil, err
	}
	return deletedEvent, nil
}

func (r *Repo) UpdateEvent(ctx context.Context, event *Event, eventID string) (*Event, error) {
	updatedEvent, err := r.dao.Update(ctx, event, eventID)
	if err != nil {
		log.Printf("Error updating event with ID %s in repository: %v", eventID, err)
		return nil, err
	}
	return updatedEvent, nil
}

func (r *Repo) ReviveEvent(ctx context.Context, eventID string) (*Event, error) {
	updatedEvent, err := r.dao.SoftUndelete(ctx, eventID)
	if err != nil {
		log.Printf("Error updating event with ID %s in repository: %v", eventID, err)
		return nil, err
	}
	return updatedEvent, nil
}

func (r *Repo) GetAllEvents(ctx context.Context) ([]Event, error) {
	eventsList, err := r.dao.List(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return eventsList, nil
}

func (r *Repo) GetEvent(ctx context.Context, eventID string) (*Event, error) {
	event, err := r.dao.FindByID(ctx, eventID)
	if err != nil {
		log.Printf("Error with event with ID %s in repository: %v", eventID, err)
		return nil, err
	}
	return event, nil
}

func (r *Repo) AddUserToEvent(ctx context.Context, eventID string, user *usrdom.User) (*Event, error) {
	event, err := r.dao.AddUserToEvent(ctx, eventID, user)
	if err != nil {
		log.Printf("dao error: %s, %v ", eventID, err)
		return nil, err
	}
	return event, nil
}
