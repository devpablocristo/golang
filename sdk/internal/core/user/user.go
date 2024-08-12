package user

import (
	"time"
)

// User representa a un usuario en el sistema con su información básica.
type User struct {
	UUID          string     // Identificador único del usuario
	Username      string     // Nombre de usuario único
	Email         string     // Correo electrónico del usuario
	PasswordHash  string     // Hash de la contraseña para mayor seguridad
	AuthUUID      string     // UUID relacionado con el servicio de autenticación, si es necesario
	PersonUUID    string     // UUID que podría apuntar a una entidad de perfil personal más detallada
	Qualification int        `validate:"gte=1,lte=10"` // Calificación del usuario, validada entre 1 y 10
	CreatedAt     time.Time  // Fecha de creación del usuario
	LoggedAt      time.Time  // Última vez que el usuario inició sesión
	UpdatedAt     time.Time  // Fecha de la última actualización de la información del usuario
	DeletedAt     *time.Time // Fecha de eliminación del usuario (si aplica, usando puntero para permitir null)
}

// InMemDB simula una base de datos en memoria para almacenar usuarios.
// La clave es el UUID del usuario.
type InMemDB map[string]*User
