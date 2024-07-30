package port

import (
	"context"

	"github.com/devpablocristo/nanlabs/domain"
)

//go:generate mockgen -source=./service.go -destination=../../mocks/service_mock.go -package=mocks
type Service interface {
	CreateCard(context.Context, *domain.Task) error
	GetTasks(context.Context) (map[string]*domain.Task, error)
	GetTask(context.Context, string) (*domain.Task, error)
	CreateTask(context.Context, *domain.Task) (*domain.Task, error)
	UpdateTask(context.Context, string) error
	DeleteTask(context.Context, string) error
}
