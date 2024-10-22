package ports

import (
	"context"

	entities "github.com/devpablocristo/golang/sdk/sg/users/internal/company/core/entities"
)

type Repository interface {
	Create(context.Context, *entities.Company) error
	FindByID(context.Context, string) (*entities.Company, error)
	Update(context.Context, *entities.Company) error
	SoftDelete(context.Context, string) error
	FindByCuit(context.Context, string) (*entities.Company, error)
}
