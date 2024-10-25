package tranport

import (
	"time"

	dto "github.com/devpablocristo/golang/sdk/sg/users/internal/core/dto"
)

type CreateUserUser struct {
	UserType         string            `json:"user_type" binding:"required,oneof=person company"`
	CreateUserPerson *CreateUserPerson `json:"person,omitempty"`
	Credentials      CreateUserCreds   `json:"credentials" binding:"required"`
	Roles            []CreateUserRole  `json:"roles,omitempty"`
	LoggedAt         *time.Time        `json:"logged_at,omitempty"`
}

// CreateUserPerson representa la estructura de datos para una persona
type CreateUserPerson struct {
	Cuil        string  `json:"cuil" binding:"required"`
	Dni         *string `json:"dni,omitempty"`
	FirstName   *string `json:"first_name,omitempty"`
	LastName    *string `json:"last_name,omitempty"`
	Nationality *string `json:"nationality,omitempty"`
	Email       string  `json:"email" binding:"required,email"`
	Phone       string  `json:"phone" binding:"required"`
}

// CreateUserCreds representa la estructura de datos de las credenciales para la API
type CreateUserCreds struct {
	Username string `json:"username" binding:"required"`
	// PasswordHash se excluye deliberadamente por razones de seguridad
}

// CreateUserRole representa la estructura de datos del rol para la API
type CreateUserRole struct {
	Name        string           `json:"name" binding:"required"`
	Permissions []CreateUserPerm `json:"permissions" binding:"required,dive,required"`
}

// CreateUserPerm representa la estructura de datos del permiso para la API
type CreateUserPerm struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func ToUserDto(req *CreateUserUser) *dto.UserDto {
	var personDto *dto.PersonDto
	if req.CreateUserPerson != nil {
		personDto = &dto.PersonDto{
			Cuil:        req.CreateUserPerson.Cuil,
			Dni:         req.CreateUserPerson.Dni,
			FirstName:   req.CreateUserPerson.FirstName,
			LastName:    req.CreateUserPerson.LastName,
			Nationality: req.CreateUserPerson.Nationality,
			Email:       req.CreateUserPerson.Email,
			Phone:       req.CreateUserPerson.Phone,
		}
	}

	// Mapeamos los roles y permisos
	rolesDto := make([]dto.RoleDto, len(req.Roles))
	for i, role := range req.Roles {
		permissionsDto := make([]dto.PermissionDto, len(role.Permissions))
		for j, perm := range role.Permissions {
			permissionsDto[j] = dto.PermissionDto{
				Name:        perm.Name,
				Description: perm.Description,
			}
		}

		rolesDto[i] = dto.RoleDto{
			Name:        role.Name,
			Permissions: permissionsDto,
		}
	}

	return &dto.UserDto{
		UserType: req.UserType,
		Person:   personDto,
		Credentials: dto.CredentialsDto{
			Username: req.Credentials.Username,
		},
		Roles:    rolesDto,
		LoggedAt: req.LoggedAt,
	}
}
