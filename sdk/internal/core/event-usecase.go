package core

import (
	"context"
	"log"

	eve "github.com/devpablocristo/qh/events/internal/core/event"
	usr "github.com/devpablocristo/qh/events/internal/core/user"
)

type EventUseCasePort interface {
	CreateEvent(context.Context, *eve.Event) error
	DeleteEvent(context.Context, string) (*Event, error)
	HardDeleteEvent(context.Context, string) (*Event, error)
	UpdateEvent(context.Context, *Event, string) (*Event, error)
	ReviveEvent(context.Context, string) (*Event, error)
	GetEvent(context.Context, string) (*Event, error)
	GetAllEvents(context.Context) ([]Event, error)
	AddUserToEvent(context.Context, string, *usr.User) (*Event, error)
}

type eventUsecase struct {
	eve eve.RepositoryPort
}

func NewUseCase(r eve.RepositoryPort) EventUseCasePort {
	return &eventUsecase{
		eve: r,
	}
}

func (u *eventUsecase) CreateEvent(ctx context.Context, event *eve.Event) error {
	if err := u.eve.CreateEvent(ctx, event); err != nil {
		log.Printf("error creating event: %v", err)
		return err
	}
	return nil
}

// func (es *eventUsecase) DeleteEvent(ctx context.Context, eventID string) (*event.Event, error) {
// 	event, err := es.repo.DeleteEvent(ctx, eventID)
// 	if err != nil {
// 		log.Printf("Error deleting event with ID %s: %v", eventID, err)
// 		return nil, err
// 	}
// 	return event, nil
// }

// func (es *eventUsecase) HardDeleteEvent(ctx context.Context, eventID string) (*event.Event, error) {
// 	event, err := es.repo.HardDeleteEvent(ctx, eventID)
// 	if err != nil {
// 		log.Printf("Error deleting event with ID %s: %v", eventID, err)
// 		return nil, err
// 	}
// 	return event, nil
// }

// func (es *eventUsecase) UpdateEvent(ctx context.Context, event *event.Event, eventID string) (*event.Event, error) {
// 	event, err := es.repo.UpdateEvent(ctx, event, eventID)
// 	if err != nil {
// 		log.Printf("Error updating event with ID %s: %v", eventID, err)
// 		return nil, err
// 	}
// 	return event, nil
// }

// func (es *eventUsecase) ReviveEvent(ctx context.Context, eventID string) (*event.Event, error) {
// 	event, err := es.repo.ReviveEvent(ctx, eventID)
// 	if err != nil {
// 		log.Printf("Error undeleting event with ID %s: %v", eventID, err)
// 		return nil, err
// 	}
// 	return event, nil
// }

// func (es *eventUsecase) GetEvent(ctx context.Context, eventID string) (*event.Event, error) {
// 	event, err := es.repo.GetEvent(ctx, eventID)
// 	if err != nil {
// 		log.Printf("Error undeleting event with ID %s: %v", eventID, err)
// 		return nil, err
// 	}
// 	return event, nil
// }

// func (es *eventUsecase) GetAllEvents(ctx context.Context) ([]event.Event, error) {
// 	events, err := es.repo.GetAllEvents(ctx)
// 	if err != nil {
// 		log.Println(err)
// 		return nil, err
// 	}
// 	return events, nil
// }

// func (es *eventUsecase) AddUserToEvent(ctx context.Context, eventID string, user *usr.User) (*event.Event, error) {
// 	event, err := es.repo.AddUserToEvent(ctx, eventID, user)
// 	if err != nil {
// 		log.Println(err)
// 		return nil, err
// 	}
// 	return event, nil
// }
