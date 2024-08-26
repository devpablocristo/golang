package coreauthentities

import (
	"time"
)

type Token struct {
	AccessToken  string
	RefreshToken string
	ExpiresAt    time.Time
}

type Session struct {
	UserUUID  string
	Token     Token
	LoggedAt  time.Time
	ExpiresAt time.Time
}

type Auth struct {
	UserUUID string
	Session  Session
}

type LoginCredentials struct {
	Username     string
	PasswordHash string
}
