package core

import (
	"context"

	eve "github.com/devpablocristo/qh/events/internal/core/event"
)

type eventUsecase struct {
	eve eve.RepositoryPort
}

func NewUseCase(r eve.RepositoryPort) EventUseCasePort {
	return &eventUsecase{
		eve: r,
	}
}

func (u *eventUsecase) CreateEvent(ctx context.Context, event *eve.Event) error {
	u.eve.CreateEvent(ctx, event)
	return nil
}

func (es *UseCase) CreateEvent(ctx context.Context, event *event.Event) (*event.Event, error) {
	event, err := es.repo.CreateEvent(ctx, event)
	if err != nil {
		log.Printf("Error creating event: %v", err)
		return nil, err
	}
	return event, nil
}

func (es *UseCase) DeleteEvent(ctx context.Context, eventID string) (*event.Event, error) {
	event, err := es.repo.DeleteEvent(ctx, eventID)
	if err != nil {
		log.Printf("Error deleting event with ID %s: %v", eventID, err)
		return nil, err
	}
	return event, nil
}

func (es *UseCase) HardDeleteEvent(ctx context.Context, eventID string) (*event.Event, error) {
	event, err := es.repo.HardDeleteEvent(ctx, eventID)
	if err != nil {
		log.Printf("Error deleting event with ID %s: %v", eventID, err)
		return nil, err
	}
	return event, nil
}

func (es *UseCase) UpdateEvent(ctx context.Context, event *event.Event, eventID string) (*event.Event, error) {
	event, err := es.repo.UpdateEvent(ctx, event, eventID)
	if err != nil {
		log.Printf("Error updating event with ID %s: %v", eventID, err)
		return nil, err
	}
	return event, nil
}

func (es *UseCase) ReviveEvent(ctx context.Context, eventID string) (*event.Event, error) {
	event, err := es.repo.ReviveEvent(ctx, eventID)
	if err != nil {
		log.Printf("Error undeleting event with ID %s: %v", eventID, err)
		return nil, err
	}
	return event, nil
}

func (es *UseCase) GetEvent(ctx context.Context, eventID string) (*event.Event, error) {
	event, err := es.repo.GetEvent(ctx, eventID)
	if err != nil {
		log.Printf("Error undeleting event with ID %s: %v", eventID, err)
		return nil, err
	}
	return event, nil
}

func (es *UseCase) GetAllEvents(ctx context.Context) ([]event.Event, error) {
	events, err := es.repo.GetAllEvents(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return events, nil
}

func (es *UseCase) AddUserToEvent(ctx context.Context, eventID string, user *usr.User) (*event.Event, error) {
	event, err := es.repo.AddUserToEvent(ctx, eventID, user)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return event, nil
}