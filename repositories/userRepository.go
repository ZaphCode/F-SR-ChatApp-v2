package repositories

import (
	"github.com/ZaphCode/F-SR-ChatApp/domain"
	"github.com/google/uuid"
)

func NewUserRepository() domain.UserRepository {
	return &userRepository{}
}

type userRepository struct{}

func (r *userRepository) FindAll() ([]domain.User, error) {
	return []domain.User{}, nil
}

func (r *userRepository) FindByID(id uuid.UUID) (domain.User, error) {
	return domain.User{}, nil
}

func (r *userRepository) Save(user *domain.User) error {
	return nil
}

func (r *userRepository) Update(user *domain.User) error {
	return nil
}

func (r *userRepository) Remove(id uuid.UUID) error {
	return nil
}
