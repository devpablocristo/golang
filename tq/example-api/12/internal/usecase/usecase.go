package usecase

import (
	"fmt"
	"log"

	entity "items/internal/entity"
)

const (
	noStock = iota
	inStock
)

const (
	activeStatus   = "ACTIVE"
	inactiveStatus = "INACTIVE"
)

// Usecases
// Esta es la inferface por donde se comunicara usescases con demas capas
type ItemUsecasePort interface {
	SaveItem(*entity.Item) (*entity.Item, error)
	GetAllItems() (entity.MapRepo, error)
	GetItemByID(entity.ID) (*entity.Item, error)
}

// el tipo de usecase es del tipo interface de inmemory
type ItemUsecase struct {
	inmemory entity.ItemRepositoryPort
	mysql    entity.ItemRepositoryPort
}

// como parametro de salida se usar la interface de usecase
func NewItemUsecase(i entity.ItemRepositoryPort, m entity.ItemRepositoryPort) ItemUsecasePort {
	return &ItemUsecase{
		inmemory: i,
		mysql:    m,
	}
}

// inmemory
// func (u *ItemUsecase) SaveItem(item *entity.Item) (*entity.Item, error) {

// 	exist, err := u.inmemory.CheckItemByCode(item.Code)
// 	if exist {
// 		log.Printf("codes must be unique %v: %s", err.Error(), item.Code)
// 		return nil, fmt.Errorf("codes must be unique %v: %s", err.Error(), item.Code)
// 	}

// 	item.Status = inactiveStatus
// 	if item.Stock > noStock {
// 		item.Status = activeStatus
// 	}

// 	savedItem, err := u.inmemory.SaveItem(item)
// 	if err != nil {
// 		log.Printf("error saving entity.Item: %v", err)
// 		return nil, fmt.Errorf("error saving entity.Item: %v", err)
// 	}

// 	return savedItem, nil
// }

// inmemory
// func (u *ItemUsecase) GetAllItems() (entity.MapRepo, error) {
// 	items, err := u.inmemory.GetAllItems()
// 	if err != nil {
// 		return nil, fmt.Errorf("error in inmemory: %w", err)
// 	}

// 	if len(items) == 0 {
// 		return nil, errItemNotFound
// 	}

// 	return items, nil
// }

// inmemory
// func (u *ItemUsecase) GetItemByID(id entity.ID) (*entity.Item, error) {
// 	item, err := u.inmemory.GetItemByID(id)
// 	if err != nil {
// 		return nil, errItemNotFound
// 	}
// 	return item, nil
// }

// mysql
func (u *ItemUsecase) SaveItem(item *entity.Item) (*entity.Item, error) {
	exist, err := u.mysql.CheckItemByCode(item.Code)
	if exist {
		log.Printf("codes must be unique %v: %s", err.Error(), item.Code)
		return nil, fmt.Errorf("codes must be unique %v: %s", err.Error(), item.Code)
	}

	item.Status = inactiveStatus
	if item.Stock > noStock {
		item.Status = activeStatus
	}

	savedItem, err := u.mysql.SaveItem(item)
	if err != nil {
		log.Printf("error saving entity.Item: %v", err)
		return nil, fmt.Errorf("error saving entity.Item: %v", err)
	}

	return savedItem, nil
}

// mysql
func (u *ItemUsecase) GetAllItems() (entity.MapRepo, error) {
	items, err := u.mysql.GetAllItems()
	if err != nil {
		return nil, fmt.Errorf("error in mysql: %w", err)
	}

	if len(items) == 0 {
		return nil, errItemNotFound
	}

	return items, nil
}

// mysql
func (u *ItemUsecase) GetItemByID(id entity.ID) (*entity.Item, error) {
	item, err := u.mysql.GetItemByID(id)
	if err != nil {
		return nil, errItemNotFound
	}
	return item, nil
}
