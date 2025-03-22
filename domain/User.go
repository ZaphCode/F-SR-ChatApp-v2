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
	Role      string
	ImageUrl  string
	CreatedAt time.Time
}

type UserService interface {
	Create(user *User) error
	Get(id uuid.UUID) (User, error)
	GetAll() ([]User, error)
	Update(user *User) error
	Delete(id uuid.UUID) error
}

type UserRepository interface {
	FindByID(id uuid.UUID) (User, error)
	FindAll() ([]User, error)
	Save(user *User) error
	Update(user *User) error
	Remove(id uuid.UUID) error
}
