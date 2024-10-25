package dto

import (
	"time"

	entities "github.com/devpablocristo/golang/sdk/sg/users/internal/core/entities"
	personentities "github.com/devpablocristo/golang/sdk/sg/users/internal/person/core/entities"
)

// UserDto representa los datos para crear un usuario en la capa de aplicación
type UserDto struct {
	UUID        string
	UserType    string
	Person      *PersonDto
	Credentials CredentialsDto
	Roles       []RoleDto
	LoggedAt    *time.Time
}

// PersonDto representa los datos de una persona en la capa de aplicación
type PersonDto struct {
	Cuil        string
	Dni         *string
	FirstName   *string
	LastName    *string
	Nationality *string
	Email       string
	Phone       string
}

// CredentialsDto representa las credenciales de un usuario en la capa de aplicación
type CredentialsDto struct {
	Username string
	// Aquí podrías incluir PasswordHash u otro campo relacionado
}

// RoleDto representa los roles de un usuario en la capa de aplicación
type RoleDto struct {
	Name        string
	Permissions []PermissionDto
}

// PermissionDto representa los permisos de un rol en la capa de aplicación
type PermissionDto struct {
	Name        string
	Description string
}

func ToUser(dto *UserDto) *entities.User {
	roles := make([]entities.Role, len(dto.Roles))
	for i, roleDto := range dto.Roles {
		permissions := make([]entities.Permission, len(roleDto.Permissions))
		for j, permDto := range roleDto.Permissions {
			permissions[j] = entities.Permission{
				Name:        permDto.Name,
				Description: permDto.Description,
			}
		}
		roles[i] = entities.Role{
			Name:        roleDto.Name,
			Permissions: permissions,
		}
	}

	return &entities.User{
		Credentials: entities.Credentials{
			Username:     dto.Credentials.Username,
			PasswordHash: "", // Asumimos que el password hash se manejará en otra parte del flujo
		},
		Roles:     roles,
		CreatedAt: time.Now(),
		LoggedAt:  dto.LoggedAtOrDefault(),
	}
}

// LoggedAtOrDefault es un método auxiliar para usar un valor por defecto en caso de que LoggedAt sea nulo
func (dto UserDto) LoggedAtOrDefault() time.Time {
	if dto.LoggedAt != nil {
		return *dto.LoggedAt
	}
	return time.Now()
}

func ToPerson(dto *PersonDto) *personentities.Person {
	return &personentities.Person{
		Cuil:        dto.Cuil,
		Dni:         getStringValue(dto.Dni), // Para evitar punteros nulos
		FirstName:   getStringValue(dto.FirstName),
		LastName:    getStringValue(dto.LastName),
		Nationality: getStringValue(dto.Nationality),
		Email:       dto.Email,
		Phone:       dto.Phone,
	}
}

// Función auxiliar para obtener el valor de un puntero a string o devolver una cadena vacía si es nil
func getStringValue(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}
