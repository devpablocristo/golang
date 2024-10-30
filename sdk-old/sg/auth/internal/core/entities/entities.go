package authent

import (
	"time"
)

type Session struct {
	UserUUID  string
	Token     string
	LoggedAt  time.Time
	ExpiresAt time.Time
}

// Auth representa la estructura de autenticación
type Auth struct {
	UserUUID string
	Session  Session
}
