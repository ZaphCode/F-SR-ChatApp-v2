package mocks

import (
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"

	"github.com/ZaphCode/F-SR-ChatApp/domain"
)

//* Users Mock

var UserA domain.User = domain.User{
	ID:        uuid.MustParse("11111111-1111-1111-1111-111111111111"),
	Username:  "User A",
	Email:     "userA@example.com",
	Password:  "$2a$10$NEeUSqIUfUSG/4hHsjuLEeK6JLigzCZbXy4DXHieYnYJurOjRvQbK", // hashed "password123"
	CreatedAt: time.Now(),
}

var UserB domain.User = domain.User{
	ID:        uuid.MustParse("22222222-2222-2222-2222-222222222222"),
	Username:  "User B",
	Email:     "userB@example.com",
	Password:  "$2a$10$HDcYGQ/64RASOHgqWrsZ1u7SAscVrlUkinbgP.m4hpHjWX3saLm1W", // hashed "password456"
	CreatedAt: time.Now(),
}

var UserC domain.User = domain.User{
	ID:        uuid.MustParse("33333333-3333-3333-3333-333333333333"),
	Username:  "User C",
	Email:     "userC@example.com",
	Password:  "$2a$10$HDcYGQ/64RASOHgqWrsZ1u7SAscVrlUkinbgP.m4hpHjWX3saLm1W", // same as User B
	CreatedAt: time.Now(),
}

//* User Repository Mock

type userRepositoryMock struct {
	mu    sync.RWMutex
	users map[uuid.UUID]domain.User
}

func NewUserRepository() domain.UserRepository {
	return &userRepositoryMock{
		mu: sync.RWMutex{},
		users: map[uuid.UUID]domain.User{
			UserA.ID: UserA,
			UserB.ID: UserB,
			UserC.ID: UserC,
		},
	}
}

func (r *userRepositoryMock) FindAll() ([]domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	users := make([]domain.User, 0, len(r.users))
	for _, u := range r.users {
		users = append(users, u)
	}
	return users, nil
}

func (r *userRepositoryMock) FindByEmail(email string) (domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, u := range r.users {
		if u.Email == email {
			return u, nil
		}
	}
	return domain.User{}, errors.New("user not found")
}

func (r *userRepositoryMock) FindByID(id uuid.UUID) (domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	u, ok := r.users[id]
	if !ok {
		return domain.User{}, errors.New("user not found")
	}
	return u, nil
}

func (r *userRepositoryMock) Save(user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[user.ID]; exists {
		return errors.New("user already exists")
	}
	r.users[user.ID] = *user
	return nil
}

func (r *userRepositoryMock) Update(id uuid.UUID, user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.users[id]; !ok {
		return errors.New("user not found")
	}
	r.users[id] = *user
	return nil
}

func (r *userRepositoryMock) Remove(id uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.users[id]; !ok {
		return errors.New("user not found")
	}
	delete(r.users, id)
	return nil
}
