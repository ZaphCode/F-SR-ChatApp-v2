package services

import (
	"errors"
	"fmt"
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

	existingUser, err := s.userRepo.FindByEmail(email)

	if err == nil && existingUser.ID != uuid.Nil {
		return domain.User{}, errors.New("the email is already in use")
	}

	hashPass, err := utils.HashPassword(password)

	if err != nil {
		return domain.User{}, fmt.Errorf("failed to hash password: %w", err)
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
		return domain.User{}, fmt.Errorf("failed to save user: %w", err)
	}

	return user, nil
}

func (s *userService) GetByID(id uuid.UUID) (domain.User, error) {
	return s.userRepo.FindByID(id)
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
	return s.userRepo.FindByEmail(email)
}

func (s *userService) GetAll() ([]domain.User, error) {
	return s.userRepo.FindAll()
}

func (s *userService) UpdateProfileImg(id uuid.UUID, img string) error {
	return nil
}

func (s *userService) Delete(id uuid.UUID) error {
	return s.userRepo.Remove(id)
}

type userServiceMock struct{}

func NewUserServiceMock() domain.UserService {
	return &userServiceMock{}
}

func (s *userServiceMock) Create(username, email, password string) (domain.User, error) {
	return domain.User{
		ID:       uuid.New(),
		Username: username,
		Email:    email,
		Password: password,
	}, nil
}

func (s *userServiceMock) GetByID(id uuid.UUID) (domain.User, error) {
	return domain.User{
		ID:       id,
		Username: "testuser",
		Email:    "test@user.com",
		Password: "testpassword",
	}, nil
}

func (s *userServiceMock) Authenticate(email, password string) (domain.User, error) {
	if email == "test@user.com" && password == "testpassword" {
		return domain.User{
			ID:       uuid.New(),
			Username: "testuser",
			Email:    email,
			Password: password,
		}, nil
	}
	return domain.User{}, errors.New("invalid credentials")
}

func (s *userServiceMock) GetByEmail(email string) (domain.User, error) {
	if email == "test@user.com" {
		return domain.User{
			ID:       uuid.New(),
			Username: "testuser",
			Email:    email,
			Password: "testpassword",
		}, nil
	}
	return domain.User{}, errors.New("user not found")
}

func (s *userServiceMock) GetAll() ([]domain.User, error) {
	return []domain.User{
		{
			ID:       uuid.New(),
			Username: "testuser",
			Email:    "test@user.com",
			Password: "testpassword",
		},
	}, nil
}

func (s *userServiceMock) UpdateProfileImg(id uuid.UUID, img string) error {
	return nil
}

func (s *userServiceMock) Delete(id uuid.UUID) error {
	return nil
}
