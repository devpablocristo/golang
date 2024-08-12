package auth

import (
	"time"
)

// Token representa un token de acceso o refresh, necesario para la autenticación y autorización.
type Token struct {
	AccessToken  string
	RefreshToken string
	ExpiresAt    time.Time
}

// Role define los roles que un usuario puede tener dentro del sistema.
type Role struct {
	Name        string
	Permissions []Permission
}

// Permission especifica los permisos asociados a roles y usuarios.
type Permission struct {
	Name        string
	Description string
}

// Credential representa las credenciales de acceso, como un correo electrónico y una contraseña.
type Credential struct {
	UserUUID string // Referencia al UUID del usuario en el microservicio de `users`
	Password string
}

// Session representa una sesión de usuario activa.
type Session struct {
	UserUUID  string
	Token     Token
	LoggedAt  time.Time
	ExpiresAt time.Time
}

// Auth es la estructura principal de autenticación que relaciona credenciales, roles y sesiones.
type Auth struct {
	UserUUID   string // Referencia al usuario en el microservicio de `users`
	Credential Credential
	Role       Role
	Session    Session
}
