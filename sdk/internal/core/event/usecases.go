package event

import (
	"context"
	"log"

	entities "github.com/devpablocristo/golang/sdk/internal/core/event/entities"
	ports "github.com/devpablocristo/golang/sdk/internal/core/event/ports"
)

type useCases struct {
	repository ports.Repository
}

func NewUseCases(r ports.Repository) ports.UseCases {
	return &useCases{
		repository: r,
	}
}

func (u *useCases) CreateEvent(ctx context.Context, Event *entities.Event) error {
	if err := u.repository.CreateEvent(ctx, Event); err != nil {
		log.Printf("error creating Event: %v", err)
		return err
	}
	return nil
}

// func (es *useCases) DeleteEvent(ctx context.Context, EventID string) (*Event.entities.Event, error) {
// 	Event, err := es.repo.DeleteEvent(ctx, EventID)
// 	if err != nil {
// 		log.Printf("Error deleting Event with ID %s: %v", EventID, err)
// 		return nil, err
// 	}
// 	return Event, nil
// }

// func (es *useCases) HardDeleteEvent(ctx context.Context, EventID string) (*Event.entities.Event, error) {
// 	Event, err := es.repo.HardDeleteEvent(ctx, EventID)
// 	if err != nil {
// 		log.Printf("Error deleting Event with ID %s: %v", EventID, err)
// 		return nil, err
// 	}
// 	return Event, nil
// }

// func (es *useCases) UpdateEvent(ctx context.Context, Event *Event.entities.Event, EventID string) (*Event.entities.Event, error) {
// 	Event, err := es.repo.UpdateEvent(ctx, Event, EventID)
// 	if err != nil {
// 		log.Printf("Error updating Event with ID %s: %v", EventID, err)
// 		return nil, err
// 	}
// 	return Event, nil
// }

// func (es *useCases) ReviveEvent(ctx context.Context, EventID string) (*Event.entities.Event, error) {
// 	Event, err := es.repo.ReviveEvent(ctx, EventID)
// 	if err != nil {
// 		log.Printf("Error undeleting Event with ID %s: %v", EventID, err)
// 		return nil, err
// 	}
// 	return Event, nil
// }

// func (es *useCases) GetEvent(ctx context.Context, EventID string) (*Event.entities.Event, error) {
// 	Event, err := es.repo.GetEvent(ctx, EventID)
// 	if err != nil {
// 		log.Printf("Error undeleting Event with ID %s: %v", EventID, err)
// 		return nil, err
// 	}
// 	return Event, nil
// }

// func (es *useCases) GetAllEvents(ctx context.Context) ([]Event.entities.Event, error) {
// 	Events, err := es.repo.GetAllEvents(ctx)
// 	if err != nil {
// 		log.Println(err)
// 		return nil, err
// 	}
// 	return Events, nil
// }

// func (es *useCases) AddUserToEvent(ctx context.Context, EventID string, user *usr.User) (*Event.entities.Event, error) {
// 	Event, err := es.repo.AddUserToEvent(ctx, EventID, user)
// 	if err != nil {
// 		log.Println(err)
// 		return nil, err
// 	}
// 	return Event, nil
// }
