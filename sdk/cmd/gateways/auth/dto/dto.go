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

func LoginRequestToDomain(lr *LoginRequest) *entities.AuthUser {
	return &entities.AuthUser{
		Username: lr.Username,
		Password: lr.Password,
	}
}

func DomainToLoginResponse(au *entities.AuthUser) *LoginRequest {
	return &LoginRequest{
		Username: au.Username,
		Password: au.Password,
	}
}
