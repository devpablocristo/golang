package person

import (
	"context"

	"github.com/devpablocristo/golang/sdk/services/person/internal/person/entities"
)

type SliceDB struct {
	db entities.Sdb
}

func NewSliceDB() *SliceDB {
	return &SliceDB{}
}

func (s *SliceDB) SavePerson(ctx context.Context, person *entities.Person) error {
	s.db = append(s.db, *person)
	return nil
}

// func (s *SliceDB) GetPerson(ctx context.Context, UUID string) (*Person, error) {
// 	for _, person := range s.sDB {
// 		if person.UUID == UUID {
// 			return &person, nil
// 		}
// 	}
// 	return nil, nil
// }

// func (s *SliceDB) ListPersons(ctx context.Context) map[string]*Person {
// 	mdb, err := slice2Map(ctx, s.sDB)
// 	if err != nil {
// 		return nil
// 	}
// 	return mdb
// }

// func (s *SliceDB) DeletePerson(ctx context.Context, UUID string) error {
// 	for i, person := range s.sDB {
// 		if person.UUID == UUID {
// 			s.sDB = append(s.sDB[:i], s.sDB[i+1:]...)
// 			return nil
// 		}
// 	}
// 	return nil
// }

// func (s *SliceDB) UpdatePerson(ctx context.Context, UUID string) error {
// 	return nil
// }

// func slice2Map(ctx context.Context, s []Person) (map[string]*Person, error) {
// 	mdb := make(map[string]*Person)
// 	for i := 0; i < len(s); i++ {
// 		mdb[s[i].UUID] = &s[i]
// 	}

// 	return mdb, nil
// }
