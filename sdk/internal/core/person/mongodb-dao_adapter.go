package person

import (
	"context"

	mdb "github.com/devpablocristo/qh/events/pkg/mongodb/mongo-driver"
)

type mongoEventDAO struct {
	inst mdb.MongoDBClientPort
}

func NewMongoEventDAO(inst mdb.MongoDBClientPort) DAOPort {
	return &mongoEventDAO{
		inst: inst,
	}
}

func (ed *mongoEventDAO) Create(ctx context.Context, person *Person) error {
	return nil
}

// func (ed *mongoEventDAO) FindByID(ctx context.Context, ID string) (*Event, error) {
// 	eventID, err := primitive.ObjectIDFromHex(ID)
// 	if err != nil {
// 		log.Printf("Error converting ID: %v", err)
// 		return nil, err
// 	}

// 	filter := bson.M{"_id": eventID}
// 	dao := &EventDAO{}
// 	if err := ed.inst.DB().Collection("events").FindOne(ctx, filter).Decode(dao); err != nil {
// 		log.Printf("Error decoding event: %v", err)
// 		return nil, err
// 	}
// 	return EventDaoToDomain(dao), nil
// }

// func (ed *mongoEventDAO) Update(ctx context.Context, event *Event, ID string) (*Event, error) {
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

// 	result, err := ed.inst.DB().Collection("events").UpdateOne(ctx, filter, update)
// 	if err != nil {
// 		log.Printf("Error updating event: %v", err)
// 		return nil, err
// 	}

// 	if result.MatchedCount == 0 {
// 		log.Printf("Error: no event found with ID: %s", eventID.Hex())
// 		return nil, fmt.Errorf("no event found with ID: %s", eventID.Hex())
// 	}

// 	var updatedDAO EventDAO
// 	if err := ed.inst.DB().Collection("events").FindOne(ctx, filter).Decode(&updatedDAO); err != nil {
// 		log.Printf("Error decoding updated event: %v", err)
// 		return nil, err
// 	}

// 	log.Printf("updated %d events", result.ModifiedCount)

// 	return EventDaoToDomain(&updatedDAO), nil
// }

// func (ed *mongoEventDAO) HardDelete(ctx context.Context, ID string) (*Event, error) {
// 	eventID, err := primitive.ObjectIDFromHex(ID)
// 	if err != nil {
// 		log.Printf("Error converting ID: %v", err)
// 		return nil, err
// 	}

// 	filter := bson.M{"_id": eventID}

// 	deletedDAO := &EventDAO{}
// 	if err := ed.inst.DB().Collection("events").FindOne(ctx, filter).Decode(deletedDAO); err != nil {
// 		log.Printf("Error decoding deleted event: %v", err)
// 		return nil, err
// 	}

// 	result, err := ed.inst.DB().Collection("events").DeleteOne(ctx, filter)
// 	if err != nil {
// 		log.Printf("Error deleting event: %v", err)
// 		return nil, err
// 	}

// 	log.Printf("deleted %d events", result.DeletedCount)
// 	return EventDaoToDomain(deletedDAO), nil
// }

// func (ed *mongoEventDAO) List(ctx context.Context) ([]Event, error) {
// 	cursor, err := ed.inst.DB().Collection("events").Find(ctx, bson.M{})
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
// 		events = append(events, *EventDaoToDomain(&dao))
// 	}
// 	return events, nil
// }

// func (ed *mongoEventDAO) SoftDelete(ctx context.Context, ID string) (*Event, error) {
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

// 	result, err := ed.inst.DB().Collection("events").UpdateOne(ctx, filter, update)
// 	if err != nil {
// 		log.Printf("Error soft deleting event: %v", err)
// 		return nil, err
// 	}

// 	var updatedDAO *EventDAO
// 	err = ed.inst.DB().Collection("events").FindOne(ctx, filter).Decode(&updatedDAO)
// 	if err != nil {
// 		log.Printf("Error decoding soft deleted event: %v", err)
// 		return nil, err
// 	}

// 	log.Printf("softdeleted %d events", result.ModifiedCount)

// 	return EventDaoToDomain(updatedDAO), nil
// }

// func (ed *mongoEventDAO) SoftUndelete(ctx context.Context, ID string) (*Event, error) {
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

// 	result, err := ed.inst.DB().Collection("events").UpdateOne(ctx, filter, update)
// 	if err != nil {
// 		log.Printf("Error soft undeleting event: %v", err)
// 		return nil, err
// 	}

// 	var updatedDAO *EventDAO
// 	err = ed.inst.DB().Collection("events").FindOne(ctx, filter).Decode(&updatedDAO)
// 	if err != nil {
// 		log.Printf("Error decoding soft undeleted event: %v", err)
// 		return nil, err
// 	}

// 	log.Printf("softundeleted %d events", result.ModifiedCount)

// 	return EventDaoToDomain(updatedDAO), nil
// }

// func (ed *mongoEventDAO) AddUserToEvent(ctx context.Context, eventID string, user *usr.User) (*Event, error) {
// 	// Implementación pendiente
// 	return nil, fmt.Errorf("AddUserToEvent method not implemented")
// }

// // helpers
// func checkEventFields(event *EventDAO) map[string]interface{} {
// 	if event == nil {
// 		return nil
// 	}

// 	updateFields := make(map[string]interface{})

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
