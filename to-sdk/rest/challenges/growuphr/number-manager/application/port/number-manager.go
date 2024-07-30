package port

import (
	"context"

	domain "github.com/devpablocristo/growuphr/number-manager/domain"
)

//go:generate mockgen -source=./number-manager.go -destination=../../../mocks/number-manager_mock.go -package=mocks
type NumberManager interface {
	AddReserveNumber(context.Context, *domain.ReservedNumber) error
	ReservedNumbers(context.Context) (map[string]*domain.ReservedNumber, error)
}
