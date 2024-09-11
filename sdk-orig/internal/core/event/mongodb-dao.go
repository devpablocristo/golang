package event

import (
	"context"

	entities "github.com/devpablocristo/golang/sdk/internal/core/event/entities"
	ports "github.com/devpablocristo/golang/sdk/internal/core/event/ports"
)

type mongoEventDAO struct {
	repository ports.Repository
}

func NewMongoEventDAO(r ports.Repository) ports.DAO {
	return &mongoEventDAO{
		repository: r,
	}
}

func (ed *mongoEventDAO) Create(ctx context.Context, event *entities.Event) error {
	// dao := EventDomainToDao(event)
	// dao.CreatedAt = time.Now()
	// if _, err := ed.repository.DB().Collection("events").InsertOne(ctx, dao); err != nil {
	// 	log.Printf("error inserting event: %v", err)
	// 	return err
	// }
	return nil
}

// func (ed *mongoEventDAO) FindByID(ctx context.Context, ID string) (*entities.Event, error) {
// 	eventID, err := primitive.ObjectIDFromHex(ID)
// 	if err != nil {
// 		log.Printf("Error converting ID: %v", err)
// 		return nil, err
// 	}

// 	filter := bson.M{"_id": eventID}
// 	dao := &EventDAO{}
// 	if err := ed.repository.DB().Collection("events").FindOne(ctx, filter).Decode(dao); err != nil {
// 		log.Printf("Error decoding event: %v", err)
// 		return nil, err
// 	}
// 	return EventDaoToDomain(dao), nil
// }

// func (ed *mongoEventDAO) Update(ctx context.Context, event *entities.Event, ID string) (*entities.Event, error) {
// 	eventID, err := primitive.ObjectIDFromHex(ID)
// 	if err != nil {
// 		log.Printf("Error converting ID: %v", err)
// 		return nil, err
// 	}

// 	dao := EventDomainToDao(event)
// 	updateFields := checkEventFields(dao)

// 	if updateFields == nil {
// 		log.Printf("Error: event is nil")
// 		return nil, fmt.Errorf("event is nil")
// 	}

// 	if len(updateFields) == 0 {
// 		log.Printf("Error: no fields to update")
// 		return nil, fmt.Errorf("no fields to update")
// 	}

// 	updateFields["updated_at"] = time.Now()

// 	filter := bson.M{"_id": eventID}
// 	update := bson.M{"$set": updateFields}

// 	result, err := ed.repository.DB().Collection("events").UpdateOne(ctx, filter, update)
// 	if err != nil {
// 		log.Printf("Error updating event: %v", err)
// 		return nil, err
// 	}

// 	if result.MatchedCount == 0 {
// 		log.Printf("Error: no event found with ID: %s", eventID.Hex())
// 		return nil, fmt.Errorf("no event found with ID: %s", eventID.Hex())
// 	}

// 	var updatedDAO EventDAO
// 	if err := ed.repository.DB().Collection("events").FindOne(ctx, filter).Decode(&updatedDAO); err != nil {
// 		log.Printf("Error decoding updated event: %v", err)
// 		return nil, err
// 	}

// 	log.Printf("updated %d events", result.ModifiedCount)

// 	return EventDaoToDomain(&updatedDAO), nil
// }

// func (ed *mongoEventDAO) HardDelete(ctx context.Context, ID string) (*entities.Event, error) {
// 	eventID, err := primitive.ObjectIDFromHex(ID)
// 	if err != nil {
// 		log.Printf("Error converting ID: %v", err)
// 		return nil, err
// 	}

// 	filter := bson.M{"_id": eventID}

// 	deletedDAO := &EventDAO{}
// 	if err := ed.repository.DB().Collection("events").FindOne(ctx, filter).Decode(deletedDAO); err != nil {
// 		log.Printf("Error decoding deleted event: %v", err)
// 		return nil, err
// 	}

// 	result, err := ed.repository.DB().Collection("events").DeleteOne(ctx, filter)
// 	if err != nil {
// 		log.Printf("Error deleting event: %v", err)
// 		return nil, err
// 	}

// 	log.Printf("deleted %d events", result.DeletedCount)
// 	return EventDaoToDomain(deletedDAO), nil
// }

// func (ed *mongoEventDAO) List(ctx context.Context) ([]Event, error) {
// 	cursor, err := ed.repository.DB().Collection("events").Find(ctx, bson.M{})
// 	if err != nil {
// 		log.Printf("Error listing events: %v", err)
// 		return nil, err
// 	}
// 	defer func() {
// 		closeErr := cursor.Close(ctx)
// 		if closeErr != nil {
// 			log.Printf("Error closing cursor: %v", closeErr)
// 		}
// 	}()

// 	var daos []EventDAO
// 	if err = cursor.All(ctx, &daos); err != nil {
// 		log.Printf("Error decoding all events: %v", err)
// 		return nil, err
// 	}

// 	var events []Event
// 	for _, dao := range daos {
// 		events = append(events, *entities.EventDaoToDomain(&dao))
// 	}
// 	return events, nil
// }

// func (ed *mongoEventDAO) SoftDelete(ctx context.Context, ID string) (*entities.Event, error) {
// 	eventID, err := primitive.ObjectIDFromHex(ID)
// 	if err != nil {
// 		log.Printf("Error converting ID: %v", err)
// 		return nil, err
// 	}

// 	filter := bson.M{"_id": eventID}
// 	update := bson.M{
// 		"$set": bson.M{
// 			"deleted_at": time.Now(),
// 			"deleted":    true,
// 		},
// 	}

// 	result, err := ed.repository.DB().Collection("events").UpdateOne(ctx, filter, update)
// 	if err != nil {
// 		log.Printf("Error soft deleting event: %v", err)
// 		return nil, err
// 	}

// 	var updatedDAO *entities.EventDAO
// 	err = ed.repository.DB().Collection("events").FindOne(ctx, filter).Decode(&updatedDAO)
// 	if err != nil {
// 		log.Printf("Error decoding soft deleted event: %v", err)
// 		return nil, err
// 	}

// 	log.Printf("softdeleted %d events", result.ModifiedCount)

// 	return EventDaoToDomain(updatedDAO), nil
// }

// func (ed *mongoEventDAO) SoftUndelete(ctx context.Context, ID string) (*entities.Event, error) {
// 	eventID, err := primitive.ObjectIDFromHex(ID)
// 	if err != nil {
// 		log.Printf("Error converting ID: %v", err)
// 		return nil, err
// 	}

// 	filter := bson.M{"_id": eventID}
// 	update := bson.M{
// 		"$set": bson.M{
// 			"deleted_at": nil,
// 			"deleted":    false,
// 		},
// 	}

// 	result, err := ed.repository.DB().Collection("events").UpdateOne(ctx, filter, update)
// 	if err != nil {
// 		log.Printf("Error soft undeleting event: %v", err)
// 		return nil, err
// 	}

// 	var updatedDAO *entities.EventDAO
// 	err = ed.repository.DB().Collection("events").FindOne(ctx, filter).Decode(&updatedDAO)
// 	if err != nil {
// 		log.Printf("Error decoding soft undeleted event: %v", err)
// 		return nil, err
// 	}

// 	log.Printf("softundeleted %d events", result.ModifiedCount)

// 	return EventDaoToDomain(updatedDAO), nil
// }

// func (ed *mongoEventDAO) AddUserToEvent(ctx context.Context, eventID string, user *usr.User) (*entities.Event, error) {
// 	// Implementación pendiente
// 	return nil, fmt.Errorf("AddUserToEvent method not implemented")
// }

// // helpers
// func checkEventFields(event *entities.EventDAO) map[string]any {
// 	if event == nil {
// 		return nil
// 	}

// 	updateFields := make(map[string]any)

// 	if event.EventName != "" {
// 		updateFields["event_name"] = event.EventName
// 	}

// 	if event.Description != "" {
// 		updateFields["description"] = event.Description
// 	}

// 	if event.Date != "" {
// 		updateFields["date"] = event.Date
// 	}

// 	if event.Location != nil && (event.Location.Address != "" || event.Location.City != "" || event.Location.State != "" || event.Location.Country != "" || event.Location.PostalCode != "") {
// 		updateFields["location"] = event.Location
// 	}

// 	if len(event.Attendees) > 0 {
// 		updateFields["attendees"] = event.Attendees
// 	}

// 	if len(event.Planners) > 0 {
// 		updateFields["planners"] = event.Planners
// 	}

// 	if len(event.Tags) > 0 {
// 		updateFields["tags"] = event.Tags
// 	}

// 	return updateFields
// }
