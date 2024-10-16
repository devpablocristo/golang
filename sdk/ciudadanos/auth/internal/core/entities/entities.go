package authent

import (
	"time"

	sdkjwt "github.com/devpablocristo/golang/sdk/pkg/jwt/v5"
)

type Session struct {
	UserUUID  string
	Token     sdkjwt.Token
	LoggedAt  time.Time
	ExpiresAt time.Time
}

// Auth representa la estructura de autenticación
type Auth struct {
	UserUUID string
	Session  Session
}
