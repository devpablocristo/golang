package port

import (
	"context"

	"github.com/devpablocristo/nanlabs/domain"
)

type Storage interface {
	SaveTask(context.Context, *domain.Task) error
	GetTask(context.Context, string) (*domain.Task, error)
	ListTasks(context.Context) map[string]*domain.Task
	DeleteTask(context.Context, string) error
	UpdateTask(context.Context, string) error
}
