package book

import (
	"errors"

	"github.com/devpablocristo/clean-pod-2/internal/platform/localdb"
	"github.com/google/uuid"
)

// extiende el repo en platform que es agnostico
type localRepo struct {
	db *localdb.LocalDB
}

func NewLocalRepo(db *localdb.LocalDB) *localRepo {
	return &localRepo{db: db}
}

func (l localRepo) Save(d *Book) error {
	if d.ID == "" {
		id := uuid.New()
		d.ID = id.String()
	}

	l.db.SaveItem(d.ID, d)
	return nil
}

func (l localRepo) Get(id string) (*Book, error) {
	item := l.db.GetItem(id)

	dev, ok := item.(*Book)
	if !ok {
		return nil, errors.New("cannot parse to developer")
	}

	return dev, nil
}

func (l localRepo) Delete(id string) error {
	l.db.DeleteItem(id)
	return nil
}

// func (l localRepo) SearchByStatus(status task.Status) ([]Developer, error) {
// 	var toRet []Book

// 	items := l.db.Dump()
// 	for _, item := range items {
// 		dev, ok := item.(*Developer)
// 		if !ok {
// 			continue
// 		}

// 		if dev.Task.Status == status {
// 			toRet = append(toRet, *dev)
// 		}
// 	}

// 	return toRet, nil
// }
