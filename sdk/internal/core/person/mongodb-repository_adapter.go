package person

import (
	"context"
	"log"
)

type MongoDbRepository struct {
	dao DAOPort
}

func NewMongoDbRepository(dao DAOPort) RepositoryPort {
	return &MongoDbRepository{
		dao: dao,
	}
}

func (r *MongoDbRepository) SavePerson(ctx context.Context, person *Person) error {
	if err := r.dao.Create(ctx, person); err != nil {
		log.Printf("error creating event in repository: %v", err)
		return err
	}
	return nil
}

// func (r *MongoDbRepository) DeleteEvent(ctx context.Context, eventID string) (*Person, error) {
// 	deletedEvent, err := r.dao.SoftDelete(ctx, eventID)
// 	if err != nil {
// 		log.Printf("Error soft deleting event with ID %s in repository: %v", eventID, err)
// 		return nil, err
// 	}
// 	return deletedEvent, nil
// }

// func (r *MongoDbRepository) HardDeleteEvent(ctx context.Context, eventID string) (*Person, error) {
// 	deletedEvent, err := r.dao.HardDelete(ctx, eventID)
// 	if err != nil {
// 		log.Printf("Error soft deleting event with ID %s in repository: %v", eventID, err)
// 		return nil, err
// 	}
// 	return deletedEvent, nil
// }

// func (r *MongoDbRepository) UpdateEvent(ctx context.Context, event *Person, eventID string) (*Person, error) {
// 	updatedEvent, err := r.dao.Update(ctx, event, eventID)
// 	if err != nil {
// 		log.Printf("Error updating event with ID %s in repository: %v", eventID, err)
// 		return nil, err
// 	}
// 	return updatedEvent, nil
// }

// func (r *MongoDbRepository) ReviveEvent(ctx context.Context, eventID string) (*Person, error) {
// 	updatedEvent, err := r.dao.SoftUndelete(ctx, eventID)
// 	if err != nil {
// 		log.Printf("Error updating event with ID %s in repository: %v", eventID, err)
// 		return nil, err
// 	}
// 	return updatedEvent, nil
// }

// func (r *MongoDbRepository) GetAllEvents(ctx context.Context) ([]domain.Person, error) {
// 	eventsList, err := r.dao.List(ctx)
// 	if err != nil {
// 		log.Println(err)
// 		return nil, err
// 	}
// 	return eventsList, nil
// }

// func (r *MongoDbRepository) GetEvent(ctx context.Context, eventID string) (*Person, error) {
// 	event, err := r.dao.FindByID(ctx, eventID)
// 	if err != nil {
// 		log.Printf("Error with event with ID %s in repository: %v", eventID, err)
// 		return nil, err
// 	}
// 	return event, nil
// }
