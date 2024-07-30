package localdb

type LocalDB struct {
	storage map[string]interface{}
}

// simplmente un repo
func New() *LocalDB {
	return &LocalDB{storage: map[string]interface{}{}}
}

func (db LocalDB) SaveItem(id string, item interface{}) {
	db.storage[id] = item
}

func (db LocalDB) GetItem(id string) interface{} {
	item, ok := db.storage[id]
	if !ok {
		return nil
	}

	return item
}

func (db LocalDB) DeleteItem(id string) {
	delete(db.storage, id)
}

func (db LocalDB) Dump() []interface{} {
	var items []interface{}
	for _, value := range db.storage {
		items = append(items, value)
	}

	return items
}
