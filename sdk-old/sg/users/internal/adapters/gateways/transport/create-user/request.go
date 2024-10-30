package tranport

import (
	dto "github.com/devpablocristo/golang/sdk/sg/users/internal/core/dto"
)

type User struct {
	Person      *Person     `json:"person"`
	Credentials Credentials `json:"credentials"`
	Roles       []Role      `json:"roles,omitempty"`
}

// Person representa la estructura de datos para una persona
type Person struct {
	Cuil        string  `json:"cuil" binding:"required"`
	Dni         *string `json:"dni,omitempty"`
	FirstName   *string `json:"first_name,omitempty"`
	LastName    *string `json:"last_name,omitempty"`
	Nationality *string `json:"nationality,omitempty"`
	Email       string  `json:"email" binding:"required,email"`
	Phone       string  `json:"phone" binding:"required"`
}

// Creds representa la estructura de datos de las credenciales para la API
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Role representa la estructura de datos del rol para la API
type Role struct {
	Name        string       `json:"name"`
	Permissions []Permission `json:"permissions"`
}

// Perm representa la estructura de datos del permiso para la API
type Permission struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func ToUserDto(req *User) *dto.UserDto {
	var personDto *dto.PersonDto
	if req.Person != nil {
		personDto = &dto.PersonDto{
			Cuil:        req.Person.Cuil,
			Dni:         req.Person.Dni,
			FirstName:   req.Person.FirstName,
			LastName:    req.Person.LastName,
			Nationality: req.Person.Nationality,
			Email:       req.Person.Email,
			Phone:       req.Person.Phone,
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
		Person: personDto,
		Credentials: dto.CredentialsDto{
			Username: req.Credentials.Username,
			Password: req.Credentials.Password,
		},
		Roles: rolesDto,
	}
}
