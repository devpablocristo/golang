package user

import (
	"context"
	"errors"
	"time"

	user "github.com/devpablocristo/golang/sdk/internal/core/user/entities"
	ports "github.com/devpablocristo/golang/sdk/internal/core/user/ports"
	"github.com/google/uuid"
)

// MapDbRepository es una implementación del repositorio usando un mapa en memoria
type MapDbRepository struct {
	db *user.InMemDB
}

// NewMapDbRepository crea un nuevo repositorio de usuarios en memoria
func NewMapDbRepository() ports.Repository {
	db := make(user.InMemDB)
	return &MapDbRepository{
		db: &db,
	}
}

// SaveUser guarda un nuevo usuario en el repositorio
func (r *MapDbRepository) SaveUser(ctx context.Context, usr *user.User) error {
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
func (r *MapDbRepository) GetUser(ctx context.Context, UUID string) (*user.User, error) {
	user, exists := (*r.db)[UUID]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// GetUserByUsername obtiene un usuario por su nombre de usuario
func (r *MapDbRepository) GetUserByUsername(ctx context.Context, username string) (*user.User, error) {
	for _, user := range *r.db {
		if user.Username == username {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

// DeleteUser elimina un usuario por su ID (UUID)
func (r *MapDbRepository) DeleteUser(ctx context.Context, UUID string) error {
	if _, exists := (*r.db)[UUID]; !exists {
		return errors.New("user not found")
	}
	delete(*r.db, UUID)
	return nil
}

// ListUsers lista todos los usuarios en el repositorio
func (r *MapDbRepository) ListUsers(ctx context.Context) (*user.InMemDB, error) {
	return r.db, nil
}

// UpdateUser actualiza la información de un usuario existente
func (r *MapDbRepository) UpdateUser(ctx context.Context, usr *user.User, UUID string) error {
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
