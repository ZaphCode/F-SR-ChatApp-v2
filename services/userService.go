package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/ZaphCode/F-SR-ChatApp/domain"
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

	hashPass, err := s.hashPassword(password)

	if err != nil {
		return domain.User{}, fmt.Errorf("failed to hash password: %w", err)
	}

	user := domain.User{
		ID:        id,
		Username:  username,
		Email:     email,
		Password:  hashPass,
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

	if err := s.verifyHashedPassword(user.Password, password); err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (s *userService) GetAll() ([]domain.User, error) {
	return s.userRepo.FindAll()
}

func (s *userService) UpdateProfileImg(id uuid.UUID, img string) error {
	user, err := s.userRepo.FindByID(id)

	if err != nil {
		return err
	}

	user.ImageUrl = img

	return s.userRepo.Update(id, &user)
}

func (s *userService) Delete(id uuid.UUID) error {
	return s.userRepo.Remove(id)
}

// Helpers

func (s *userService) hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (s *userService) verifyHashedPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
