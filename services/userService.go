package services

import (
	"time"

	"github.com/ZaphCode/F-SR-ChatApp/domain"
	"github.com/ZaphCode/F-SR-ChatApp/utils"
	"github.com/google/uuid"
)

func NewUserService(userRepo domain.UserRepository) domain.UserService {
	return &userService{userRepo}
}

type userService struct {
	userRepo domain.UserRepository
}

func (s *userService) Create(username, email, password string) (domain.User, error) {
	id, err := uuid.NewUUID()

	if err != nil {
		return domain.User{}, err
	}

	hashPass, err := utils.HashPassword(password)

	if err != nil {
		return domain.User{}, err
	}

	user := domain.User{
		ID:        id,
		Username:  username,
		Email:     email,
		Password:  hashPass,
		Verified:  false,
		CreatedAt: time.Now(),
	}

	if err := s.userRepo.Save(&user); err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (s *userService) GetByID(id uuid.UUID) (domain.User, error) {
	return domain.User{}, nil
}

func (s *userService) Authenticate(email, password string) (domain.User, error) {
	user, err := s.userRepo.FindByEmail(email)

	if err != nil {
		return domain.User{}, err
	}

	if err := utils.VerifyHashedPassword(user.Password, password); err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (s *userService) GetByEmail(email string) (domain.User, error) {
	return domain.User{}, nil
}

func (s *userService) GetAll() ([]domain.User, error) {
	return nil, nil
}

func (s *userService) Update(id uuid.UUID, user domain.User) error {
	return nil
}

func (s *userService) Delete(id uuid.UUID) error {
	return s.userRepo.Remove(id)
}
