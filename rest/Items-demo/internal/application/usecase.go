package application

import (
	entity "Items/internal/domain"
	"errors"
	"log"
)

type UseCase interface {
	AddItem(item entity.Item) (entity.Item, error)
}

type useCase struct {
	repo entity.Repository
}

func NewUseCaseService(repository entity.Repository) UseCase {
	return &useCase{repo: repository}
}

func (u *useCase) AddItem(item entity.Item) (entity.Item, error) {

	itemAdded, err := u.repo.AddItem(item)
	if err != nil {
		log.Fatalln("Error adding item")
		return itemAdded, errors.New("Error adding item")
	}

	return itemAdded, nil
}
