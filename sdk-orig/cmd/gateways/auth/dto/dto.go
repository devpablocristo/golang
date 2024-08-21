package dto

import (
	"github.com/devpablocristo/golang/sdk/internal/core/auth/entities"

	mdw "github.com/devpablocristo/golang/sdk/pkg/middleware/gin"
)

type LoginResponse struct {
	Token string `json:"token"`
}

func LoginRequestToDomain(lr *mdw.LoginRequest) *entities.LoginCredentials {
	return &entities.LoginCredentials{
		Username:     lr.Username,
		PasswordHash: lr.PasswordHash,
	}
}

func DomainToLoginResponse(lc *entities.LoginCredentials) *mdw.LoginRequest {
	return &mdw.LoginRequest{
		Username:     lc.Username,
		PasswordHash: lc.PasswordHash,
	}
}
