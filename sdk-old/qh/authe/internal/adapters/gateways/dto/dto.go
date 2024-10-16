package dto

import (
	mware "github.com/devpablocristo/golang/sdk/pkg/middleware/gin"
	entities "github.com/devpablocristo/golang/sdk/qh/authe/internal/core/entities"
)

type LoginResponse struct {
	Token string `json:"token"`
}

func LoginRequestToDomain(lr *mware.LoginRequest) *entities.LoginCredentials {
	return &entities.LoginCredentials{
		Username:     lr.Username,
		PasswordHash: lr.PasswordHash,
	}
}

func DomainToLoginResponse(lc *entities.LoginCredentials) *mware.LoginRequest {
	return &mware.LoginRequest{
		Username:     lc.Username,
		PasswordHash: lc.PasswordHash,
	}
}
