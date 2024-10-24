package dto

import (
	"errors"
	"time"

	entities "github.com/devpablocristo/golang/sdk/sg/users/internal/core/entities"
)

type UserDTO struct {
	UUID        string         `json:"uuid"` required
	PersonUUID  *string        `json:"person_uuid,omitempty"` 
	CompanyUUID *string        `json:"company_uuid,omitempty"`
	UserType    string         `json:"user_type"` 
	Credentials CredentialsDTO `json:"credentials"`
	Roles       []RoleDTO      `json:"roles"`
	LoggedAt    *time.Time     `json:"logged_at,omitempty"`
}

type RoleDTO struct {
	Name        string          `json:"name"`
	Permissions []PermissionDTO `json:"permissions"`
}

type PermissionDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CredentialsDTO struct {
	Username     string `json:"username"`
	PasswordHash string `json:"passwordhash"`
}

// ToPermissionDTO convierte una entidad de dominio Permission a un DTO PermissionDTO
func ToPermissionDTO(permission entities.Permission) PermissionDTO {
	return PermissionDTO{
		Name:        permission.Name,
		Description: permission.Description,
	}
}

// ToRoleDTO convierte una entidad de dominio Role a un DTO RoleDTO
func ToRoleDTO(role entities.Role) RoleDTO {
	permissions := make([]PermissionDTO, 0, len(role.Permissions))
	for _, perm := range role.Permissions {
		permissions = append(permissions, ToPermissionDTO(perm))
	}

	return RoleDTO{
		Name:        role.Name,
		Permissions: permissions,
	}
}

// ToUserDTO convierte una entidad de dominio User a un DTO UserDTO
func ToUserDTO(user *entities.User) *UserDTO {
	if user == nil {
		return nil
	}

	roles := make([]RoleDTO, 0, len(user.Roles))
	for _, role := range user.Roles {
		roles = append(roles, ToRoleDTO(role))
	}

	var loggedAt *time.Time
	if !user.LoggedAt.IsZero() {
		loggedAt = &user.LoggedAt
	}

	return &UserDTO{
		UUID:        user.UUID,
		PersonUUID:  user.PersonUUID,
		CompanyUUID: user.CompanyUUID,
		UserType:    string(user.UserType),
		Credentials: CredentialsDTO{
			Username: user.Credentials.Username,
		},
		Roles:    roles,
		LoggedAt: loggedAt,
	}
}

// FromPermissionDTO convierte un DTO PermissionDTO a una entidad de dominio Permission
func FromPermissionDTO(dtoPermission PermissionDTO) entities.Permission {
	return entities.Permission{
		Name:        dtoPermission.Name,
		Description: dtoPermission.Description,
	}
}

// FromRoleDTO convierte un DTO RoleDTO a una entidad de dominio Role
func FromRoleDTO(dtoRole RoleDTO) entities.Role {
	permissions := make([]entities.Permission, 0, len(dtoRole.Permissions))
	for _, permDTO := range dtoRole.Permissions {
		permissions = append(permissions, FromPermissionDTO(permDTO))
	}

	return entities.Role{
		Name:        dtoRole.Name,
		Permissions: permissions,
	}
}

// FromUserDTO convierte un DTO UserDTO a una entidad de dominio User
// Nota: Debido a que UserDTO no contiene campos como PasswordHash, CreatedAt, UpdatedAt, etc.,
// estos deben ser proporcionados externamente.
func FromUserDTO(dtoUser *UserDTO, passwordHash string, createdAt, updatedAt time.Time) (*entities.User, error) {
	if dtoUser == nil {
		return nil, errors.New("dtoUser is nil")
	}

	// Validar que solo uno de los UUIDs esté presente
	if (dtoUser.PersonUUID != nil && dtoUser.CompanyUUID != nil) || (dtoUser.PersonUUID == nil && dtoUser.CompanyUUID == nil) {
		return nil, errors.New("Debe proporcionar solo uno de person_uuid o company_uuid")
	}

	// Validar UserType
	var userType entities.UserType
	switch dtoUser.UserType {
	case "person":
		userType = entities.PersonType
	case "company":
		userType = entities.CompanyType
	default:
		return nil, errors.New("user_type inválido")
	}

	// Crear la entidad de dominio User
	user := &entities.User{
		UUID:        dtoUser.UUID,
		PersonUUID:  dtoUser.PersonUUID,
		CompanyUUID: dtoUser.CompanyUUID,
		UserType:    userType,
		Credentials: entities.Credentials{
			Username:     dtoUser.Credentials.Username,
			PasswordHash: passwordHash,
		},
		Roles:     []entities.Role{},
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	// Convertir Roles
	for _, roleDTO := range dtoUser.Roles {
		user.Roles = append(user.Roles, FromRoleDTO(roleDTO))
	}

	// Asignar LoggedAt si está presente
	if dtoUser.LoggedAt != nil {
		user.LoggedAt = *dtoUser.LoggedAt
	} else {
		user.LoggedAt = time.Time{} // Zero value
	}

	// DeletedAt no está presente en el DTO; se puede asignar nil o manejar según sea necesario
	user.DeletedAt = nil

	return user, nil
}
