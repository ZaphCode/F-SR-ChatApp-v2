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

// Delete implements domain.UserService.
func (s *userService) Delete(id uuid.UUID) error {
	panic("unimplemented")
}

// Get implements domain.UserService.
func (s *userService) Get(id uuid.UUID) (domain.User, error) {
	panic("unimplemented")
}

// GetAll implements domain.UserService.
func (s *userService) GetAll() ([]domain.User, error) {
	panic("unimplemented")
}

// Update implements domain.UserService.
func (s *userService) Update(user *domain.User) error {
	panic("unimplemented")
}

func (s *userService) Create(user *domain.User) error {
	return s.userRepo.Save(user)
}
