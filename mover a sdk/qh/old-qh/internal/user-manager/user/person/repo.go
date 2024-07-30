package person

import (
	"context"
	"log"
)

type Repo struct {
	dao PortMongoPersonDAO
}

func NewRepo(dao PortMongoPersonDAO) RepoPort {
	return &Repo{
		dao: dao,
	}
}

func (r *Repo) CreatePerson(ctx context.Context, event *Person) (*Person, error) {
	createdEvent, err := r.dao.Create(ctx, event)
	if err != nil {
		log.Printf("Error creating event in repository: %v", err)
		return nil, err
	}
	return createdEvent, nil
}

// func (r *Repo) DeleteEvent(ctx context.Context, eventID string) (*Person, error) {
// 	deletedEvent, err := r.dao.SoftDelete(ctx, eventID)
// 	if err != nil {
// 		log.Printf("Error soft deleting event with ID %s in repository: %v", eventID, err)
// 		return nil, err
// 	}
// 	return deletedEvent, nil
// }

// func (r *Repo) HardDeleteEvent(ctx context.Context, eventID string) (*Person, error) {
// 	deletedEvent, err := r.dao.HardDelete(ctx, eventID)
// 	if err != nil {
// 		log.Printf("Error soft deleting event with ID %s in repository: %v", eventID, err)
// 		return nil, err
// 	}
// 	return deletedEvent, nil
// }

// func (r *Repo) UpdateEvent(ctx context.Context, event *Person, eventID string) (*Person, error) {
// 	updatedEvent, err := r.dao.Update(ctx, event, eventID)
// 	if err != nil {
// 		log.Printf("Error updating event with ID %s in repository: %v", eventID, err)
// 		return nil, err
// 	}
// 	return updatedEvent, nil
// }

// func (r *Repo) ReviveEvent(ctx context.Context, eventID string) (*Person, error) {
// 	updatedEvent, err := r.dao.SoftUndelete(ctx, eventID)
// 	if err != nil {
// 		log.Printf("Error updating event with ID %s in repository: %v", eventID, err)
// 		return nil, err
// 	}
// 	return updatedEvent, nil
// }

// func (r *Repo) GetAllEvents(ctx context.Context) ([]domain.Person, error) {
// 	eventsList, err := r.dao.List(ctx)
// 	if err != nil {
// 		log.Println(err)
// 		return nil, err
// 	}
// 	return eventsList, nil
// }

// func (r *Repo) GetEvent(ctx context.Context, eventID string) (*Person, error) {
// 	event, err := r.dao.FindByID(ctx, eventID)
// 	if err != nil {
// 		log.Printf("Error with event with ID %s in repository: %v", eventID, err)
// 		return nil, err
// 	}
// 	return event, nil
// }
