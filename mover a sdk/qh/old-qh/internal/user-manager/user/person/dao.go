package person

import (
	"context"
)

// type EventDAO struct {
// 	ID          primitive.ObjectID `bson:"_id,omitempty"`
// 	CreatedAt   time.Time          `bson:"created_at"`
// 	UpdatedAt   time.Time          `bson:"updated_at"`
// 	DeletedAt   time.Time          `bson:"deleted_at"`
// 	Deleted     bool               `bson:"deleted"`
// 	Version     int                `bson:"version"`
// 	EventName   string             `bson:"event_name"`
// 	Description string             `bson:"description"`
// 	Date        string             `bson:"date"`
// 	Location    *LocationDAO       `bson:"location"`
// 	Organizer   *OrganizerDAO      `bson:"organizer"`
// 	Attendees   []AttendeeDAO      `bson:"attendees"`
// 	Planners    []PlannerDAO       `bson:"planners"`
// 	Tags        []string           `bson:"tags"`
// }

// type LocationDAO struct {
// 	Address    string `bson:"address"`
// 	City       string `bson:"city"`
// 	State      string `bson:"state"`
// 	Country    string `bson:"country"`
// 	PostalCode string `bson:"postal_code"`
// }

// type OrganizerDAO struct {
// 	Name  string `bson:"name"`
// 	Email string `bson:"email"`
// 	Phone string `bson:"phone"`
// }

// type AttendeeDAO struct {
// 	Person *PersonDAO `bson:"person"`
// }

// type PlannerDAO struct {
// 	Person *PersonDAO `bson:"person"`
// }

// type PersonDAO struct {
// 	Name  string `bson:"name"`
// 	Email string `bson:"email"`
// }

// func personDAOToPerson(dao *PersonDAO) *.Person {
// 	return &.Person{
// 		Name:  dao.Name,
// 		Email: dao.Email,
// 	}
// }

// func personToPersonDAO(person *.Person) *PersonDAO {
// 	return &PersonDAO{
// 		Name:  person.Name,
// 		Email: person.Email,
// 	}
// }

// func daoToDomain(dao *EventDAO) *.Person {
// 	if dao == nil {
// 		return nil
// 	}

// 	attendees := make([].Attendee, len(dao.Attendees))
// 	for i, daoAttendee := range dao.Attendees {
// 		attendees[i] = .Attendee{
// 			Person: *personDAOToPerson(daoAttendee.Person),
// 		}
// 	}

// 	planners := make([].Planner, len(dao.Planners))
// 	for i, daoPlanner := range dao.Planners {
// 		planners[i] = .Planner{
// 			Person: *personDAOToPerson(daoPlanner.Person),
// 		}
// 	}

// 	return &.Person{
// 		EventName:   dao.EventName,
// 		Description: dao.Description,
// 		Date:        dao.Date,
// 		Location: .Location{
// 			Address:    dao.Location.Address,
// 			City:       dao.Location.City,
// 			State:      dao.Location.State,
// 			Country:    dao.Location.Country,
// 			PostalCode: dao.Location.PostalCode,
// 		},
// 		Organizer: .Organizer{
// 			Name:  dao.Organizer.Name,
// 			Email: dao.Organizer.Email,
// 			Phone: dao.Organizer.Phone,
// 		},
// 		Attendees: attendees,
// 		Planners:  planners,
// 		Tags:      dao.Tags,
// 	}
// }

// func domainToDAO( *.Person) *EventDAO {
// 	if  == nil {
// 		return nil
// 	}

// 	daoAttendees := make([]AttendeeDAO, len(.Attendees))
// 	for i, attendee := range .Attendees {
// 		daoAttendees[i] = AttendeeDAO{
// 			Person: personToPersonDAO(&attendee.Person),
// 		}
// 	}

// 	daoPlanners := make([]PlannerDAO, len(.Planners))
// 	for i, planner := range .Planners {
// 		daoPlanners[i] = PlannerDAO{
// 			Person: personToPersonDAO(&planner.Person),
// 		}
// 	}

// 	return &EventDAO{
// 		EventName:   .PersonName,
// 		Description: .Description,
// 		Date:        .Date,
// 		Location: &LocationDAO{
// 			Address:    .Location.Address,
// 			City:       .Location.City,
// 			State:      .Location.State,
// 			Country:    .Location.Country,
// 			PostalCode: .Location.PostalCode,
// 		},
// 		Organizer: &OrganizerDAO{
// 			Name:  .Organizer.Name,
// 			Email: .Organizer.Email,
// 			Phone: .Organizer.Phone,
// 		},
// 		Attendees: daoAttendees,
// 		Planners:  daoPlanners,
// 		Tags:      .Tags,
// 	}
// }

type PortMongoPersonDAO interface {
	//FindByID(context.Context, string) (*.Person, error)
	Create(context.Context, *Person) (*Person, error)
	// Update(context.Context, *.Person, string) (*.Person, error)
	// HardDelete(context.Context, string) (*.Person, error)
	// List(context.Context) ([].Person, error)
	// SoftDelete(context.Context, string) (*.Person, error)
	// SoftUndelete(context.Context, string) (*.Person, error)
}

type mongoPersonDAO struct {
	mongoService MongoDBServicePort
}

func NewmongoPersonDAO(ms MongoDBServicePort) PortMongoPersonDAO {
	return &mongoPersonDAO{
		mongoService: ms,
	}
}

// func (evDao *mongoPersonDAO) FindByID(ctx context.Context, ID string) (*.Person, error) {
// 	eventID, err := primitive.ObjectIDFromHex(ID)
// 	if err != nil {
// 		log.Printf("Error converting ID: %v", err)
// 		return nil, err
// 	}

// 	filter := bson.M{"_id": eventID}
// 	dao := &EventDAO{}
// 	if err := evDao.mongoService.GetCollection(ctx).FindOne(ctx, filter).Decode(dao); err != nil {
// 		log.Printf("Error decoding event: %v", err)
// 		return nil, err
// 	}
// 	return daoToDomain(dao), nil
// }

func (evDao *mongoPersonDAO) Create(ctx context.Context, event *Person) (*Person, error) {
	// dao := domainToDAO(event)
	// dao.CreatedAt = time.Now()
	// result, err := evDao.mongoService.GetCollection(ctx).InsertOne(ctx, dao)
	// if err != nil {
	// 	log.Printf("Error inserting event: %v", err)
	// 	return nil, err
	// }
	// dao.ID = result.InsertedID.(primitive.ObjectID)
	// return daoToDomain(dao), nil
	return nil, nil
}

// func (evDao *mongoPersonDAO) Update(ctx context.Context, event *.Person, ID string) (*.Person, error) {
// 	eventID, err := primitive.ObjectIDFromHex(ID)
// 	if err != nil {
// 		log.Printf("Error converting ID: %v", err)
// 		return nil, err
// 	}

// 	dao := domainToDAO(event)

// 	updateFields := checkEventFields(dao)

// 	if updateFields == nil {
// 		log.Printf("Error: event is nil")
// 		return nil, errors.New("event is nil")
// 	}

// 	if len(updateFields) == 0 {
// 		log.Printf("Error: no fields to update")
// 		return nil, errors.New("no fields to update")
// 	}

// 	updateFields["updated_at"] = time.Now()

// 	filter := bson.M{"_id": eventID}
// 	update := bson.M{"$set": updateFields}

// 	result, err := evDao.mongoService.GetCollection(ctx).UpdateOne(ctx, filter, update)
// 	if err != nil {
// 		log.Printf("Error updating event: %v", err)
// 		return nil, err
// 	}

// 	if result.MatchedCount == 0 {
// 		log.Printf("Error: no event found with ID: %s", eventID.Hex())
// 		return nil, fmt.Errorf("no event found with ID: %s", eventID.Hex())
// 	}

// 	var updatedDAO EventDAO
// 	if err := evDao.mongoService.GetCollection(ctx).FindOne(ctx, filter).Decode(&updatedDAO); err != nil {
// 		log.Printf("Error decoding updated event: %v", err)
// 		return nil, err
// 	}

// 	log.Printf("updated %d events", result.ModifiedCount)

// 	return daoToDomain(&updatedDAO), nil
// }

// func (evDao *mongoPersonDAO) HardDelete(ctx context.Context, ID string) (*.Person, error) {
// 	eventID, err := primitive.ObjectIDFromHex(ID)
// 	if err != nil {
// 		log.Printf("Error converting ID: %v", err)
// 		return nil, err
// 	}

// 	filter := bson.M{"_id": eventID}

// 	deletedDAO := &EventDAO{}
// 	if err := evDao.mongoService.GetCollection(ctx).FindOne(ctx, filter).Decode(deletedDAO); err != nil {
// 		log.Printf("Error decoding deleted event: %v", err)
// 		return nil, err
// 	}

// 	result, err := evDao.mongoService.GetCollection(ctx).DeleteOne(ctx, filter)
// 	if err != nil {
// 		log.Printf("Error deleting event: %v", err)
// 		return nil, err
// 	}

// 	log.Printf("deleted %d events", result.DeletedCount)

// 	return daoToDomain(deletedDAO), nil
// }

// func (evDao *mongoPersonDAO) List(ctx context.Context) ([].Person, error) {
// 	cursor, err := evDao.mongoService.GetCollection(ctx).Find(ctx, bson.M{})
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

// 	var events [].Person
// 	for _, dao := range daos {
// 		events = append(events, *daoToDomain(&dao))
// 	}
// 	return events, nil
// }

// func (evDao *mongoPersonDAO) SoftDelete(ctx context.Context, ID string) (*.Person, error) {
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

// 	result, err := evDao.mongoService.GetCollection(ctx).UpdateOne(ctx, filter, update)
// 	if err != nil {
// 		log.Printf("Error soft deleting event: %v", err)
// 		return nil, err
// 	}

// 	var updatedDAO *EventDAO
// 	err = evDao.mongoService.GetCollection(ctx).FindOne(ctx, filter).Decode(&updatedDAO)
// 	if err != nil {
// 		log.Printf("Error decoding soft deleted event: %v", err)
// 		return nil, err
// 	}

// 	log.Printf("softdeleted %d events", result.ModifiedCount)

// 	return daoToDomain(updatedDAO), nil
// }

// func (evDao *mongoPersonDAO) SoftUndelete(ctx context.Context, ID string) (*.Person, error) {
// 	eventID, err := primitive.ObjectIDFromHex(ID)
// 	if err != nil {
// 		log.Printf("Error converting ID: %v", err)
// 		return nil, err
// 	}

// 	filter := bson.M{"_id": eventID}
// 	update := bson.M{
// 		"$set": bson.M{
// 			"deleted_at": time.Now(),
// 			"deleted":    false,
// 		},
// 	}

// 	result, err := evDao.mongoService.GetCollection(ctx).UpdateOne(ctx, filter, update)
// 	if err != nil {
// 		log.Printf("Error soft deleting event: %v", err)
// 		return nil, err
// 	}

// 	var updatedDAO *EventDAO
// 	err = evDao.mongoService.GetCollection(ctx).FindOne(ctx, filter).Decode(&updatedDAO)
// 	if err != nil {
// 		log.Printf("Error decoding soft deleted event: %v", err)
// 		return nil, err
// 	}

// 	log.Printf("softdeleted %d events", result.ModifiedCount)

// 	return daoToDomain(updatedDAO), nil
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

// 	if event.Location.Address != "" || event.Location.City != "" || event.Location.State != "" || event.Location.Country != "" || event.Location.PostalCode != "" {
// 		updateFields["location"] = event.Location
// 	}

// 	if event.Organizer.Name != "" || event.Organizer.Email != "" || event.Organizer.Phone != "" {
// 		updateFields["organizer"] = event.Organizer
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