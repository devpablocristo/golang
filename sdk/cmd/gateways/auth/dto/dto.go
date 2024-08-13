package dto

import (
	"github.com/devpablocristo/golang/sdk/internal/core/auth/entities"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func LoginRequestToDomain(lr *LoginRequest) *entities.LogingCredentials {
	return &entities.LogingCredentials{
		Username: lr.Username,
		Password: lr.Password,
	}
}

func DomainToLoginResponse(au *entities.LogingCredentials) *LoginRequest {
	return &LoginRequest{
		Username: au.Username,
		Password: au.Password,
	}
}
