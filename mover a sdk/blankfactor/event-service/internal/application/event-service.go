package application

import (
	"context"

	port "github.com/devpablocristo/blankfactor/event-service/internal/application/port"
	domain "github.com/devpablocristo/blankfactor/event-service/internal/domain"
)

type EventService struct {
	EventRepo port.EventRepo
}

func NewEventService(er port.EventRepo) port.EventService {
	return &EventService{
		EventRepo: er,
	}
}

func (es *EventService) CreateEvent(ctx context.Context, event *domain.Event) (*domain.Event, error) {
	event, err := es.EventRepo.CreateEvent(ctx, event)
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (es *EventService) GetAllEvents(ctx context.Context) ([]*domain.Event, error) {
	events, err := es.EventRepo.GetAllEvents(ctx)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (es *EventService) GetOverlappingEvents(ctx context.Context) ([][]*domain.Event, error) {
	events, err := es.EventRepo.GetAllEvents(ctx)
	if err != nil {
		return nil, err
	}

	overlappingEvents := [][]*domain.Event{}
	for i := 0; i < len(events); i++ {
		for j := i + 1; j < len(events); j++ {
			if isOverlapping(events[i], events[j]) {
				overlappingEvents = append(overlappingEvents, []*domain.Event{events[i], events[j]})
			}
		}
	}

	return overlappingEvents, nil
}

func isOverlapping(event1, event2 *domain.Event) bool {
	return event1.StartTime.Before(event2.EndTime) && event2.StartTime.Before(event1.EndTime)
}
