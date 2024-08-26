package entities

import (
	"time"
)

type User struct {
	UUID          string
	Email         string
	Credentials   Credentials
	PersonUUID    string
	Qualification int `validate:"gte=1,lte=10"` // Calificación del usuario, validada entre 1 y 10
	Roles         []Role
	CreatedAt     time.Time  // Fecha de creación del usuario
	LoggedAt      time.Time  // Última vez que el usuario inició sesión
	UpdatedAt     time.Time  // Fecha de la última actualización de la información del usuario
	DeletedAt     *time.Time // Fecha de eliminación del usuario (si aplica, usando puntero para permitir null)
}

type Role struct {
	Name        string
	Permissions []Permission
}

type Permission struct {
	Name        string
	Description string
}

type Credentials struct {
	Username     string
	PasswordHash string
}

type InMemDB map[string]*User
