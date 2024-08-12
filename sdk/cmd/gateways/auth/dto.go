package auth

import (
	"time"

	"github.com/devpablocristo/golang/sdk/internal/core/user"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func toDomain(loginReq *LoginRequest) *user.User {
	return &user.User{
		Username: loginReq.Username,
		Password: loginReq.Password,
		LoggedAt: time.Now(),
	}
}
