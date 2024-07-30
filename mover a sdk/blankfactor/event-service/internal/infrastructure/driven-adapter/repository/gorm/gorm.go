package gorm

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	domain "github.com/devpablocristo/blankfactor/event-service/internal/domain"
)

type EventGormRepository struct {
	DB *gorm.DB
}

// NewDatabase : intializes and returns mysql db
func NewEventRepository(db *gorm.DB) *EventGormRepository {
	return &EventGormRepository{
		DB: db,
	}
}

func (r *EventGormRepository) CreateEvent(ctx context.Context, event *domain.Event) (*domain.Event, error) {
	e := &eventDAO{
		StartTime: event.StartTime,
		EndTime:   event.EndTime,
	}

	err := r.DB.Create(e).Error
	if err != nil {
		return nil, fmt.Errorf("error creating event: %v", err)
	}

	id := e.ID

	savedEvent, err := r.GetEventByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error creating event: %v", err)
	}

	return savedEvent, nil
}

func (r *EventGormRepository) GetEventByID(ctx context.Context, id int64) (*domain.Event, error) {
	var e eventDAO
	err := r.DB.Where("id = ?", id).First(&e).Error
	if err != nil {
		return nil, fmt.Errorf("error getting event by ID: %v", err)
	}

	event := e.dao2Event()

	return event, nil
}

func (r *EventGormRepository) UpdateEvent(ctx context.Context, event *domain.Event, id int64) error {
	return nil
}

func (r *EventGormRepository) DeleteEvent(ctx context.Context, id int64) error {
	return nil
}

func (r *EventGormRepository) GetAllEvents(ctx context.Context) ([]*domain.Event, error) {
	var events []*eventDAO

	err := r.DB.Find(&events).Error
	if err != nil {
		return nil, fmt.Errorf("error getting events: %v", err)
	}

	var domainEvents []*domain.Event
	for _, e := range events {
		domainEvents = append(domainEvents, e.dao2Event())
	}

	return domainEvents, nil
}
