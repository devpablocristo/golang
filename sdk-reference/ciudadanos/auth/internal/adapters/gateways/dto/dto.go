package dto

import (
	sdkjwt "github.com/devpablocristo/golang/sdk/pkg/jwt/v5"
	mware "github.com/devpablocristo/golang/sdk/pkg/middleware/gin"
)

type LoginResponse struct {
	Token string `json:"token"`
}

func LoginRequestToDomain(lr *mware.LoginRequest) *sdkjwt.LoginCredentials {
	return &sdkjwt.LoginCredentials{
		Username:     lr.Username,
		PasswordHash: lr.PasswordHash,
	}
}

func DomainToLoginResponse(lc *sdkjwt.LoginCredentials) *mware.LoginRequest {
	return &mware.LoginRequest{
		Username:     lc.Username,
		PasswordHash: lc.PasswordHash,
	}
}
