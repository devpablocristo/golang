package ctypes

import (
	domain "github.com/devpablocristo/qh/internal/user-manager/user"
)

type UserDTO struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

func UserDtoToDomain(u *UserDTO) *domain.User {
	return &domain.User{
		ID:   u.ID,
		Type: u.Type,
	}
}
