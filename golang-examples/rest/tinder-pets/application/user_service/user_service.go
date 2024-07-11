package application

import (
	"github.com/luka385/tinder-pets/domain"
	"github.com/luka385/tinder-pets/infrastructure/mongodb"
)

type UserService struct {
	userRepository mongodb.UserRepositoryPort
}

func NewUserService(ur mongodb.UserRepositoryPort) UserServicePort {
	return &UserService{
		userRepository: ur,
	}
}

func (s *UserService) CreateUser(user *domain.User) error {
	return s.userRepository.Create(user)
}

func (s *UserService) UpdateUser(user *domain.User) error {
	return s.userRepository.Update(user)
}

func (s *UserService) DeleteUser(id string) error {
	return s.userRepository.Delete(id)
}

func (s *UserService) GetUserByID(id string) (*domain.User, error) {
	return s.userRepository.GetByID(id)
}

func (s *UserService) GetAllUser() ([]*domain.User, error) {
	return s.userRepository.GetAll()
}
