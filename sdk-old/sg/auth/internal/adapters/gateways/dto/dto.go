package dto

import (
	sdktypes "github.com/devpablocristo/golang/sdk/pkg/types"
)

type LoginResponse struct {
	Token string `json:"token"`
}

func LoginRequestToDomain(lr *sdktypes.LoginCredentials) *sdktypes.LoginCredentials {
	return &sdktypes.LoginCredentials{
		Username:     lr.Username,
		PasswordHash: lr.PasswordHash,
	}
}

func DomainToLoginResponse(lc *sdktypes.LoginCredentials) *sdktypes.LoginCredentials {
	return &sdktypes.LoginCredentials{
		Username:     lc.Username,
		PasswordHash: lc.PasswordHash,
	}
}
