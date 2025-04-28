package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Username  string
	Email     string
	Password  string
	Verified  bool
	ImageUrl  string
	CreatedAt time.Time
}

type UserService interface {
	Create(username, email, password string) (User, error)
	Authenticate(email, password string) (User, error)
	GetByID(id uuid.UUID) (User, error)
	GetAll() ([]User, error)
	Update(id uuid.UUID, user User) error
	Delete(id uuid.UUID) error
}

type UserRepository interface {
	Save(user *User) error
	FindByID(id uuid.UUID) (User, error)
	FindAll() ([]User, error)
	FindByEmail(email string) (User, error)
	Update(id uuid.UUID, user *User) error
	Remove(id uuid.UUID) error
}
