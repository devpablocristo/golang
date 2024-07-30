package port

import (
	"context"

	domain "github.com/devpablocristo/growuphr/number-manager/domain"
)

//go:generate mockgen -source=./storage.go -destination=../../../mocks/storage_mock.go -package=mocks
type Storage interface {
	Create(context.Context, *domain.ReservedNumber) error
	Read(context.Context, string) (*domain.ReservedNumber, error)
	List(context.Context) map[string]*domain.ReservedNumber
	Delete(context.Context, string) error
	CheckForUsername(context.Context, string) (*domain.ReservedNumber, bool)
	CheckForNumber(context.Context, int) (*domain.ReservedNumber, bool)
}
