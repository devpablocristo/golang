package monitoring

import (
	"context"

	"github.com/devpablocristo/golang/sdk/internal/core/monitoring/ports"
)

type useCases struct {
	repository ports.Repository
}

// NewUseCases crea una nueva instancia de casos de uso de monitoreo.
func NewUseCases(r ports.Repository) ports.UseCases {
	return &useCases{
		repository: r,
	}
}

// CheckDbConn verifica la conexión a la base de datos usando el repositorio.
func (u *useCases) CheckDbConn(ctx context.Context) error {
	return u.repository.CheckDbConn(ctx)
}
