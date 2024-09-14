package event

import (
	"context"
	"log"

	entities "github.com/devpablocristo/golang/sdk/services/events-api/internal/event/entities"
	ports "github.com/devpablocristo/golang/sdk/services/events-api/internal/event/ports"
)

type MongoDbRepository struct {
	dao ports.DAO
}

func NewMongoDbRepository(dao ports.DAO) ports.Repository {
	return &MongoDbRepository{
		dao: dao,
	}
}

func (r *MongoDbRepository) CreateEvent(ctx context.Context, event *entities.Event) error {
	if err := r.dao.Create(ctx, event); err != nil {
		log.Printf("Error creating event in repository: %v", err)
		return err
	}
	return nil
}

// func (r *MongoDbRepository) DeleteEvent(ctx context.Context, eventID string) (*Event, error) {
// 	deletedEvent, err := r.dao.SoftDelete(ctx, eventID)
// 	if err != nil {
// 		log.Printf("Error soft deleting event with ID %s in repository: %v", eventID, err)
// 		return nil, err
// 	}
// 	return deletedEvent, nil
// }

// func (r *MongoDbRepository) HardDeleteEvent(ctx context.Context, eventID string) (*Event, error) {
// 	deletedEvent, err := r.dao.HardDelete(ctx, eventID)
// 	if err != nil {
// 		log.Printf("Error soft deleting event with ID %s in repository: %v", eventID, err)
// 		return nil, err
// 	}
// 	return deletedEvent, nil
// }

// func (r *MongoDbRepository) UpdateEvent(ctx context.Context, event *Event, eventID string) (*Event, error) {
// 	updatedEvent, err := r.dao.Update(ctx, event, eventID)
// 	if err != nil {
// 		log.Printf("Error updating event with ID %s in repository: %v", eventID, err)
// 		return nil, err
// 	}
// 	return updatedEvent, nil
// }

// func (r *MongoDbRepository) ReviveEvent(ctx context.Context, eventID string) (*Event, error) {
// 	updatedEvent, err := r.dao.SoftUndelete(ctx, eventID)
// 	if err != nil {
// 		log.Printf("Error updating event with ID %s in repository: %v", eventID, err)
// 		return nil, err
// 	}
// 	return updatedEvent, nil
// }

// func (r *MongoDbRepository) GetAllEvents(ctx context.Context) ([]Event, error) {
// 	eventsList, err := r.dao.List(ctx)
// 	if err != nil {
// 		log.Println(err)
// 		return nil, err
// 	}
// 	return eventsList, nil
// }

// func (r *MongoDbRepository) GetEvent(ctx context.Context, eventID string) (*Event, error) {
// 	event, err := r.dao.FindByID(ctx, eventID)
// 	if err != nil {
// 		log.Printf("Error with event with ID %s in repository: %v", eventID, err)
// 		return nil, err
// 	}
// 	return event, nil
// }

// func (r *MongoDbRepository) AddUserToEvent(ctx context.Context, eventID string, user *usr.User) (*Event, error) {
// 	event, err := r.dao.AddUserToEvent(ctx, eventID, user)
// 	if err != nil {
// 		log.Printf("dao error: %s, %v ", eventID, err)
// 		return nil, err
// 	}
// 	return event, nil
// }
