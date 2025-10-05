package services

import (
	"errors"

	"github.com/ZaphCode/F-SR-ChatApp/domain"
	"github.com/google/uuid"
)

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
