package services

import (
	"github.com/ZaphCode/F-SR-ChatApp/domain"
	"github.com/google/uuid"
)

func NewUserService(userRepo domain.UserRepository) domain.UserService {
	return &userService{userRepo}
}

type userService struct {
	userRepo domain.UserRepository
}

func (s *userService) Create(user *domain.User) error {
	return s.userRepo.Save(user)
}

func (s *userService) Get(id uuid.UUID) (domain.User, error) {
	return domain.User{}, nil
}

func (s *userService) GetAll() ([]domain.User, error) {
	return nil, nil
}

func (s *userService) Update(user *domain.User) error {
	return nil
}

func (s *userService) Delete(id uuid.UUID) error {
	return s.userRepo.Remove(id)
}
