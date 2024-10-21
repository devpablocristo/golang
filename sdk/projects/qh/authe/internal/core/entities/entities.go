package authent

import (
	"time"
)

// Token representa un token de autenticación JWT
type Token struct {
	AccessToken  string
	RefreshToken string
	ExpiresAt    time.Time
}

// Session representa una sesión de usuario
type Session struct {
	UserUUID  string
	Token     Token
	LoggedAt  time.Time
	ExpiresAt time.Time
}

// Auth representa la estructura de autenticación
type Auth struct {
	UserUUID string
	Session  Session
}

// LoginCredentials representa las credenciales de inicio de sesión del usuario
type LoginCredentials struct {
	Username     string
	PasswordHash string
}
