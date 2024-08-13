package user

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/devpablocristo/golang/sdk/internal/core/user/coreports"
	"github.com/devpablocristo/golang/sdk/internal/core/user/entities"
)

// mapDbRepository es una implementación del repositorio usando un mapa en memoria
type mapDbRepository struct {
	db *entities.InMemDB
}

// NewMapDbRepository crea un nuevo repositorio de usuarios en memoria
func NewMapDbRepository() coreports.Repository {
	db := make(entities.InMemDB)
	return &mapDbRepository{
		db: &db,
	}
}

// SaveUser guarda un nuevo usuario en el repositorio
func (r *mapDbRepository) SaveUser(ctx context.Context, usr *entities.User) error {
	if usr.Username == "" {
		return errors.New("username is required")
	}

	// Generar un nuevo UUID para el usuario
	usr.UUID = uuid.New().String()
	usr.CreatedAt = time.Now()

	(*r.db)[usr.UUID] = usr
	return nil
}

// GetUser obtiene un usuario por su ID (UUID)
func (r *mapDbRepository) GetUser(ctx context.Context, UUID string) (*entities.User, error) {
	user, exists := (*r.db)[UUID]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// GetUserByUsername obtiene un usuario por su nombre de usuario
func (r *mapDbRepository) GetUserByUsername(ctx context.Context, username string) (*entities.User, error) {
	for _, user := range *r.db {
		if user.Username == username {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

// DeleteUser elimina un usuario por su ID (UUID)
func (r *mapDbRepository) DeleteUser(ctx context.Context, UUID string) error {
	if _, exists := (*r.db)[UUID]; !exists {
		return errors.New("user not found")
	}
	delete(*r.db, UUID)
	return nil
}

// ListUsers lista todos los usuarios en el repositorio
func (r *mapDbRepository) ListUsers(ctx context.Context) (*entities.InMemDB, error) {
	return r.db, nil
}

// UpdateUser actualiza la información de un usuario existente
func (r *mapDbRepository) UpdateUser(ctx context.Context, usr *entities.User, UUID string) error {
	existingUser, exists := (*r.db)[UUID]
	if !exists {
		return errors.New("user not found")
	}

	if usr.Username != "" {
		existingUser.Username = usr.Username
	}
	if usr.PasswordHash != "" {
		existingUser.PasswordHash = usr.PasswordHash
	}
	// Mantener la fecha de creación original
	usr.CreatedAt = existingUser.CreatedAt

	(*r.db)[UUID] = existingUser
	return nil
}
