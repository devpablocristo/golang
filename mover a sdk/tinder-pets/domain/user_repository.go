package domain

type UserRepository interface {
	Create(user *User) error
	Update(user *User) error
	Delete(id string) error
	GetByID(id string) (*User, error)
	GetAll() ([]*User, error)
}
